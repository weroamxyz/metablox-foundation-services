package controllers

import (
	"encoding/json"
	"github.com/MetaBloxIO/metablox-foundation-services/comm/regutil"
	"github.com/MetaBloxIO/metablox-foundation-services/contract"
	"github.com/MetaBloxIO/metablox-foundation-services/credentials"
	"github.com/MetaBloxIO/metablox-foundation-services/did"
	"github.com/MetaBloxIO/metablox-foundation-services/errval"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

func SendDocToRegistry(c *gin.Context) error {
	document := models.CreateDIDDocument()
	if err := c.BindJSON(document); err != nil {
		return err
	}

	splitString, valid := did.PrepareDID(document.ID)
	if !valid {
		return errval.ErrDIDFormat
	}

	err := contract.UploadDocument(document, splitString[2], credentials.IssuerPrivateKey)
	if err != nil {
		return err
	}

	return nil
}

func RegisterDIDForUser(c *gin.Context) (map[string]interface{}, error) {
	// 1.new param instance
	register := models.NewRegisterDID()
	if err := c.BindJSON(register); err != nil {
		return nil, err
	}

	jsonStr, err := json.MarshalIndent(register, "", "\t")
	logger.Infoln("received request params:", string(jsonStr))

	// 2.check did format
	_, valid := did.PrepareDID(register.Did)
	if !valid {
		return nil, errval.ErrDIDFormat
	}
	// 3.check account format
	if !regutil.IsETHAddress(register.Account) {
		return nil, errval.ErrETHAddress
	}
	// 4. handle biz logic
	tx, err := contract.RegisterDID(register, credentials.IssuerPrivateKey)
	if err != nil {
		return nil, err
	}
	// 5. wrap response
	data := make(map[string]interface{})
	data["txHash"] = tx.Hash().Hex()
	return data, nil
}
