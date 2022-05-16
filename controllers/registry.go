package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/MetaBloxIO/metablox-foundation-services/contract"
	"github.com/MetaBloxIO/metablox-foundation-services/did"
	"github.com/MetaBloxIO/metablox-foundation-services/errval"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
)

func SendDocToRegistry(c *gin.Context) error {
	/*document := models.CreateDIDDocument()
	if err := c.BindJSON(document); err != nil {
		return err
	}*/

	document := models.GenerateTestDIDDocument()

	splitString, valid := did.PrepareDID(document.ID)
	if !valid {
		return errval.ErrDIDFormat
	}

	err := contract.UploadDocument(document, splitString[2], issuerPrivateKey)
	if err != nil {
		return err
	}

	return nil
}
