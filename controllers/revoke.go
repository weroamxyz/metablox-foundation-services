package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/MetaBloxIO/metablox-foundation-services/contract"
	"github.com/MetaBloxIO/metablox-foundation-services/credentials"
	"github.com/MetaBloxIO/metablox-foundation-services/errval"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/MetaBloxIO/metablox-foundation-services/presentations"
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
