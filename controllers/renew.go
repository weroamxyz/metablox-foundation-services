package controllers

import (
	"github.com/MetaBloxIO/metablox-foundation-services/contract"
	"github.com/MetaBloxIO/metablox-foundation-services/credentials"
	"github.com/MetaBloxIO/metablox-foundation-services/errval"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/MetaBloxIO/metablox-foundation-services/presentations"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

func RenewVC(c *gin.Context) (*models.VerifiableCredential, error) {
	var input *struct {
		Presentation    models.VerifiablePresentation
		PublicKeyString []byte
	}

	if err := c.BindJSON(&input); err != nil {
		return nil, err
	}

	pubKey, err := crypto.UnmarshalPubkey(input.PublicKeyString)
	if err != nil {
		return nil, err
	}

	err = CheckNonce(c.ClientIP(), input.Presentation.Proof.Nonce)
	if err != nil {
		return nil, err
	}

	DeleteNonce(c.ClientIP())
	for i, vc := range input.Presentation.VerifiableCredential {
		ConvertCredentialSubject(&vc)
		input.Presentation.VerifiableCredential[i] = vc
	}

	success, err := presentations.VerifyVP(&input.Presentation, pubKey, &issuerPrivateKey.PublicKey)
	if err != nil {
		return nil, err
	}

	if !success {
		return nil, errval.ErrVerifyPresent
	}

	err = credentials.RenewVC(&input.Presentation.VerifiableCredential[0], issuerPrivateKey)
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
