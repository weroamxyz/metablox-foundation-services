package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/metabloxDID/contract"
	"github.com/metabloxDID/credentials"
	"github.com/metabloxDID/errval"
	"github.com/metabloxDID/models"
	"github.com/metabloxDID/presentations"
)

func RevokeVC(c *gin.Context) (*models.VerifiableCredential, error) {

	vp := models.CreatePresentation()

	if err := c.BindJSON(&vp); err != nil {
		return nil, err
	}

	err := CheckNonce(c.ClientIP(), vp.Proof.Nonce)
	if err != nil {
		return nil, err
	}

	DeleteNonce(c.ClientIP())

	success, err := presentations.VerifyVP(vp)
	if err != nil {
		return nil, err
	}

	if !success {
		return nil, errval.ErrVerifyPresent
	}

	err = credentials.RevokeVC(&vp.VerifiableCredential[0])
	if err != nil {
		return nil, err
	}

	vcBytes := [32]byte{}
	copy(vcBytes[:], credentials.ConvertVCToBytes(vp.VerifiableCredential[0]))
	err = contract.RevokeVC(vcBytes)
	if err != nil {
		return nil, err
	}

	return &vp.VerifiableCredential[0], nil
}
