package models

import (
	"errors"
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
	Expires      string `json:"expires"`
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

func CreateSubjectInfo() *SubjectInfo {
	return &SubjectInfo{}
}

func CreateVCProof() *VCProof {
	return &VCProof{}
}

func CreateResolutionOptions() *ResolutionOptions {
	return &ResolutionOptions{}
}
