package service

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"github.com/MetaBloxIO/did-sdk-go"
	"github.com/MetaBloxIO/metablox-foundation-services/dao"
	"strconv"
	"strings"
	"time"
)

// create credential used to access wifi using the information provided in wifiAccessInfo.
// If a wifi credential already exists for this DID, return the existing credential
func CreateWifiAccessVC(issuerDocument *did.DIDDocument, wifiAccessInfo *did.WifiAccessInfo, issuerPrivKey *ecdsa.PrivateKey) (*did.VerifiableCredential, error) {
	var vc *did.VerifiableCredential
	exists, err := dao.CheckWifiAccessForExistence(wifiAccessInfo.ID)
	if err != nil {
		return nil, err
	}
	if exists {
		vc, err = dao.GetWifiAccessFromDB(wifiAccessInfo.ID)
		if err != nil {
			return nil, err
		}
		vc.Context = []string{did.ContextSecp256k1, did.ContextCredential}
		vc.Type = []string{did.TypeCredential, did.TypeWifi}
		vc.ID = did.BaseIDString + vc.ID
		err = did.ConvertTimesFromDBFormat(vc)
		if err != nil {
			return nil, err
		}
		vc.Proof = did.CreateProof(issuerDocument.Authentication)
	} else {

		vc, err = did.CreateVC(issuerDocument)
		if err != nil {
			return nil, err
		}

		vc.Type = append(vc.Type, did.TypeWifi)
		vc.Description = "Example Wifi Access Credential" //TODO: probably should fix this placeholder value at some point
		vc.CredentialSubject = *wifiAccessInfo

		//Upload VC to DB and generate ID. Has to be done before creating signature, as changing the ID will change the signature
		err = did.ConvertTimesToDBFormat(vc)
		if err != nil {
			return nil, err
		}

		generatedID, err := dao.UploadWifiAccessVC(*vc)
		if err != nil {
			return nil, err
		}

		err = did.ConvertTimesFromDBFormat(vc)
		if err != nil {
			return nil, err
		}

		vc.ID = did.BaseIDString + strconv.Itoa(generatedID)
	}

	//Create the proof's signature using a stringified version of the VC and the issuer's private key.
	//This way, the signature can be verified by re-stringifying the VC and looking up the public key in the issuer's DID document.
	//Verification will only succeed if the VC was unchanged since the signature and if the issuer
	//public key matches the private key used to make the signature
	hashedVC := sha256.Sum256(did.ConvertVCToBytes(*vc))

	signatureData, err := did.CreateJWSSignature(issuerPrivKey, hashedVC[:])
	if err != nil {
		return nil, err
	}
	vc.Proof.JWSSignature = signatureData

	return vc, nil
}

// create credential used by miners using the information provided in miningLicenseInfo.
// If a mining credential already exists for this DID, return the existing credential
func CreateMiningLicenseVC(issuerDocument *did.DIDDocument, miningLicenseInfo *did.MiningLicenseInfo, issuerPrivKey *ecdsa.PrivateKey) (*did.VerifiableCredential, error) {
	var vc *did.VerifiableCredential
	exists, err := dao.CheckMiningLicenseForExistence(miningLicenseInfo.ID)
	if err != nil {
		return nil, err
	}
	if exists {
		vc, err = dao.GetMiningLicenseFromDB(miningLicenseInfo.ID)
		if err != nil {
			return nil, err
		}
		vc.Context = []string{did.ContextSecp256k1, did.ContextCredential}
		vc.Type = []string{did.TypeCredential, did.TypeMining}
		vc.ID = did.BaseIDString + vc.ID
		err = did.ConvertTimesFromDBFormat(vc)
		if err != nil {
			return nil, err
		}
		vc.Proof = did.CreateProof(issuerDocument.Authentication)
	} else {

		vc, err = did.CreateVC(issuerDocument)
		if err != nil {
			return nil, err
		}

		vc.Type = append(vc.Type, did.TypeMining)
		vc.Description = "Example Mining License Credential" //TODO: probably should fix this placeholder value at some point
		vc.CredentialSubject = *miningLicenseInfo

		//Upload VC to DB and generate ID. Has to be done before creating signature, as changing the ID will change the signature
		err = did.ConvertTimesToDBFormat(vc)
		if err != nil {
			return nil, err
		}

		generatedID, err := dao.UploadMiningLicenseVC(*vc)
		if err != nil {
			return nil, err
		}

		err = did.ConvertTimesFromDBFormat(vc)
		if err != nil {
			return nil, err
		}

		vc.ID = did.BaseIDString + strconv.Itoa(generatedID)
	}

	//Create the proof's signature using a stringified version of the VC and the issuer's private key.
	//This way, the signature can be verified by re-stringifying the VC and looking up the public key in the issuer's DID document.
	//Verification will only succeed if the VC was unchanged since the signature and if the issuer
	//public key matches the private key used to make the signature
	hashedVC := sha256.Sum256(did.ConvertVCToBytes(*vc))

	signatureData, err := did.CreateJWSSignature(issuerPrivKey, hashedVC[:])
	if err != nil {
		return nil, err
	}
	vc.Proof.JWSSignature = signatureData

	return vc, nil
}

// update the provided credential's expiration date as well as its signature
func RenewVC(vc *did.VerifiableCredential, issuerPrivKey *ecdsa.PrivateKey) error {
	splitID := strings.Split(vc.ID, "/")
	idNum := splitID[len(splitID)-1]

	revoked, err := dao.GetCredentialStatusByID(idNum)
	if err != nil {
		return err
	}

	if revoked { //can't renew a revoked credential
		return did.ErrRenewRevoked
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

	hashedVC := sha256.Sum256(did.ConvertVCToBytes(*vc))

	signatureData, err := did.CreateJWSSignature(issuerPrivKey, hashedVC[:]) //since the expiration date has changed, the signature must also change
	if err != nil {
		return err
	}
	vc.Proof.JWSSignature = signatureData

	return nil
}

// revoke the provided credential. No need to update the signature since we're making the credential invalid anyways
func RevokeVC(vc *did.VerifiableCredential) error {
	vc.Revoked = true
	splitID := strings.Split(vc.ID, "/")
	idNum := splitID[len(splitID)-1]

	err := dao.RevokeVC(idNum)
	if err != nil {
		return err
	}
	return nil
}
