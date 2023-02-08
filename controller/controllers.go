package controller

import (
	"github.com/MetaBloxIO/metablox-foundation-services/did"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

// used to map ip addresses to their assigned nonces
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

// Get a nonce for the user, which will be needed when creating presentations. Lasts 1 minute.
func GenerateNonceHandler(c *gin.Context) {
	nonce := CreateNonce(c.ClientIP())
	ResponseSuccess(c, nonce)
}

// get public key of the issuer
func GetIssuerPublicKeyHandler(c *gin.Context) {
	ResponseSuccess(c, crypto.FromECDSAPub(&did.IssuerPrivateKey.PublicKey))
}

func RegisterDIDHandler(c *gin.Context) {
	hash, err := RegisterDIDForUser(c)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}
	ResponseSuccessWithMsgAndData(c, "DID has been successfully registered to registry", hash)
}
