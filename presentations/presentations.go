package presentations

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"errors"
	"time"

	"github.com/dappley/go-dappley/crypto/keystore/secp256k1"
	"github.com/metabloxDID/credentials"
	"github.com/metabloxDID/did"
	"github.com/metabloxDID/key"
	"github.com/metabloxDID/models"
	"github.com/multiformats/go-multibase"
)

func CreatePresentation(credentials []models.VerifiableCredential, holderDocument models.DIDDocument, holderPrivKey *ecdsa.PrivateKey, nonce string) (*models.VerifiablePresentation, error) {
	presentationProof := models.CreateVPProof()
	presentationProof.Type = "EcdsaSecp256k1Signature2019"
	presentationProof.VerificationMethod = holderDocument.Authentication
	presentationProof.JWSSignature = ""
	presentationProof.Created = time.Now().Format(time.RFC3339)
	presentationProof.ProofPurpose = "Authentication"
	presentationProof.Nonce = nonce
	context := []string{"https://www.w3.org/2018/credentials/v1", "https://ns.did.ai/suites/secp256k1-2019/v1/"}
	presentationType := []string{"VerifiablePresentation"}
	presentation := models.NewPresentation(context, presentationType, credentials, holderDocument.ID, *presentationProof)
	//Create the proof's signature using a stringified version of the VP and the holder's private key.
	//This way, the signature can be verified by re-stringifying the VP and looking up the public key in the holder's DID document.
	//Verification will only succeed if the VP was unchanged since the signature and if the holder
	//public key matches the private key used to make the signature

	//This proof is only for the presentation itself; each credential also needs to be individually verified
	hashedVP := sha256.Sum256(ConvertVPToBytes(*presentation))

	signatureData, err := key.CreateJWSSignature(holderPrivKey, hashedVP[:])
	if err != nil {
		return nil, err
	}
	presentation.Proof.JWSSignature = signatureData
	return presentation, nil
}

//Need to first verify the presentation's proof using the holder's DID document. Afterwards, need to verify
//the proof of each credential included inside the presentation
func VerifyVP(presentation *models.VerifiablePresentation, nonce string) (bool, error) {
	resolutionMeta, holderDoc, _ := did.Resolve(presentation.Holder, models.CreateResolutionOptions())
	if resolutionMeta.Error != "" {
		return false, errors.New("failed to resolve holder document: " + resolutionMeta.Error)
	}

	targetVM, err := holderDoc.RetrieveVerificationMethod(presentation.Proof.VerificationMethod)
	if err != nil {
		return false, err
	}

	//currently only support EcdsaSecp256k1Signature2019, but it's possible we could introduce more
	var success bool
	switch presentation.Proof.Type {
	case "EcdsaSecp256k1Signature2019":
		if targetVM.MethodType != "EcdsaSecp256k1VerificationKey2019" {
			return false, errors.New("must use a verification method with a type of 'EcdsaSecp256k1VerificationKey2019' to verify a 'EcdsaSecp256k1Signature2019' proof")
		}
		success, err = VerifyVPSecp256k1(presentation, targetVM, nonce)
	default:
		return false, errors.New("unable to verify unknown proof type " + presentation.Proof.Type)
	}

	if !success {
		return false, err
	}

	for _, credential := range presentation.VerifiableCredential {
		success, err = credentials.VerifyVC(&credential)
		if !success {
			return false, err
		}
	}

	return true, nil
}

func VerifyVPSecp256k1(presentation *models.VerifiablePresentation, targetVM models.VerificationMethod, nonce string) (bool, error) {
	//presentation must include the requested nonce
	if presentation.Proof.Nonce != nonce {
		return false, nil
	}
	copiedVP := *presentation
	//have to make sure to remove the signature from the copy, as the original did not have a signature at the time the signature was generated
	copiedVP.Proof.JWSSignature = ""
	hashedVP := sha256.Sum256(ConvertVPToBytes(copiedVP))
	_, pubData, err := multibase.Decode(targetVM.MultibaseKey)
	if err != nil {
		return false, err
	}
	pubKey, err := secp256k1.ToECDSAPublicKey(pubData)
	if err != nil {
		return false, err
	}
	result, err := key.VerifyJWSSignature(presentation.Proof.JWSSignature, pubKey, hashedVP[:])
	if err != nil {
		return false, err
	}
	return result, nil
}

func ConvertVPToBytes(vp models.VerifiablePresentation) []byte {
	var convertedBytes []byte
	for _, item := range vp.Context {
		convertedBytes = bytes.Join([][]byte{convertedBytes, []byte(item)}, []byte{})
	}

	for _, item := range vp.Type {
		convertedBytes = bytes.Join([][]byte{convertedBytes, []byte(item)}, []byte{})
	}

	for _, item := range vp.VerifiableCredential {
		convertedBytes = bytes.Join([][]byte{convertedBytes, credentials.ConvertVCToBytes(item)}, []byte{})
	}

	convertedBytes = bytes.Join([][]byte{convertedBytes, []byte(vp.Holder), []byte(vp.Proof.Type), []byte(vp.Proof.Created), []byte(vp.Proof.VerificationMethod), []byte(vp.Proof.ProofPurpose), []byte(vp.Proof.JWSSignature), []byte(vp.Proof.Nonce)}, []byte{})
	return convertedBytes
}
