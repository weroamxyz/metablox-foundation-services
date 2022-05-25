package controllers

import (
	"errors"

	"github.com/MetaBloxIO/metablox-foundation-services/contract"
	"github.com/MetaBloxIO/metablox-foundation-services/credentials"
	"github.com/MetaBloxIO/metablox-foundation-services/did"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/gin-gonic/gin"
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

	newVC, err := credentials.CreateWifiAccessVC(issuerDocument, wifiInfo, credentials.IssuerPrivateKey)
	if err != nil {
		return nil, err
	}

	err = contract.CreateVC(newVC, c.Param("did"), credentials.IssuerPrivateKey)
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

	newVC, err := credentials.CreateMiningLicenseVC(issuerDocument, miningInfo, credentials.IssuerPrivateKey)
	if err != nil {
		return nil, err
	}

	err = contract.CreateVC(newVC, c.Param("did"), credentials.IssuerPrivateKey)
	if err != nil {
		return nil, err
	}

	return newVC, nil
}

func IssueStakingVC(c *gin.Context) (*models.VerifiableCredential, error) {
	didString := credentials.IssuerDID

	stakingInfo := models.CreateStakingVCInfo()

	if err := c.BindJSON(&stakingInfo); err != nil {
		return nil, err
	}

	opts := models.CreateResolutionOptions()
	resolutionMeta, issuerDocument, _ := did.Resolve(didString, opts)
	if resolutionMeta.Error != "" {
		return nil, errors.New(resolutionMeta.Error)
	}

	newVC, err := credentials.CreateStakingVC(issuerDocument, stakingInfo, credentials.IssuerPrivateKey)
	if err != nil {
		return nil, err
	}

	err = contract.CreateVC(newVC, c.Param("did"), credentials.IssuerPrivateKey)
	if err != nil {
		return nil, err
	}

	return newVC, nil
}
