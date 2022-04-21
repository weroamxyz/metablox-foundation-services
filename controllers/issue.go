package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/metabloxDID/contract"
	"github.com/metabloxDID/credentials"
	"github.com/metabloxDID/did"
	"github.com/metabloxDID/errval"
	"github.com/metabloxDID/models"
)

func CheckIfValidIssuer(did string) bool {
	for _, issuer := range ValidIssuers {
		if did == issuer {
			return true
		}
	}
	return false
}

func IssueWifiVC(c *gin.Context) (*models.VerifiableCredential, error) {
	didString := "did:metablox:" + c.Param("did")
	if !CheckIfValidIssuer(didString) {
		return nil, errval.ErrInvalidIssuer
	}
	var input struct {
		AuthenticationInfo *models.AuthenticationInfo
		WifiAccessInfo     *models.WifiAccessInfo
	}

	if err := c.BindJSON(&input); err != nil {
		return nil, err
	}

	err := CheckNonce(c.ClientIP(), input.AuthenticationInfo.Nonce)
	if err != nil {
		return nil, err
	}
	DeleteNonce(c.ClientIP())

	opts := models.CreateResolutionOptions()
	resolutionMeta, issuerDocument, _ := did.Resolve(didString, opts)
	if resolutionMeta.Error != "" {
		return nil, err
	}

	success, err := did.AuthenticateDocumentHolder(issuerDocument, input.AuthenticationInfo.Signature, input.AuthenticationInfo.Nonce)
	if err != nil {
		return nil, err
	}
	if !success {
		return nil, errval.ErrAuthFailed
	}

	newVC, err := credentials.CreateWifiAccessVC(issuerDocument, input.WifiAccessInfo, issuerPrivateKey)
	if err != nil {
		return nil, err
	}

	vcBytes := [32]byte{}
	copy(vcBytes[:], credentials.ConvertVCToBytes(*newVC))
	err = contract.CreateVC(vcBytes)
	if err != nil {
		return nil, err
	}

	return newVC, nil
}

func IssueMiningVC(c *gin.Context) (*models.VerifiableCredential, error) {
	didString := "did:metablox:" + c.Param("did")
	if !CheckIfValidIssuer(didString) {
		return nil, errval.ErrInvalidIssuer
	}
	var input struct {
		AuthenticationInfo *models.AuthenticationInfo
		MiningLicenseInfo  *models.MiningLicenseInfo
	}

	if err := c.BindJSON(&input); err != nil {
		return nil, err
	}

	err := CheckNonce(c.ClientIP(), input.AuthenticationInfo.Nonce)
	if err != nil {
		return nil, err
	}
	DeleteNonce(c.ClientIP())

	opts := models.CreateResolutionOptions()
	resolutionMeta, issuerDocument, _ := did.Resolve(didString, opts)
	if resolutionMeta.Error != "" {
		return nil, err
	}

	success, err := did.AuthenticateDocumentHolder(issuerDocument, input.AuthenticationInfo.Signature, input.AuthenticationInfo.Nonce)
	if err != nil {
		return nil, err
	}
	if !success {
		return nil, errval.ErrAuthFailed
	}

	newVC, err := credentials.CreateMiningLicenseVC(issuerDocument, input.MiningLicenseInfo, issuerPrivateKey)
	if err != nil {
		return nil, err
	}

	vcBytes := [32]byte{}
	copy(vcBytes[:], credentials.ConvertVCToBytes(*newVC))
	err = contract.CreateVC(vcBytes)
	if err != nil {
		return nil, err
	}

	return newVC, nil
}
