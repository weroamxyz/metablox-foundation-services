package dao

import (
	"errors"
	"net/mail"

	"github.com/MetaBloxIO/metablox-foundation-services/models"
)

func InsertSubmittedEmail(req *models.EmailSubmission) error {

	if len(req.Emails) == 0 {
		return errors.New("Empty Email list")
	}

	for _, email := range req.Emails {
		_, err := mail.ParseAddress(email)
		if err != nil {
			continue
		}

		sqlStr := `INSERT IGNORE INTO email_collect ( email, DID, SN )
			VALUES( ?, ?, ? )`
		_, err = SqlDB.Exec(sqlStr, email, req.DID, req.SN)
		if err != nil {
			return err
		}
	}

	return nil
}
