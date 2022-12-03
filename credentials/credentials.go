package credentials

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/MetaBloxIO/metablox-foundation-services/dao"
	"github.com/MetaBloxIO/metablox-foundation-services/did"
	"github.com/MetaBloxIO/metablox-foundation-services/errval"
	"github.com/MetaBloxIO/metablox-foundation-services/key"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/ethereum/go-ethereum/crypto"
)

var IssuerDID string
var IssuerPrivateKey *ecdsa.PrivateKey

// All credential ids use a format of this value plus a number. ex. 'http://metablox.com/credentials/5'
// Only the number is stored in the db as the ID; the full string is only used in formal credentials
const baseIDString = "http://metablox.com/credentials/"

// read in private key for issuer and generate the DID from it
func InitializeValues() error {
	var err error
	IssuerPrivateKey, err = key.GetIssuerPrivateKey()
	if err != nil {
		return err
	}
	IssuerDID = did.GenerateDIDString(IssuerPrivateKey)
	return nil
}

// create a credential proof using the provided verification method string
func CreateProof(vm string) models.VCProof {
	vcProof := models.CreateVCProof()
	vcProof.Type = models.Secp256k1Sig
	vcProof.VerificationMethod = vm
	vcProof.JWSSignature = ""
	vcProof.Created = time.Now().Format(time.RFC3339)
	vcProof.ProofPurpose = models.PurposeAuth
	vcProof.PublicKeyString = crypto.FromECDSAPub(&IssuerPrivateKey.PublicKey)
	return *vcProof
}

// convert issuance and expiration times of credential from db format to RFC3339
func ConvertTimesFromDBFormat(vc *models.VerifiableCredential) error {
	issuanceTime, err := time.Parse("2006-01-02 15:04:05", vc.IssuanceDate)
	if err != nil {
		return err
	}
	vc.IssuanceDate = issuanceTime.Format(time.RFC3339)

	expirationTime, err := time.Parse("2006-01-02 15:04:05", vc.ExpirationDate)
	if err != nil {
		return err
	}
	vc.ExpirationDate = expirationTime.Format(time.RFC3339)
	return nil
}

// convert issuance and expiration times of credential from RFC3339 to db format
func ConvertTimesToDBFormat(vc *models.VerifiableCredential) error {
	issuanceTime, err := time.Parse(time.RFC3339, vc.IssuanceDate)
	if err != nil {
		return err
	}
	vc.IssuanceDate = issuanceTime.Format("2006-01-02 15:04:05")

	expirationTime, err := time.Parse(time.RFC3339, vc.ExpirationDate)
	if err != nil {
		return err
	}
	vc.ExpirationDate = expirationTime.Format("2006-01-02 15:04:05")
	return nil
}

// Base function for creating VCs. Called by any function that creates a type of VC to initialize universal values
func CreateVC(issuerDocument *models.DIDDocument) (*models.VerifiableCredential, error) {
	context := []string{models.ContextSecp256k1, models.ContextCredential}
	vcType := []string{models.TypeCredential}
	loc, _ := time.LoadLocation("UTC")
	expirationDate := time.Now().In(loc).AddDate(10, 0, 0).Format(time.RFC3339) //arbitrarily setting VCs to last for 10 years for the moment, can change when necessary
	description := ""

	vcProof := CreateProof(issuerDocument.Authentication)

	newVC := models.NewVerifiableCredential(context, "0", vcType, issuerDocument.ID, time.Now().In(loc).Format(time.RFC3339), expirationDate, description, nil, vcProof, false)

	return newVC, nil
}

// create credential used to access wifi using the information provided in wifiAccessInfo.
// If a wifi credential already exists for this DID, return the existing credential
func CreateWifiAccessVC(issuerDocument *models.DIDDocument, wifiAccessInfo *models.WifiAccessInfo, issuerPrivKey *ecdsa.PrivateKey) (*models.VerifiableCredential, error) {
	var vc *models.VerifiableCredential
	exists, err := dao.CheckWifiAccessForExistence(wifiAccessInfo.ID)
	if err != nil {
		return nil, err
	}
	if exists {
		vc, err = dao.GetWifiAccessFromDB(wifiAccessInfo.ID)
		if err != nil {
			return nil, err
		}
		vc.Context = []string{models.ContextSecp256k1, models.ContextCredential}
		vc.Type = []string{models.TypeCredential, models.TypeWifi}
		vc.ID = baseIDString + vc.ID
		err = ConvertTimesFromDBFormat(vc)
		if err != nil {
			return nil, err
		}
		vc.Proof = CreateProof(issuerDocument.Authentication)
	} else {

		vc, err = CreateVC(issuerDocument)
		if err != nil {
			return nil, err
		}

		vc.Type = append(vc.Type, models.TypeWifi)
		vc.Description = "Example Wifi Access Credential" //TODO: probably should fix this placeholder value at some point
		vc.CredentialSubject = *wifiAccessInfo

		//Upload VC to DB and generate ID. Has to be done before creating signature, as changing the ID will change the signature
		err = ConvertTimesToDBFormat(vc)
		if err != nil {
			return nil, err
		}

		generatedID, err := dao.UploadWifiAccessVC(*vc)
		if err != nil {
			return nil, err
		}

		err = ConvertTimesFromDBFormat(vc)
		if err != nil {
			return nil, err
		}

		vc.ID = baseIDString + strconv.Itoa(generatedID)
	}

	//Create the proof's signature using a stringified version of the VC and the issuer's private key.
	//This way, the signature can be verified by re-stringifying the VC and looking up the public key in the issuer's DID document.
	//Verification will only succeed if the VC was unchanged since the signature and if the issuer
	//public key matches the private key used to make the signature
	hashedVC := sha256.Sum256(ConvertVCToBytes(*vc))

	signatureData, err := key.CreateJWSSignature(issuerPrivKey, hashedVC[:])
	if err != nil {
		return nil, err
	}
	vc.Proof.JWSSignature = signatureData

	return vc, nil
}

// create credential used by miners using the information provided in miningLicenseInfo.
// If a mining credential already exists for this DID, return the existing credential
func CreateMiningLicenseVC(issuerDocument *models.DIDDocument, miningLicenseInfo *models.MiningLicenseInfo, issuerPrivKey *ecdsa.PrivateKey) (*models.VerifiableCredential, error) {
	var vc *models.VerifiableCredential
	exists, err := dao.CheckMiningLicenseForExistence(miningLicenseInfo.ID)
	if err != nil {
		return nil, err
	}
	if exists {
		vc, err = dao.GetMiningLicenseFromDB(miningLicenseInfo.ID)
		if err != nil {
			return nil, err
		}
		vc.Context = []string{models.ContextSecp256k1, models.ContextCredential}
		vc.Type = []string{models.TypeCredential, models.TypeMining}
		vc.ID = baseIDString + vc.ID
		err = ConvertTimesFromDBFormat(vc)
		if err != nil {
			return nil, err
		}
		vc.Proof = CreateProof(issuerDocument.Authentication)
	} else {

		vc, err = CreateVC(issuerDocument)
		if err != nil {
			return nil, err
		}

		vc.Type = append(vc.Type, models.TypeMining)
		vc.Description = "Example Mining License Credential" //TODO: probably should fix this placeholder value at some point
		vc.CredentialSubject = *miningLicenseInfo

		//Upload VC to DB and generate ID. Has to be done before creating signature, as changing the ID will change the signature
		err = ConvertTimesToDBFormat(vc)
		if err != nil {
			return nil, err
		}

		generatedID, err := dao.UploadMiningLicenseVC(*vc)
		if err != nil {
			return nil, err
		}

		err = ConvertTimesFromDBFormat(vc)
		if err != nil {
			return nil, err
		}

		vc.ID = baseIDString + strconv.Itoa(generatedID)
	}

	//Create the proof's signature using a stringified version of the VC and the issuer's private key.
	//This way, the signature can be verified by re-stringifying the VC and looking up the public key in the issuer's DID document.
	//Verification will only succeed if the VC was unchanged since the signature and if the issuer
	//public key matches the private key used to make the signature
	hashedVC := sha256.Sum256(ConvertVCToBytes(*vc))

	signatureData, err := key.CreateJWSSignature(issuerPrivKey, hashedVC[:])
	if err != nil {
		return nil, err
	}
	vc.Proof.JWSSignature = signatureData

	return vc, nil
}

// update the provided credential's expiration date as well as its signature
func RenewVC(vc *models.VerifiableCredential, issuerPrivKey *ecdsa.PrivateKey) error {
	splitID := strings.Split(vc.ID, "/")
	idNum := splitID[len(splitID)-1]

	revoked, err := dao.GetCredentialStatusByID(idNum)
	if err != nil {
		return err
	}

	if revoked { //can't renew a revoked credential
		return errval.ErrRenewRevoked
	}

	oldExpirationDate, err := time.Parse(time.RFC3339, vc.ExpirationDate)
	if err != nil {
		return err
	}

	newExpirationDate := oldExpirationDate.AddDate(1, 0, 0) //TODO: come up with better logic for how much to extend expiration date by
	vc.ExpirationDate = newExpirationDate.Format(time.RFC3339)
	err = dao.UpdateVCExpirationDate(idNum, newExpirationDate.Format("2006-01-02 15:04:05"))
	if err != nil {
		return err
	}

	vc.Proof.JWSSignature = ""
	vc.Proof.Created = time.Now().Format(time.RFC3339)

	hashedVC := sha256.Sum256(ConvertVCToBytes(*vc))

	signatureData, err := key.CreateJWSSignature(issuerPrivKey, hashedVC[:]) //since the expiration date has changed, the signature must also change
	if err != nil {
		return err
	}
	vc.Proof.JWSSignature = signatureData

	return nil
}

// revoke the provided credential. No need to update the signature since we're making the credential invalid anyways
func RevokeVC(vc *models.VerifiableCredential) error {
	vc.Revoked = true
	splitID := strings.Split(vc.ID, "/")
	idNum := splitID[len(splitID)-1]

	err := dao.RevokeVC(idNum)
	if err != nil {
		return err
	}
	return nil
}

// convert credential to a JSON format. Currently unused
func VCToJson(vc *models.VerifiableCredential) ([]byte, error) {
	jsonVC, err := json.Marshal(vc)
	if err != nil {
		return nil, err
	}
	return jsonVC, nil
}

// convert JSON formatted credential to object. Currently unused
func JsonToVC(jsonVC []byte) (*models.VerifiableCredential, error) {
	vc := models.CreateVerifiableCredential()
	err := json.Unmarshal(jsonVC, vc)
	if err != nil {
		return nil, err
	}
	return vc, nil
}

// Need to make sure that the stated issuer of the VC actually created it (using the proof alongside the issuer's verification methods),
// as well as check that the issuer is a trusted source
func VerifyVC(vc *models.VerifiableCredential) (bool, error) {
	if vc.Issuer != IssuerDID { //issuer of VC must be the same issuer stored here
		return false, errval.ErrUnknownIssuer
	}

	resolutionMeta, issuerDoc, _ := did.Resolve(vc.Issuer, models.CreateResolutionOptions())
	if resolutionMeta.Error != "" {
		return false, errors.New(resolutionMeta.Error)
	}

	//get verification method from the issuer DID document which is listed in the vc proof
	targetVM, err := issuerDoc.RetrieveVerificationMethod(vc.Proof.VerificationMethod)
	if err != nil {
		return false, err
	}

	//get public key stored in the vc proof
	publicKey, err := crypto.UnmarshalPubkey(vc.Proof.PublicKeyString)
	if err != nil {
		return false, err
	}

	//currently only support EcdsaSecp256k1Signature2019, but it's possible we could introduce more
	switch vc.Proof.Type {
	case models.Secp256k1Sig:
		if targetVM.MethodType != models.Secp256k1Key { //vm must be the same type as the proof
			return false, errval.ErrSecp256k1WrongVMType
		}

		success := key.CompareAddresses(targetVM, publicKey) //vm must have the address that matches the proof's public key
		if !success {
			return false, errval.ErrWrongAddress
		}

		return VerifyVCSecp256k1(vc, publicKey)
	default:
		return false, errval.ErrUnknownProofType
	}
}

// Verify that the provided public key matches the signature in the proof.
// Since we've made sure that the address in the issuer vm matches this public key,
// verifying the signature here proves that the signature was made with the issuer's private key
func VerifyVCSecp256k1(vc *models.VerifiableCredential, pubKey *ecdsa.PublicKey) (bool, error) {
	copiedVC := *vc
	//have to make sure to remove the signature from the copy, as the original did not have a signature at the time the signature was generated
	copiedVC.Proof.JWSSignature = ""
	hashedVC := sha256.Sum256(ConvertVCToBytes(copiedVC))

	result, err := key.VerifyJWSSignature(vc.Proof.JWSSignature, pubKey, hashedVC[:])
	if err != nil {
		return false, err
	}
	return result, nil
}

// convert credential to bytes so it can be hashed
func ConvertVCToBytes(vc models.VerifiableCredential) []byte {
	var convertedBytes []byte

	for _, item := range vc.Context {
		convertedBytes = bytes.Join([][]byte{convertedBytes, []byte(item)}, []byte{})
	}

	for _, item := range vc.Type {
		convertedBytes = bytes.Join([][]byte{convertedBytes, []byte(item)}, []byte{})
	}

	convertedBytes = bytes.Join([][]byte{convertedBytes, []byte(vc.Issuer), []byte(vc.IssuanceDate), []byte(vc.ExpirationDate), []byte(vc.Description)}, []byte{})

	switch vc.Type[1] {
	case models.TypeWifi:
		wifiAccessInfo := vc.CredentialSubject.(models.WifiAccessInfo)
		convertedBytes = bytes.Join([][]byte{convertedBytes, []byte(wifiAccessInfo.ID), []byte(wifiAccessInfo.Type)}, []byte{})
	case models.TypeMining:
		miningLicenseInfo := vc.CredentialSubject.(models.MiningLicenseInfo)
		convertedBytes = bytes.Join([][]byte{convertedBytes, []byte(miningLicenseInfo.ID), []byte(miningLicenseInfo.Name), []byte(miningLicenseInfo.Model), []byte(miningLicenseInfo.Serial)}, []byte{})
	}

	convertedBytes = bytes.Join([][]byte{convertedBytes, []byte(vc.Proof.Type), []byte(vc.Proof.Created), []byte(vc.Proof.VerificationMethod), []byte(vc.Proof.ProofPurpose), []byte(vc.Proof.JWSSignature), vc.Proof.PublicKeyString}, []byte{})
	return convertedBytes
}
