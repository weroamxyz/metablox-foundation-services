package credentials

import (
	"testing"

	"github.com/metabloxDID/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateAndVerifyVC(t *testing.T) {
	document := models.GenerateTestDIDDocument()
	document.ID = "did:metablox:sampleIssuer"
	subjectInfo := models.GenerateTestSubjectInfo()
	privKey := models.GenerateTestPrivKey()
	vc, err := CreateVC(document, subjectInfo, privKey)
	assert.Nil(t, err)
	sampleVC := models.GenerateTestVC()
	assert.Equal(t, sampleVC.Context, vc.Context)
	assert.Equal(t, sampleVC.Type, vc.Type)
	assert.Equal(t, sampleVC.Issuer, vc.Issuer)
	assert.Equal(t, sampleVC.Description, vc.Description)
	assert.Equal(t, sampleVC.Proof.Type, vc.Proof.Type)
	assert.Equal(t, sampleVC.Proof.JWSSignature, sampleVC.Proof.JWSSignature)
	assert.Equal(t, sampleVC.Proof.ProofPurpose, vc.Proof.ProofPurpose)
	assert.Equal(t, sampleVC.Proof.VerificationMethod, sampleVC.Proof.VerificationMethod)
	success, err := VerifyVCSecp256k1(vc, document.VerificationMethod[0])
	assert.Nil(t, err)
	assert.True(t, success)
}
