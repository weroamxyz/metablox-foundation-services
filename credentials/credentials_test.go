package credentials

import (
	"errors"
	"testing"

	"github.com/metabloxDID/dao"
	"github.com/metabloxDID/models"
	"github.com/metabloxDID/settings"
	"github.com/stretchr/testify/assert"
)

func TestCreateMiningLicenseVC(t *testing.T) {
	err := settings.Init()
	assert.Nil(t, err)
	err = dao.TestDBInit()
	assert.Nil(t, err)
	t.Cleanup(dao.Close)
	t.Cleanup(dao.DeleteTestCredentialsTable)
	t.Cleanup(dao.DeleteTestMiningLicenseTable)

	err = dao.CreateTestCredentialsTable()
	assert.Nil(t, err)
	err = dao.CreateTestMiningLicenseTable()
	assert.Nil(t, err)

	issuerDocument := models.GenerateTestDIDDocument()
	issuerDocument.ID = "did:metablox:sampleIssuer"
	miningLicenseInfo := models.GenerateTestMiningLicenseInfo()
	issuerPrivKey := models.GenerateTestPrivKey()
	vc, err := CreateMiningLicenseVC(issuerDocument, miningLicenseInfo, issuerPrivKey)
	assert.Nil(t, err)
	sampleVC := models.GenerateTestMiningLicenseVC()
	assert.Equal(t, sampleVC.Context, vc.Context)
	assert.Equal(t, sampleVC.Type, vc.Type)
	assert.Equal(t, sampleVC.Issuer, vc.Issuer)
	assert.Equal(t, sampleVC.Description, vc.Description)
	assert.Equal(t, sampleVC.Proof.Type, vc.Proof.Type)
	assert.Equal(t, sampleVC.Proof.ProofPurpose, vc.Proof.ProofPurpose)
	assert.Equal(t, sampleVC.Proof.VerificationMethod, vc.Proof.VerificationMethod)
	_, err = CreateMiningLicenseVC(issuerDocument, miningLicenseInfo, issuerPrivKey)
	assert.Equal(t, errors.New("mining license vc already exists for user"), err)
}

func TestCreateWifiAccessVC(t *testing.T) {
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

	issuerDocument := models.GenerateTestDIDDocument()
	issuerDocument.ID = "did:metablox:sampleIssuer"
	wifiAccessInfo := models.GenerateTestWifiAccessInfo()
	issuerPrivKey := models.GenerateTestPrivKey()
	vc, err := CreateWifiAccessVC(issuerDocument, wifiAccessInfo, issuerPrivKey)
	assert.Nil(t, err)
	sampleVC := models.GenerateTestWifiAccessVC()
	assert.Equal(t, sampleVC.Context, vc.Context)
	assert.Equal(t, sampleVC.Type, vc.Type)
	assert.Equal(t, sampleVC.Issuer, vc.Issuer)
	assert.Equal(t, sampleVC.Description, vc.Description)
	assert.Equal(t, sampleVC.Proof.Type, vc.Proof.Type)
	assert.Equal(t, sampleVC.Proof.ProofPurpose, vc.Proof.ProofPurpose)
	assert.Equal(t, sampleVC.Proof.VerificationMethod, vc.Proof.VerificationMethod)
}

func TestRenewVC(t *testing.T) {
	err := settings.Init()
	assert.Nil(t, err)
	err = dao.TestDBInit()
	assert.Nil(t, err)
	t.Cleanup(dao.Close)
	t.Cleanup(dao.DeleteTestCredentialsTable)

	err = dao.CreateTestCredentialsTable()
	assert.Nil(t, err)

	vc := models.GenerateTestWifiAccessVC()
	vc.ExpirationDate = "2022-03-31T12:53:19-07:00"
	err = dao.InsertSampleIntoCredentials(vc)
	assert.Nil(t, err)

	err = RenewVC(vc)
	assert.Nil(t, err)
	assert.Equal(t, "2023-03-31T12:53:19-07:00", vc.ExpirationDate)

	dbVC, err := dao.RetrieveSampleFromCredentials("1")
	assert.Nil(t, err)
	assert.Equal(t, "2023-03-31 12:53:19", dbVC.ExpirationDate)
}

func TestRevokeVC(t *testing.T) {
	err := settings.Init()
	assert.Nil(t, err)
	err = dao.TestDBInit()
	assert.Nil(t, err)
	t.Cleanup(dao.Close)
	t.Cleanup(dao.DeleteTestCredentialsTable)

	err = dao.CreateTestCredentialsTable()
	assert.Nil(t, err)

	vc := models.GenerateTestWifiAccessVC()
	err = dao.InsertSampleIntoCredentials(vc)
	assert.Nil(t, err)

	err = RevokeVC(vc)
	assert.Nil(t, err)
	assert.True(t, vc.Revoked)

	dbVC, err := dao.RetrieveSampleFromCredentials("1")
	assert.Nil(t, err)
	assert.True(t, dbVC.Revoked)
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
