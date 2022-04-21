package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/metabloxDID/contract"
	"github.com/metabloxDID/did"
	"github.com/metabloxDID/errval"
	"github.com/metabloxDID/models"
)

func SendDocToRegistry(c *gin.Context) error {
	document := models.CreateDIDDocument()
	if err := c.BindJSON(document); err != nil {
		return err
	}

	_, valid := did.PrepareDID(document.ID)
	if !valid {
		return errval.ErrDIDFormat
	}

	docBytes := [32]byte{}
	copy(docBytes[:], did.ConvertDocToBytes(*document))
	err := contract.UploadDocument(docBytes)
	if err != nil {
		return err
	}

	return nil
}
