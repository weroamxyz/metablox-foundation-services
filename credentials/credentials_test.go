package credentials

import (
	"errors"
	"testing"

	"github.com/metabloxDID/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateVC(t *testing.T) {
	issuerDocument := models.GenerateTestDIDDocument()
	issuerDocument.ID = "did:metablox:sampleIssuer"
	subjectInfo := models.GenerateTestSubjectInfo()
	issuerPrivKey := models.GenerateTestPrivKey()
	vc, err := CreateVC(issuerDocument, subjectInfo, issuerPrivKey)
	assert.Nil(t, err)
	sampleVC := models.GenerateTestVC()
	assert.Equal(t, sampleVC.Context, vc.Context)
	assert.Equal(t, sampleVC.Type, vc.Type)
	assert.Equal(t, sampleVC.Issuer, vc.Issuer)
	assert.Equal(t, sampleVC.Description, vc.Description)
	assert.Equal(t, sampleVC.Proof.Type, vc.Proof.Type)
	assert.Equal(t, sampleVC.Proof.ProofPurpose, vc.Proof.ProofPurpose)
	assert.Equal(t, sampleVC.Proof.VerificationMethod, vc.Proof.VerificationMethod)
}

func TestVerifyVC(t *testing.T) {
	vc := models.GenerateTestVC()
	issuerDocument := models.GenerateTestDIDDocument()

	success, err := VerifyVCSecp256k1(vc, issuerDocument.VerificationMethod[0])
	assert.Nil(t, err)
	assert.True(t, success)
	vc.Type = append(vc.Type, "Modified")
	success, err = VerifyVCSecp256k1(vc, issuerDocument.VerificationMethod[0])
	assert.Equal(t, errors.New("square/go-jose: error in cryptographic primitive"), err)
	assert.False(t, success)
}
