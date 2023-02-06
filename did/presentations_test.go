package did

import (
	"testing"

	"github.com/MetaBloxIO/metablox-foundation-services/dao"
	"github.com/MetaBloxIO/metablox-foundation-services/errval"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/MetaBloxIO/metablox-foundation-services/settings"
	"github.com/stretchr/testify/assert"
)

func TestCreateVP(t *testing.T) {
	err := settings.Init()
	assert.Nil(t, err)
	err = dao.TestDBInit()
	assert.Nil(t, err)
	t.Cleanup(dao.Close)
	t.Cleanup(dao.DeleteTestCredentialsTable)
	t.Cleanup(dao.DeleteTestWifiAccessTable)

	err = dao.CreateTestCredentialsTable()
	assert.Nil(t, err)

	err = dao.CreateTestWifiAccessTable()
	assert.Nil(t, err)
	IssuerPrivateKey = models.GenerateTestPrivKey()

	issuerDocument := models.GenerateTestDIDDocument()
	issuerDocument.ID = "did:metablox:sampleIssuer"
	wifiAccessInfo := models.GenerateTestWifiAccessInfo()
	issuerPrivKey := models.GenerateTestPrivKey()
	vc, err := CreateWifiAccessVC(issuerDocument, wifiAccessInfo, issuerPrivKey)
	assert.Nil(t, err)
	credentialArray := make([]models.VerifiableCredential, 0)
	credentialArray = append(credentialArray, *vc)
	issuerDocument.ID = "did:metablox:7rb6LjVKYSEf4LLRqbMQGgdeE8MYXkfS7dhjvJzUckEX"

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
	samplePresentation := models.GenerateTestPresentation()

	success, err := VerifyVPSecp256k1(samplePresentation, &models.GenerateTestPrivKey().PublicKey)
	assert.Nil(t, err)
	assert.True(t, success)
	samplePresentation.Type = append(samplePresentation.Type, "Modified")
	success, err = VerifyVPSecp256k1(samplePresentation, &models.GenerateTestPrivKey().PublicKey)
	assert.Equal(t, errval.ErrJWSAuthentication, err)
	assert.False(t, success)
}
