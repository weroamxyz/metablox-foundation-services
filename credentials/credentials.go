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

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/metabloxDID/dao"
	"github.com/metabloxDID/did"
	"github.com/metabloxDID/errval"
	"github.com/metabloxDID/key"
	"github.com/metabloxDID/models"
	"github.com/multiformats/go-multibase"
)

const sampleTrustedIssuer = "did:metablox:sampleIssuer"
const baseIDString = "http://metablox.com/credentials/"

//Base function for creating VCs. Called by any function that creates a type of VC
func CreateVC(issuerDocument *models.DIDDocument) (*models.VerifiableCredential, error) {
	context := []string{models.ContextCredential, models.ContextSecp256k1}
	vcType := []string{models.TypeCredential}
	expirationDate := time.Now().AddDate(10, 0, 0).Format(time.RFC3339) //arbitrarily setting VCs to last for 10 years for the moment, can change when necessary
	description := ""

	vcProof := models.CreateVCProof()
	vcProof.Type = models.Secp256k1Sig
	vcProof.VerificationMethod = issuerDocument.Authentication
	vcProof.JWSSignature = ""
	vcProof.Created = time.Now().Format(time.RFC3339)
	vcProof.ProofPurpose = models.PurposeAuth

	newVC := models.NewVerifiableCredential(context, "0", vcType, "", issuerDocument.ID, time.Now().Format(time.RFC3339), expirationDate, description, nil, *vcProof, false)

	return newVC, nil
}

func CreateWifiAccessVC(issuerDocument *models.DIDDocument, wifiAccessInfo *models.WifiAccessInfo, issuerPrivKey *ecdsa.PrivateKey) (*models.VerifiableCredential, error) {
	exists, err := dao.CheckWifiAccessForExistence(wifiAccessInfo.ID)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errval.ErrWifiExists
	}

	newVC, err := CreateVC(issuerDocument)
	if err != nil {
		return nil, err
	}

	newVC.Type = append(newVC.Type, models.TypeWifi)
	newVC.SubType = models.TypeWifi
	newVC.Description = "Example Wifi Access Credential"
	newVC.CredentialSubject = *wifiAccessInfo

	//Upload VC to DB and generate ID. Has to be done before creating signature, as changing the ID will change the signature
	generatedID, err := dao.UploadWifiAccessVC(*newVC)
	if err != nil {
		return nil, err
	}
	newVC.ID = baseIDString + strconv.Itoa(generatedID)

	//Create the proof's signature using a stringified version of the VC and the issuer's private key.
	//This way, the signature can be verified by re-stringifying the VC and looking up the public key in the issuer's DID document.
	//Verification will only succeed if the VC was unchanged since the signature and if the issuer
	//public key matches the private key used to make the signature
	hashedVC := sha256.Sum256(ConvertVCToBytes(*newVC))

	signatureData, err := key.CreateJWSSignature(issuerPrivKey, hashedVC[:])
	if err != nil {
		return nil, err
	}
	newVC.Proof.JWSSignature = signatureData

	return newVC, nil
}

func CreateMiningLicenseVC(issuerDocument *models.DIDDocument, miningLicenseInfo *models.MiningLicenseInfo, issuerPrivKey *ecdsa.PrivateKey) (*models.VerifiableCredential, error) {
	exists, err := dao.CheckMiningLicenseForExistence(miningLicenseInfo.ID)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errval.ErrMiningExists
	}

	newVC, err := CreateVC(issuerDocument)
	if err != nil {
		return nil, err
	}

	newVC.Type = append(newVC.Type, models.TypeMining)
	newVC.SubType = models.TypeMining
	newVC.Description = "Example Mining License Credential"
	newVC.CredentialSubject = *miningLicenseInfo

	//Upload VC to DB and generate ID. Has to be done before creating signature, as changing the ID will change the signature
	generatedID, err := dao.UploadMiningLicenseVC(*newVC)
	if err != nil {
		return nil, err
	}
	newVC.ID = baseIDString + strconv.Itoa(generatedID)

	//Create the proof's signature using a stringified version of the VC and the issuer's private key.
	//This way, the signature can be verified by re-stringifying the VC and looking up the public key in the issuer's DID document.
	//Verification will only succeed if the VC was unchanged since the signature and if the issuer
	//public key matches the private key used to make the signature
	hashedVC := sha256.Sum256(ConvertVCToBytes(*newVC))

	signatureData, err := key.CreateJWSSignature(issuerPrivKey, hashedVC[:])
	if err != nil {
		return nil, err
	}
	newVC.Proof.JWSSignature = signatureData

	return newVC, nil
}

func RenewVC(vc *models.VerifiableCredential) error {
	splitID := strings.Split(vc.ID, "/")
	idNum := splitID[len(splitID)-1]

	revoked, err := dao.GetCredentialStatusByID(idNum)
	if err != nil {
		return err
	}

	if revoked {
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
	return nil
}

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

func VCToJson(vc *models.VerifiableCredential) ([]byte, error) {
	jsonVC, err := json.Marshal(vc)
	if err != nil {
		return nil, err
	}
	return jsonVC, nil
}

func JsonToVC(jsonVC []byte) (*models.VerifiableCredential, error) {
	vc := models.CreateVerifiableCredential()
	err := json.Unmarshal(jsonVC, vc)
	if err != nil {
		return nil, err
	}
	return vc, nil
}

//Need to make sure that the stated issuer of the VC actually created it (using the proof alongside the issuer's verification methods),
//as well as check that the issuer is a trusted source
func VerifyVC(vc *models.VerifiableCredential) (bool, error) {
	//can modify to match the DID of the actual trusted issuer(s). May also want different
	//trusted issuers for different types of VCs
	if vc.Issuer != sampleTrustedIssuer {
		return false, errval.ErrUnknownIssuer
	}

	resolutionMeta, issuerDoc, _ := did.Resolve(vc.Issuer, models.CreateResolutionOptions())
	if resolutionMeta.Error != "" {
		return false, errors.New(resolutionMeta.Error)
	}

	targetVM, err := issuerDoc.RetrieveVerificationMethod(vc.Proof.VerificationMethod)
	if err != nil {
		return false, err
	}

	//currently only support EcdsaSecp256k1Signature2019, but it's possible we could introduce more
	switch vc.Proof.Type {
	case models.Secp256k1Sig:
		if targetVM.MethodType != models.Secp256k1Key {
			return false, errval.ErrSecp256k1WrongVMType
		}
		return VerifyVCSecp256k1(vc, targetVM)
	default:
		return false, errval.ErrUnknownProofType
	}
}

func VerifyVCSecp256k1(vc *models.VerifiableCredential, targetVM models.VerificationMethod) (bool, error) {
	copiedVC := *vc
	//have to make sure to remove the signature from the copy, as the original did not have a signature at the time the signature was generated
	copiedVC.Proof.JWSSignature = ""
	hashedVC := sha256.Sum256(ConvertVCToBytes(copiedVC))
	_, pubData, err := multibase.Decode(targetVM.MultibaseKey)
	if err != nil {
		return false, err
	}

	pubKey, err := crypto.UnmarshalPubkey(pubData)
	if err != nil {
		return false, err
	}
	result, err := key.VerifyJWSSignature(vc.Proof.JWSSignature, pubKey, hashedVC[:])
	if err != nil {
		return false, err
	}
	return result, nil
}

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

	convertedBytes = bytes.Join([][]byte{convertedBytes, []byte(vc.Proof.Type), []byte(vc.Proof.Created), []byte(vc.Proof.VerificationMethod), []byte(vc.Proof.ProofPurpose), []byte(vc.Proof.JWSSignature)}, []byte{})
	return convertedBytes
}
