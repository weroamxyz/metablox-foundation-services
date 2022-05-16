package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/metabloxDID/contract"
	"github.com/metabloxDID/credentials"
	"github.com/metabloxDID/did"
	"github.com/metabloxDID/models"
)

func IssueWifiVC(c *gin.Context) (*models.VerifiableCredential, error) {
	didString := credentials.IssuerDID

	wifiInfo := models.CreateWifiAccessInfo()

	if err := c.BindJSON(&wifiInfo); err != nil {
		return nil, err
	}

	opts := models.CreateResolutionOptions()
	resolutionMeta, issuerDocument, _ := did.Resolve(didString, opts)
	if resolutionMeta.Error != "" {
		return nil, errors.New(resolutionMeta.Error)
	}

	newVC, err := credentials.CreateWifiAccessVC(issuerDocument, wifiInfo, issuerPrivateKey)
	if err != nil {
		return nil, err
	}

	err = contract.CreateVC(newVC, c.Param("did"), issuerPrivateKey)
	if err != nil {
		return nil, err
	}

	return newVC, nil
}

func IssueMiningVC(c *gin.Context) (*models.VerifiableCredential, error) {
	didString := credentials.IssuerDID

	miningInfo := models.CreateMiningLicenseInfo()

	if err := c.BindJSON(&miningInfo); err != nil {
		return nil, err
	}

	opts := models.CreateResolutionOptions()
	resolutionMeta, issuerDocument, _ := did.Resolve(didString, opts)
	if resolutionMeta.Error != "" {
		return nil, errors.New(resolutionMeta.Error)
	}

	newVC, err := credentials.CreateMiningLicenseVC(issuerDocument, miningInfo, issuerPrivateKey)
	if err != nil {
		return nil, err
	}

	err = contract.CreateVC(newVC, c.Param("did"), issuerPrivateKey)
	if err != nil {
		return nil, err
	}

	return newVC, nil
}
