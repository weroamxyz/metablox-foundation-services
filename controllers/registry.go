package controllers

import (
	"github.com/ethereum/go-ethereum/crypto"
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
	//privateKey := models.GenerateTestPrivKey()
	privateKey, _ := crypto.HexToECDSA("fdebd2c79a17bbea3f69b6ec146bc49b968a63bd24ec342e1bd22830d13f2687")

	_, valid := did.PrepareDID(document.ID)
	if !valid {
		return errval.ErrDIDFormat
	}

	/*docBytes := [32]byte{}
	copy(docBytes[:], did.ConvertDocToBytes(*document))*/
	err := contract.UploadDocument(document, privateKey)
	if err != nil {
		return err
	}

	return nil
}
