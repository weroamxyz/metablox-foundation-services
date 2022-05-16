package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/metabloxDID/contract"
	"github.com/metabloxDID/did"
	"github.com/metabloxDID/errval"
	"github.com/metabloxDID/models"
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
