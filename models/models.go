package models

type DIDDocument struct {
	Context            string               `json:"@context"`
	ID                 string               `json:"id"`
	Created            string               `json:"created"`
	Updated            string               `json:"updated"`
	Version            int                  `json:"version"`
	VerificationMethod []VerificationMethod `json:"verificationMethod"`
	Authentication     string               `json:"authentication"`
}

type VerificationMethod struct {
	ID         string `json:"id"`
	MethodType string `json:"type"`
	Controller string `json:"controller"`
	Key        string `json:"publicKeyHex"`
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

func CreateDIDDocument() *DIDDocument {
	return &DIDDocument{}
}
