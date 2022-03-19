package models

type DIDDocument struct {
	Context            string
	Created            string
	Updated            string
	Version            int
	ID                 string
	VerificationMethod []VerificationMethod
	Authentication     string
}

type VerificationMethod struct {
	ID         string
	MethodType string
	Controller string
	Key        string
}
