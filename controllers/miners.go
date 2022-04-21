package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/metabloxDID/dao"
	"github.com/metabloxDID/did"
	"github.com/metabloxDID/errval"
	"github.com/metabloxDID/models"
)

func GetMinerList(c *gin.Context) ([]models.MinerInfo, error) {
	didString := "did:metablox:" + c.Param("did")
	authenticationInfo := models.CreateAuthenticationInfo()
	err := c.BindJSON(authenticationInfo)

	if err != nil {
		return nil, err
	}

	err = CheckNonce(c.ClientIP(), authenticationInfo.Nonce)
	if err != nil {
		return nil, err
	}

	opts := models.CreateResolutionOptions()
	resolutionMeta, doc, _ := did.Resolve(didString, opts)
	if resolutionMeta.Error != "" {
		return nil, errors.New(resolutionMeta.Error)
	}
	success, err := did.AuthenticateDocumentHolder(doc, authenticationInfo.Signature, authenticationInfo.Nonce)
	if err != nil {
		return nil, err
	}
	if !success {
		return nil, errval.ErrAuthFailed
	}
	minerList, err := dao.GetMinerList()
	if err != nil {
		return nil, err
	}

	return minerList, nil
}
