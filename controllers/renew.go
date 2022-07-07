package controllers

import (
	"github.com/MetaBloxIO/metablox-foundation-services/contract"
	"github.com/MetaBloxIO/metablox-foundation-services/credentials"
	"github.com/MetaBloxIO/metablox-foundation-services/errval"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/MetaBloxIO/metablox-foundation-services/presentations"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

//renew the first credential in the provided presentation. Currently always increase the expiration date by 1 year
func RenewVC(c *gin.Context) (*models.VerifiableCredential, error) {
	input := models.CreatePresentation()

	if err := c.BindJSON(&input); err != nil {
		return nil, err
	}

	err := CheckNonce(c.ClientIP(), input.Proof.Nonce) //presentation must have a valid nonce
	if err != nil {
		return nil, err
	}

	DeleteNonce(c.ClientIP())
	for i, vc := range input.VerifiableCredential {
		ConvertCredentialSubject(&vc)
		input.VerifiableCredential[i] = vc
	}

	success, err := presentations.VerifyVP(input)
	if err != nil {
		logger.Warn(err)
	}

	if !success {
		return nil, errval.ErrVerifyPresent
	}

	err = credentials.RenewVC(&input.VerifiableCredential[0], credentials.IssuerPrivateKey)
	if err != nil {
		return nil, err
	}

	vcBytes := [32]byte{}
	copy(vcBytes[:], credentials.ConvertVCToBytes(input.VerifiableCredential[0]))
	err = contract.RenewVC(vcBytes) //currently does nothing
	if err != nil {
		return nil, err
	}

	return &input.VerifiableCredential[0], nil
}
