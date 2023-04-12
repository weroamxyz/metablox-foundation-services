package service

import (
	"github.com/MetaBloxIO/did-sdk-go"
	"github.com/MetaBloxIO/metablox-foundation-services/dao"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/pkg/errors"
)

func HandleEmailSubmission(req *models.EmailSubmission) error {

	_, valid := did.PrepareDID(req.DID)
	if !valid {
		return errors.New("Invalid DID")
	}

	if req.SN == "" {
		return errors.New("Empty SN")
	}

	err := dao.InsertSubmittedEmail(req)

	return err
}
