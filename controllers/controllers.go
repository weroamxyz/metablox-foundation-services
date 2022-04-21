package controllers

import (
	"crypto/ecdsa"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metabloxDID/contract"
	"github.com/metabloxDID/credentials"
	"github.com/metabloxDID/dao"
	"github.com/metabloxDID/did"
	"github.com/metabloxDID/key"
	"github.com/metabloxDID/models"
	"github.com/metabloxDID/presentations"
	"github.com/mitchellh/mapstructure"
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
	return nil
}

//If user provided correct nonce, then assign them a new one and return that value
func CreateNonce(ip string) string {
	NonceLookup[ip] = time.Now().Format("2006-01-02 15:04:05.999999999 -0700 MST")
	return NonceLookup[ip]
}

//Compare the nonce a user has given with the one they are assigned. Current time must also be within 1 minute of the nonce's value
func CheckNonce(ip, givenNonce string) (bool, error) {
	assignedNonce, found := NonceLookup[ip]
	if !found {
		return false, errors.New("no nonce assigned to user")
	}

	nonceTime, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", assignedNonce)
	if err != nil {
		return false, err
	}
	if nonceTime.Add(time.Minute).Before(time.Now()) {
		delete(NonceLookup, ip) //remove expired nonces
		return false, errors.New("nonce has expired")
	}

	if assignedNonce != givenNonce {
		return false, errors.New("provided nonce is incorrect")
	}

	return true, nil
}

//delete a nonce after it has been successfully used in an operation
func DeleteNonce(ip string) {
	delete(NonceLookup, ip)
}

func CheckIfValidIssuer(did string) bool {
	for _, issuer := range ValidIssuers {
		if did == issuer {
			return true
		}
	}
	return false
}

func IssueWifiVCHandler(c *gin.Context) {
	didString := "did:metablox:" + c.Param("did")
	if !CheckIfValidIssuer(didString) {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Provided did is not a valid issuer")
		return
	}

	var input struct {
		AuthenticationInfo *models.AuthenticationInfo
		WifiAccessInfo     *models.WifiAccessInfo
	}

	if err := c.BindJSON(&input); err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	success, err := CheckNonce(c.ClientIP(), input.AuthenticationInfo.Nonce)
	if !success {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed to verify nonce: "+err.Error())
		return
	}
	DeleteNonce(c.ClientIP())

	opts := models.CreateResolutionOptions()
	resolutionMeta, issuerDocument, _ := did.Resolve(didString, opts)
	if resolutionMeta.Error != "" {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed to resolve did '"+didString+"': "+resolutionMeta.Error)
		return
	}

	success, err = did.AuthenticateDocumentHolder(issuerDocument, input.AuthenticationInfo.Signature, input.AuthenticationInfo.Nonce)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error authenticating signature: "+err.Error())
		return
	}
	if !success {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Authentication failed")
		return
	}

	newVC, err := credentials.CreateWifiAccessVC(issuerDocument, input.WifiAccessInfo, issuerPrivateKey)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error creating wifi access VC: "+err.Error())
		return
	}

	vcBytes := [32]byte{}
	copy(vcBytes[:], credentials.ConvertVCToBytes(*newVC))
	err = contract.CreateVC(vcBytes)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error registering wifi access VC: "+err.Error())
		return
	}

	ResponseSuccess(c, newVC)
}

func IssueMiningVCHandler(c *gin.Context) {
	didString := "did:metablox:" + c.Param("did")
	if !CheckIfValidIssuer(didString) {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Provided did is not a valid issuer")
		return
	}
	var input struct {
		AuthenticationInfo *models.AuthenticationInfo
		MiningLicenseInfo  *models.MiningLicenseInfo
	}

	if err := c.BindJSON(&input); err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	success, err := CheckNonce(c.ClientIP(), input.AuthenticationInfo.Nonce)
	if !success {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed to verify nonce: "+err.Error())
		return
	}
	DeleteNonce(c.ClientIP())

	opts := models.CreateResolutionOptions()
	resolutionMeta, issuerDocument, _ := did.Resolve(didString, opts)
	if resolutionMeta.Error != "" {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed to resolve did '"+didString+"': "+resolutionMeta.Error)
		return
	}

	success, err = did.AuthenticateDocumentHolder(issuerDocument, input.AuthenticationInfo.Signature, input.AuthenticationInfo.Nonce)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error authenticating signature: "+err.Error())
		return
	}
	if !success {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Authentication failed")
		return
	}

	newVC, err := credentials.CreateMiningLicenseVC(issuerDocument, input.MiningLicenseInfo, issuerPrivateKey)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error creating mining license VC: "+err.Error())
		return
	}

	vcBytes := [32]byte{}
	copy(vcBytes[:], credentials.ConvertVCToBytes(*newVC))
	err = contract.CreateVC(vcBytes)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error registering mining license VC: "+err.Error())
		return
	}

	ResponseSuccess(c, newVC)
}

func RenewVCHandler(c *gin.Context) {
	didString := "did:metablox:" + c.Param("did")

	var input struct {
		AuthenticationInfo *models.AuthenticationInfo
		Presentation       *models.VerifiablePresentation
	}

	if err := c.BindJSON(&input); err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	success, err := CheckNonce(c.ClientIP(), input.AuthenticationInfo.Nonce)
	if !success {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed to verify nonce: "+err.Error())
		return
	}
	defer DeleteNonce(c.ClientIP()) //we re-use the nonce for the presentation, so only delete the nonce after the controller is completely finished

	success, err = CheckNonce(c.ClientIP(), input.Presentation.Proof.Nonce)
	if !success {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed to verify presentation nonce: "+err.Error())
		return
	}

	if input.Presentation.VerifiableCredential[0].Issuer != didString {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Provided did does not match issuer of credential")
		return
	}

	opts := models.CreateResolutionOptions()
	resolutionMeta, issuerDocument, _ := did.Resolve(didString, opts)
	if resolutionMeta.Error != "" {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed to resolve did '"+didString+"': "+resolutionMeta.Error)
		return
	}

	success, err = did.AuthenticateDocumentHolder(issuerDocument, input.AuthenticationInfo.Signature, input.AuthenticationInfo.Nonce)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error authenticating signature: "+err.Error())
		return
	}
	if !success {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Authentication failed")
		return
	}

	success, err = presentations.VerifyVP(input.Presentation)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error verifying presentation: "+err.Error())
		return
	}

	if !success {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed to verify presentation")
		return
	}

	err = credentials.RenewVC(&input.Presentation.VerifiableCredential[0])
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed to renew credential: "+err.Error())
	}

	vcBytes := [32]byte{}
	copy(vcBytes[:], credentials.ConvertVCToBytes(input.Presentation.VerifiableCredential[0]))
	err = contract.RenewVC(vcBytes)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed to renew credential in registry: "+err.Error())
		return
	}

	ResponseSuccess(c, input.Presentation.VerifiableCredential[0])
}

func RevokeVCHandler(c *gin.Context) {
	didString := "did:metablox:" + c.Param("did")
	var input struct {
		AuthenticationInfo *models.AuthenticationInfo
		Presentation       *models.VerifiablePresentation
	}

	if err := c.BindJSON(&input); err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	success, err := CheckNonce(c.ClientIP(), input.AuthenticationInfo.Nonce)
	if !success {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed to verify nonce: "+err.Error())
		return
	}
	defer DeleteNonce(c.ClientIP()) //we re-use the nonce for the presentation, so only delete the nonce after the controller is finished

	success, err = CheckNonce(c.ClientIP(), input.Presentation.Proof.Nonce)
	if !success {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed to verify presentation nonce: "+err.Error())
		return
	}

	if input.Presentation.VerifiableCredential[0].Issuer != didString {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Provided did does not match issuer of credential")
		return
	}

	opts := models.CreateResolutionOptions()
	resolutionMeta, issuerDocument, _ := did.Resolve(didString, opts)
	if resolutionMeta.Error != "" {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed to resolve did '"+didString+"': "+resolutionMeta.Error)
		return
	}

	success, err = did.AuthenticateDocumentHolder(issuerDocument, input.AuthenticationInfo.Signature, input.AuthenticationInfo.Nonce)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error authenticating signature: "+err.Error())
		return
	}
	if !success {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Authentication failed")
		return
	}

	success, err = presentations.VerifyVP(input.Presentation)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error verifying presentation: "+err.Error())
		return
	}

	if !success {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed to verify presentation")
		return
	}

	err = credentials.RevokeVC(&input.Presentation.VerifiableCredential[0])
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed to revoke credential: "+err.Error())
	}

	vcBytes := [32]byte{}
	copy(vcBytes[:], credentials.ConvertVCToBytes(input.Presentation.VerifiableCredential[0]))
	err = contract.RevokeVC(vcBytes)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed to revoke credential in registry: "+err.Error())
		return
	}

	ResponseSuccess(c, input.Presentation.VerifiableCredential[0])
}

func SendDocToRegistryHandler(c *gin.Context) {
	response := &ResponseData{}
	if err := c.ShouldBindJSON(response); err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	documentData := response.Data.([]interface{})[0].(map[string]interface{})
	document := models.CreateDIDDocument()
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{ErrorUnused: true, Result: document})
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error creating decoder: "+err.Error())
		return
	}
	err = decoder.Decode(documentData)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error reading did document: "+err.Error())
		return
	}

	_, valid := did.PrepareDID(document.ID)
	if !valid {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Improperly formatted did")
		return
	}

	docBytes := [32]byte{}
	copy(docBytes[:], did.ConvertDocToBytes(*document))
	err = contract.UploadDocument(docBytes)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error registering document: "+err.Error())
		return
	}

	ResponseSuccessWithMsg(c, "DID document has been successfully uploaded to registry")
}

func GetMinerListHandler(c *gin.Context) {
	didString := "did:metablox:" + c.Param("did")
	authenticationInfo := models.CreateAuthenticationInfo()
	err := c.BindJSON(authenticationInfo)

	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error reading input: "+err.Error())
		return
	}

	success, err := CheckNonce(c.ClientIP(), authenticationInfo.Nonce)
	if !success {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed to verify nonce: "+err.Error())
		return
	}

	opts := models.CreateResolutionOptions()
	resolutionMeta, doc, _ := did.Resolve(didString, opts)
	if resolutionMeta.Error != "" {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error resolving did document: "+resolutionMeta.Error)
		return
	}
	success, err = did.AuthenticateDocumentHolder(doc, authenticationInfo.Signature, authenticationInfo.Nonce)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error authenticating signature: "+err.Error())
		return
	}
	if !success {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Authentication failed")
		return
	}
	minerList, err := dao.GetMinerList()
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error retrieving miner info: "+err.Error())
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
