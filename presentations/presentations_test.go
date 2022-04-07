package presentations

import (
	"errors"
	"testing"

	"github.com/metabloxDID/credentials"
	"github.com/metabloxDID/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateAndVerifyVP(t *testing.T) {
	document := models.GenerateTestDIDDocument()
	document.ID = "did:metablox:sampleIssuer"
	subjectInfo := models.GenerateTestSubjectInfo()
	privKey := models.GenerateTestPrivKey()
	vc, err := credentials.CreateVC(document, subjectInfo, privKey)
	assert.Nil(t, err)
	credentialArray := make([]models.VerifiableCredential, 0)
	credentialArray = append(credentialArray, *vc)
	document.ID = "did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo"

	presentation, err := CreatePresentation(credentialArray, *document, privKey)
	assert.Nil(t, err)

	samplePresentation := models.GenerateTestPresentation()
	assert.Equal(t, samplePresentation.Context, presentation.Context)
	assert.Equal(t, samplePresentation.Type, presentation.Type)
	assert.Equal(t, samplePresentation.Holder, presentation.Holder)
	assert.Equal(t, samplePresentation.Proof.Type, presentation.Proof.Type)
	assert.Equal(t, samplePresentation.Proof.ProofPurpose, presentation.Proof.ProofPurpose)
	assert.Equal(t, samplePresentation.Proof.VerificationMethod, presentation.Proof.VerificationMethod)
	success, err := VerifyVPSecp256k1(presentation, document.VerificationMethod[0])
	assert.Nil(t, err)
	assert.True(t, success)
	presentation.Type = append(presentation.Type, "Modified")
	success, err = VerifyVPSecp256k1(presentation, document.VerificationMethod[0])
	assert.Equal(t, errors.New("square/go-jose: error in cryptographic primitive"), err)
	assert.False(t, success)
}
