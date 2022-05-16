package controllers

import (
	"crypto/ecdsa"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/MetaBloxIO/metablox-foundation-services/contract"
	"github.com/MetaBloxIO/metablox-foundation-services/credentials"
	"github.com/MetaBloxIO/metablox-foundation-services/did"
	"github.com/MetaBloxIO/metablox-foundation-services/key"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
)

var NonceLookup map[string]string

var issuerPrivateKey *ecdsa.PrivateKey

func InitializeValues() error {
	var err error
	NonceLookup = make(map[string]string)
	issuerPrivateKey, err = key.GetIssuerPrivateKey()
	if err != nil {
		return err
	}
	credentials.IssuerDID = did.GenerateDIDString(issuerPrivateKey)
	return nil
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

	err := contract.RegisterVCIssuer(inputs.CredentialKey, inputs.Did, issuerPrivateKey)
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

	err := contract.UpdateVCValue(inputs.CredentialKey, inputs.FieldName, inputs.NewValue, issuerPrivateKey)
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
