package presentations

import (
	"errors"
	"testing"

	"github.com/metabloxDID/credentials"
	"github.com/metabloxDID/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateVP(t *testing.T) {
	issuerDocument := models.GenerateTestDIDDocument()
	issuerDocument.ID = "did:metablox:sampleIssuer"
	subjectInfo := models.GenerateTestSubjectInfo()
	issuerPrivKey := models.GenerateTestPrivKey()
	vc, err := credentials.CreateVC(issuerDocument, subjectInfo, issuerPrivKey)
	assert.Nil(t, err)
	credentialArray := make([]models.VerifiableCredential, 0)
	credentialArray = append(credentialArray, *vc)
	issuerDocument.ID = "did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo"

	presentation, err := CreatePresentation(credentialArray, *issuerDocument, issuerPrivKey, "sampleNonce")
	assert.Nil(t, err)

	samplePresentation := models.GenerateTestPresentation()
	assert.Equal(t, samplePresentation.Context, presentation.Context)
	assert.Equal(t, samplePresentation.Type, presentation.Type)
	assert.Equal(t, samplePresentation.Holder, presentation.Holder)
	assert.Equal(t, samplePresentation.Proof.Type, presentation.Proof.Type)
	assert.Equal(t, samplePresentation.Proof.ProofPurpose, presentation.Proof.ProofPurpose)
	assert.Equal(t, samplePresentation.Proof.VerificationMethod, presentation.Proof.VerificationMethod)
	assert.Equal(t, samplePresentation.Proof.Nonce, presentation.Proof.Nonce)
}

func TestVerifyVP(t *testing.T) {
	issuerDocument := models.GenerateTestDIDDocument()
	samplePresentation := models.GenerateTestPresentation()

	success, err := VerifyVPSecp256k1(samplePresentation, issuerDocument.VerificationMethod[0], "sampleNonce")
	assert.Nil(t, err)
	assert.True(t, success)
	samplePresentation.Type = append(samplePresentation.Type, "Modified")
	success, err = VerifyVPSecp256k1(samplePresentation, issuerDocument.VerificationMethod[0], "sampleNonce")
	assert.Equal(t, errors.New("square/go-jose: error in cryptographic primitive"), err)
	assert.False(t, success)
	success, err = VerifyVPSecp256k1(samplePresentation, issuerDocument.VerificationMethod[0], "sampleNonce2")
	assert.Nil(t, err)
	assert.False(t, success)
}
