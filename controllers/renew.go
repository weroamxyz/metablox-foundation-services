package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/metabloxDID/contract"
	"github.com/metabloxDID/credentials"
	"github.com/metabloxDID/did"
	"github.com/metabloxDID/errval"
	"github.com/metabloxDID/models"
	"github.com/metabloxDID/presentations"
)

func RenewVC(c *gin.Context) (*models.VerifiableCredential, error) {
	didString := "did:metablox:" + c.Param("did")

	var input struct {
		AuthenticationInfo *models.AuthenticationInfo
		Presentation       *models.VerifiablePresentation
	}

	if err := c.BindJSON(&input); err != nil {
		return nil, err
	}

	err := CheckNonce(c.ClientIP(), input.AuthenticationInfo.Nonce)
	if err != nil {
		return nil, err
	}
	defer DeleteNonce(c.ClientIP()) //we re-use the nonce for the presentation, so only delete the nonce after the controller is completely finished

	err = CheckNonce(c.ClientIP(), input.Presentation.Proof.Nonce)
	if err != nil {
		return nil, err
	}

	if input.Presentation.VerifiableCredential[0].Issuer != didString {
		return nil, errval.ErrDIDNotIssuer
	}

	opts := models.CreateResolutionOptions()
	resolutionMeta, issuerDocument, _ := did.Resolve(didString, opts)
	if resolutionMeta.Error != "" {
		return nil, errors.New(resolutionMeta.Error)
	}

	success, err := did.AuthenticateDocumentHolder(issuerDocument, input.AuthenticationInfo.Signature, input.AuthenticationInfo.Nonce)
	if err != nil {
		return nil, err
	}
	if !success {
		return nil, errval.ErrAuthFailed
	}

	success, err = presentations.VerifyVP(input.Presentation)
	if err != nil {
		return nil, err
	}

	if !success {
		return nil, errval.ErrVerifyPresent
	}

	err = credentials.RenewVC(&input.Presentation.VerifiableCredential[0])
	if err != nil {
		return nil, err
	}

	vcBytes := [32]byte{}
	copy(vcBytes[:], credentials.ConvertVCToBytes(input.Presentation.VerifiableCredential[0]))
	err = contract.RenewVC(vcBytes)
	if err != nil {
		return nil, err
	}

	return &input.Presentation.VerifiableCredential[0], nil
}
