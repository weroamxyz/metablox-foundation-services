package controllers

import (
	"crypto/sha256"
	"net/http"

	"github.com/MetaBloxIO/metablox-foundation-services/contract"
	"github.com/MetaBloxIO/metablox-foundation-services/credentials"
	"github.com/MetaBloxIO/metablox-foundation-services/key"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/MetaBloxIO/metablox-foundation-services/presentations"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

var NonceLookup map[string]string

func InitializeValues() {
	NonceLookup = make(map[string]string)
}

func IssueWifiVCHandler(c *gin.Context) {
	newVC, err := IssueWifiVC(c)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccess(c, newVC)
}

func IssueMiningVCHandler(c *gin.Context) {
	newVC, err := IssueMiningVC(c)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccess(c, newVC)
}

func IssueStakingVCHandler(c *gin.Context) {
	newVC, err := IssueStakingVC(c)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccess(c, newVC)
}

func RenewVCHandler(c *gin.Context) {
	renewedVC, err := RenewVC(c)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccess(c, renewedVC)
}

func RevokeVCHandler(c *gin.Context) {
	revokedVC, err := RevokeVC(c)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccess(c, revokedVC)
}

func SendDocToRegistryHandler(c *gin.Context) {
	err := SendDocToRegistry(c)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccessWithMsg(c, "DID document has been successfully uploaded to registry")
}

func GetMinerListHandler(c *gin.Context) {
	minerList, err := GetMinerList(c)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccess(c, minerList)
}

func GenerateNonceHandler(c *gin.Context) {
	nonce := CreateNonce(c.ClientIP())
	ResponseSuccess(c, nonce)
}

//temporary testing function to generate signatures using the test private key
func GenerateTestSignatures(c *gin.Context) {
	message := c.Param("message")
	privKey := models.GenerateTestPrivKey()
	signature, err := key.CreateJWSSignature(privKey, []byte(message))
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed to create signature: "+err.Error())
		return
	}

	ResponseSuccessWithMsg(c, "Generated signature '"+signature+"'")
}

func AssignIssuer(c *gin.Context) {
	var inputs struct {
		CredentialKey string
		Did           string
	}

	if err := c.BindJSON(&inputs); err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed: "+err.Error())
		return
	}

	err := contract.RegisterVCIssuer(inputs.CredentialKey, inputs.Did, credentials.IssuerPrivateKey)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed: "+err.Error())
		return
	}

	ResponseSuccessWithMsg(c, "Success")
}

func SetVCAttribute(c *gin.Context) {
	var inputs struct {
		CredentialKey string
		FieldName     string
		NewValue      string
	}

	if err := c.BindJSON(&inputs); err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed: "+err.Error())
		return
	}

	err := contract.UpdateVCValue(inputs.CredentialKey, inputs.FieldName, inputs.NewValue, credentials.IssuerPrivateKey)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed: "+err.Error())
		return
	}

	ResponseSuccessWithMsg(c, "Success")
}

func ReadVCChangedEvents(c *gin.Context) {
	var inputs struct {
		CredentialKey string
	}

	if err := c.BindJSON(&inputs); err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed: "+err.Error())
		return
	}

	err := contract.ReadVCChangedEvents(inputs.CredentialKey)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed: "+err.Error())
		return
	}

	ResponseSuccessWithMsg(c, "Success")
}

func GetIssuerPublicKeyHandler(c *gin.Context) {
	ResponseSuccess(c, crypto.FromECDSAPub(&credentials.IssuerPrivateKey.PublicKey))
}

func GenerateTestingPresentationSignatures(c *gin.Context) {
	presentation := models.CreatePresentation()
	c.BindJSON(presentation)
	for i, vc := range presentation.VerifiableCredential {
		vc.Proof.JWSSignature = ""
		ConvertCredentialSubject(&vc)
		hashedVC := sha256.Sum256(credentials.ConvertVCToBytes(vc))

		signatureData, _ := key.CreateJWSSignature(models.GenerateTestPrivKey(), hashedVC[:])
		vc.Proof.JWSSignature = signatureData
		presentation.VerifiableCredential[i] = vc
	}
	presentation.Proof.JWSSignature = ""
	hashedVP := sha256.Sum256(presentations.ConvertVPToBytes(*presentation))

	signatureData, _ := key.CreateJWSSignature(models.GenerateTestPrivKey(), hashedVP[:])
	presentation.Proof.JWSSignature = signatureData
	ResponseSuccess(c, presentation)
}
