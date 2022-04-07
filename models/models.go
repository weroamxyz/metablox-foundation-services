package models

import (
	"crypto/ecdsa"
	"errors"
	"time"

	"github.com/dappley/go-dappley/crypto/keystore/secp256k1"
)

type DIDDocument struct {
	Context            []string             `json:"@context"`
	ID                 string               `json:"id"`
	Created            string               `json:"created"`
	Updated            string               `json:"updated"`
	Version            int                  `json:"version"`
	VerificationMethod []VerificationMethod `json:"verificationMethod"`
	Authentication     string               `json:"authentication"`
}

type VerificationMethod struct {
	ID           string `json:"id"`
	MethodType   string `json:"type" mapstructure:"type"`
	Controller   string `json:"controller"`
	MultibaseKey string `json:"publicKeyMultibase" mapstructure:"publicKeyMultibase"`
}

type ResolutionOptions struct {
	Accept string `json:"accept"`
}

type RepresentationResolutionOptions struct {
	Accept string `json:"accept"`
}

type ResolutionMetadata struct {
	Error string `json:"error"`
}

type RepresentationResolutionMetadata struct {
	ContentType string `json:"contentType"`
	Error       string `json:"error"`
}

type DocumentMetadata struct {
	Created       string   `json:"created"`
	Updated       string   `json:"updated"`
	Deactivated   string   `json:"deactivated"`
	NextUpdate    string   `json:"nextUpdate"`
	VersionID     string   `json:"versionId"`
	NextVersionID string   `json:"nextVersionId"`
	EquivalentID  []string `json:"equivalentId"`
	CanonicalID   string   `json:"canonicalId"`
}

type VerifiableCredential struct {
	Context           []string    `json:"@context"`
	Type              []string    `json:"type"`
	Issuer            string      `json:"issuer"`
	IssuanceDate      string      `json:"issuanceDate"`
	ExpirationDate    string      `json:"expirationDate"`
	Description       string      `json:"description"`
	CredentialSubject SubjectInfo `json:"credentialSubject"`
	Proof             VCProof     `json:"proof"`
}

//This can be a type of input form to set up the VC.
//Temp fields here currently, will be changed in the future
type SubjectInfo struct {
	ID           string   `json:"id"`
	Type         []string `json:"type"`
	GivenName    string   `json:"givenName"`
	FamilyName   string   `json:"familyName"`
	Gender       string   `json:"gender"`
	BirthCountry string   `json:"birthCountry"`
	BirthDate    string   `json:"birthName"`
}

type VCProof struct {
	Type               string `json:"type"`
	Created            string `json:"created"`
	VerificationMethod string `json:"verificationMethod"`
	ProofPurpose       string `json:"proofPurpose"`
	JWSSignature       string `json:"jws"` //signature is created from a hash of the issuer's DID document
}

func CreateDIDDocument() *DIDDocument {
	return &DIDDocument{}
}

func (doc DIDDocument) RetrieveVerificationMethod(vmID string) (VerificationMethod, error) {
	for _, vm := range doc.VerificationMethod {
		if vm.ID == vmID {
			return vm, nil
		}
	}
	return VerificationMethod{}, errors.New("failed to find verification method with ID " + vmID)
}

func CreateVerifiableCredential() *VerifiableCredential {
	return &VerifiableCredential{}
}

func InitializeVerifiableCredential(context, vctype []string, issuer, expirationDate, description string, subject SubjectInfo, proof VCProof) *VerifiableCredential {
	return &VerifiableCredential{Context: context, Type: vctype, Issuer: issuer, IssuanceDate: time.Now().Format(time.RFC3339), ExpirationDate: expirationDate, Description: description, CredentialSubject: subject, Proof: proof}
}

func CreateSubjectInfo() *SubjectInfo {
	return &SubjectInfo{}
}

func CreateVCProof() *VCProof {
	return &VCProof{}
}

func CreateResolutionOptions() *ResolutionOptions {
	return &ResolutionOptions{}
}

func GenerateTestPrivKey() *ecdsa.PrivateKey {
	privKey, _ := secp256k1.ToECDSAPrivateKey([]byte{165, 190, 153, 12, 246, 178, 211, 170, 147, 144, 51, 73, 48, 27, 20, 79, 61, 110, 201, 118, 99, 219, 50, 252, 135, 12, 107, 237, 245, 95, 170, 17})
	return privKey
}

func GenerateTestDIDDocument() *DIDDocument {
	document := CreateDIDDocument()
	document.Context = append(document.Context, "https://w3id.org/did/v1")
	document.Context = append(document.Context, "https://ns.did.ai/suites/secp256k1-2019/v1/")
	document.ID = "did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo"
	document.Created = "2022-03-31T12:53:19-07:00"
	document.Updated = "2022-03-31T12:53:19-07:00"
	document.Version = 1
	document.VerificationMethod = append(document.VerificationMethod, VerificationMethod{ID: "did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo#verification", MethodType: "EcdsaSecp256k1VerificationKey2019", Controller: "did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo", MultibaseKey: "zR4TQJaWaLA3vvYukULRJoxTsRmqCMsWuEJdDE8CJwRFCUijDGwCBP89xVcWdLRQaEM6b7wD294xCs8byy3CdDMYa"})
	document.Authentication = "did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo#verification"
	return document
}

func GenerateTestSubjectInfo() *SubjectInfo {
	sampleSubject := CreateSubjectInfo()
	sampleSubject.ID = "did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo"
	sampleSubject.Type = make([]string, 0)
	sampleSubject.Type = append(sampleSubject.Type, "sampleType")
	sampleSubject.GivenName = "John"
	sampleSubject.FamilyName = "Jacobs"
	sampleSubject.Gender = "Male"
	sampleSubject.BirthCountry = "Canada"
	sampleSubject.BirthDate = "2022-03-22"
	return sampleSubject
}

func GenerateTestVC() *VerifiableCredential {
	vc := CreateVerifiableCredential()
	vc.Context = append(vc.Context, "https://www.w3.org/2018/credentials/v1")
	vc.Context = append(vc.Context, "https://ns.did.ai/suites/secp256k1-2019/v1/")
	vc.Type = append(vc.Type, "VerifiableCredential")
	vc.Type = append(vc.Type, "PermanentResidentCard")
	vc.Issuer = "did:metablox:sampleIssuer"
	vc.ExpirationDate = "2032-03-31T12:53:19-07:00"
	vc.Description = "Government of Example Permanent Resident Card"
	vcProof := CreateVCProof()
	vcProof.Type = "EcdsaSecp256k1Signature2019"
	vcProof.VerificationMethod = "did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo#verification"
	vcProof.JWSSignature = "eyJhbGciOiJFUzI1NiJ9..b79nsPjFxYqE0Wta211yA7Rj-MtxMfHsG9dE7V7DGqrK-kMa66d7yjJ0lunAnIUCL7RO55NZ_OuWN-3NK_0J_w"
	vcProof.Created = "2022-03-31T12:53:19-07:00"
	vcProof.ProofPurpose = "Authentication"
	vc.Proof = *vcProof
	return vc
}
