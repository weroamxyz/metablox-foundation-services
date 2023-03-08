package controller

import (
	"errors"
	"github.com/MetaBloxIO/did-sdk-go"
	"github.com/MetaBloxIO/metablox-foundation-services/contract"
	"github.com/MetaBloxIO/metablox-foundation-services/service"
	"github.com/gin-gonic/gin"
)

// IssueWifiVC issue a Wi-Fi access credential using inputted WifiAccessInfo, or return the credential that already exists for the DID in the input
func IssueWifiVC(c *gin.Context) (*did.VerifiableCredential, error) {
	didString := did.IssuerDID

	wifiInfo := did.CreateWifiAccessInfo()

	if err := c.BindJSON(&wifiInfo); err != nil {
		return nil, err
	}

	opts := did.CreateResolutionOptions()
	resolutionMeta, issuerDocument, _ := did.Resolve(didString, opts, contract.GetRegistry())
	if resolutionMeta.Error != "" {
		return nil, errors.New(resolutionMeta.Error)
	}

	newVC, err := service.CreateWifiAccessVC(issuerDocument, wifiInfo, did.IssuerPrivateKey)
	if err != nil {
		return nil, err
	}

	//TODO: currently does nothing
	err = contract.CreateVC(newVC, c.Param("did"), did.IssuerPrivateKey)
	if err != nil {
		return nil, err
	}

	return newVC, nil
}

// IssueMiningVC issue a mining license credential using inputted MiningLicenseInfo, or return the credential that already exists for the DID in the input
func IssueMiningVC(c *gin.Context) (*did.VerifiableCredential, error) {
	didString := did.IssuerDID

	miningInfo := did.CreateMiningLicenseInfo()

	if err := c.BindJSON(&miningInfo); err != nil {
		return nil, err
	}

	opts := did.CreateResolutionOptions()
	resolutionMeta, issuerDocument, _ := did.Resolve(didString, opts, contract.GetRegistry())
	if resolutionMeta.Error != "" {
		return nil, errors.New(resolutionMeta.Error)
	}

	newVC, err := service.CreateMiningLicenseVC(issuerDocument, miningInfo, did.IssuerPrivateKey)
	if err != nil {
		return nil, err
	}

	err = contract.CreateVC(newVC, c.Param("did"), did.IssuerPrivateKey)
	if err != nil {
		return nil, err
	}

	return newVC, nil
}
