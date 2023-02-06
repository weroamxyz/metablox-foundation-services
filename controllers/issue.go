package controllers

import (
	"errors"
	"github.com/MetaBloxIO/metablox-foundation-services/contract"
	"github.com/MetaBloxIO/metablox-foundation-services/did"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/gin-gonic/gin"
)

// IssueWifiVC issue a Wi-Fi access credential using inputted WifiAccessInfo, or return the credential that already exists for the DID in the input
func IssueWifiVC(c *gin.Context) (*models.VerifiableCredential, error) {
	didString := did.IssuerDID

	wifiInfo := models.CreateWifiAccessInfo()

	if err := c.BindJSON(&wifiInfo); err != nil {
		return nil, err
	}

	opts := models.CreateResolutionOptions()
	resolutionMeta, issuerDocument, _ := did.Resolve(didString, opts)
	if resolutionMeta.Error != "" {
		return nil, errors.New(resolutionMeta.Error)
	}

	newVC, err := did.CreateWifiAccessVC(issuerDocument, wifiInfo, did.IssuerPrivateKey)
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
