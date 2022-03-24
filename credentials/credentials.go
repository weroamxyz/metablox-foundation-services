package credentials

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/dappley/go-dappley/crypto/keystore/secp256k1"
	"github.com/metabloxDID/did"
	"github.com/metabloxDID/models"
)

const sampleTrustedIssuer = "did:metablox:sampleIssuer"

//In the future, will probably need to set up multiple different creation functions for different types of VCs.
//This function serves as an example of making a resident card
func CreateVC(issuerDocument *models.DIDDocument, subjectInfo *models.SubjectInfo, issuerPrivKey []byte) (*models.VerifiableCredential, error) {
	newVC := models.CreateVerifiableCredential()
	newVC.Context = make([]string, 0)
	newVC.Context = append(newVC.Context, "https://www.w3.org/2018/credentials/v1")
	newVC.Type = make([]string, 0)
	newVC.Type = append(newVC.Type, "VerifiableCredential")
	newVC.Type = append(newVC.Type, "PermanentResidentCard")
	newVC.Issuer = issuerDocument.ID
	newVC.IssuanceDate = time.Now().Format(time.RFC3339)
	newVC.ExpirationDate = time.Now().AddDate(10, 0, 0).Format(time.RFC3339) //arbitrarily setting VCs to last for 10 years for the moment, can change when necessary
	newVC.Description = "Government of Example Permanent Resident Card"
	newVC.CredentialSubject = *subjectInfo //subject info is gathered ahead of time through an input form or some other means

	vcProof := models.CreateVCProof()
	vcProof.Type = "Secp256k1"
	vcProof.VerificationMethod = issuerDocument.Authentication
	vcProof.SignatureValue = ""
	newVC.Proof = *vcProof
	//Create the proof's signature using a stringified version of the VC and the issuer's private key.
	//This way, the signature can be verified by re-stringifying the VC and looking up the public key in the issuer's DID document.
	//Verification will only succeed if the VC was unchanged since the signature and if the issuer
	//public key matches the private key used to make the signature
	stringVC := fmt.Sprintf("%v", *newVC)
	hashedVC := sha256.Sum256([]byte(stringVC))

	signatureData, err := secp256k1.Sign(hashedVC[:], issuerPrivKey)
	if err != nil {
		return nil, err
	}
	newVC.Proof.SignatureValue = string(signatureData)

	return newVC, nil
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
		return false, errors.New("unknown issuer")
	}

	resolutionMeta, issuerDoc, _ := did.Resolve(vc.Issuer, models.CreateResolutionOptions())
	if resolutionMeta.Error != "" {
		return false, errors.New("failed to resolve issuer document: " + resolutionMeta.Error)
	}

	targetVM, err := issuerDoc.RetrieveVerificationMethod(vc.Proof.VerificationMethod)
	if err != nil {
		return false, err
	}

	if targetVM.MethodType != vc.Proof.Type {
		return false, errors.New("proof type (" + vc.Proof.Type + ") does not match verification method type(" + targetVM.MethodType + ")")
	}

	//currently only support Secp256k1, but it's possible we could introduce more
	switch vc.Proof.Type {
	case "Secp256k1":
		return VerifyVCSecp256k1(vc, targetVM)
	default:
		return false, errors.New("unable to verify unknown proof type " + vc.Proof.Type)
	}
}

func VerifyVCSecp256k1(vc *models.VerifiableCredential, targetVM models.VerificationMethod) (bool, error) {

	copiedVC := *vc
	//have to make sure to remove the signature from the copy, as the original did not have a signature at the time the signature was generated
	copiedVC.Proof.SignatureValue = ""
	stringVC := fmt.Sprintf("%v", copiedVC)
	hashedVC := sha256.Sum256([]byte(stringVC))
	pubData, err := hex.DecodeString(targetVM.Key)
	if err != nil {
		return false, err
	}
	result, err := secp256k1.Verify(hashedVC[:], []byte(vc.Proof.SignatureValue), pubData)
	if err != nil {
		return false, err
	}
	return result, nil
}
