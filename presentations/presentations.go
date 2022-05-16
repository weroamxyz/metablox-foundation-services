package presentations

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/multiformats/go-multibase"
	"github.com/MetaBloxIO/metablox-foundation-services/credentials"
	"github.com/MetaBloxIO/metablox-foundation-services/did"
	"github.com/MetaBloxIO/metablox-foundation-services/errval"
	"github.com/MetaBloxIO/metablox-foundation-services/key"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
)

func CreatePresentation(credentials []models.VerifiableCredential, holderDocument models.DIDDocument, holderPrivKey *ecdsa.PrivateKey, nonce string) (*models.VerifiablePresentation, error) {
	presentationProof := models.CreateVPProof()
	presentationProof.Type = models.Secp256k1Sig
	presentationProof.VerificationMethod = holderDocument.Authentication
	presentationProof.JWSSignature = ""
	presentationProof.Created = time.Now().Format(time.RFC3339)
	presentationProof.ProofPurpose = "Authentication"
	presentationProof.Nonce = nonce
	context := []string{models.ContextCredential, models.ContextSecp256k1}
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
func VerifyVP(presentation *models.VerifiablePresentation) (bool, error) {
	resolutionMeta, holderDoc, _ := did.Resolve(presentation.Holder, models.CreateResolutionOptions())
	if resolutionMeta.Error != "" {
		return false, errors.New(resolutionMeta.Error)
	}

	targetVM, err := holderDoc.RetrieveVerificationMethod(presentation.Proof.VerificationMethod)
	if err != nil {
		return false, err
	}

	//currently only support EcdsaSecp256k1Signature2019, but it's possible we could introduce more
	var success bool
	switch presentation.Proof.Type {
	case models.Secp256k1Sig:
		if targetVM.MethodType != models.Secp256k1Key {
			return false, errval.ErrSecp256k1WrongVMType
		}
		success, err = VerifyVPSecp256k1(presentation, targetVM)
	default:
		return false, errval.ErrUnknownProofType
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

func VerifyVPSecp256k1(presentation *models.VerifiablePresentation, targetVM models.VerificationMethod) (bool, error) {
	copiedVP := *presentation
	//have to make sure to remove the signature from the copy, as the original did not have a signature at the time the signature was generated
	copiedVP.Proof.JWSSignature = ""
	hashedVP := sha256.Sum256(ConvertVPToBytes(copiedVP))
	_, pubData, err := multibase.Decode(targetVM.MultibaseKey)
	if err != nil {
		return false, err
	}

	pubKey, err := crypto.UnmarshalPubkey(pubData)
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
