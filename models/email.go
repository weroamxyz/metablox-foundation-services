package models

type EmailSubmission struct {
	SN     string   `json:"sn"`
	DID    string   `json:"did"`
	Emails []string `json:"emails"`
}
