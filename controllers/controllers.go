package controllers

import (
	"crypto/ecdsa"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/metabloxDID/contract"
	"github.com/metabloxDID/key"
	"github.com/metabloxDID/models"
)

var NonceLookup map[string]string
var ValidIssuers []string
var issuerPrivateKey *ecdsa.PrivateKey

func InitializeValues() error {
	var err error
	NonceLookup = make(map[string]string)
	ValidIssuers = []string{"did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo"}
	issuerPrivateKey, err = key.GetIssuerPrivateKey()
	if err != nil {
		return err
	}
	testKey, _ := crypto.HexToECDSA("fdebd2c79a17bbea3f69b6ec146bc49b968a63bd24ec342e1bd22830d13f2687")
	err = contract.TestSignatures(testKey, "secret message")
	if err != nil {
		fmt.Println(err)
	}
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
