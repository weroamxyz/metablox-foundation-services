package controllers

import (
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
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

func Init() {
	NonceLookup = make(map[string]string)
	ValidIssuers = []string{"did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo"}
}

//Compare the nonce a user has given with the one they are assigned. If user does not have an assigned nonce, give them one.
func CheckNonce(ip, givenNonce string) (bool, string) {
	assignedNonce, found := NonceLookup[ip]
	if !found {
		NonceLookup[ip] = time.Now().String()
		return false, NonceLookup[ip]
	}

	if givenNonce == assignedNonce {
		return true, ""
	}

	return false, assignedNonce
}

//If user provided correct nonce, then assign them a new one and return that value
func UpdateNonce(ip string) string {
	NonceLookup[ip] = time.Now().String()
	return NonceLookup[ip]
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
	response := &ResponseData{}
	if err := c.ShouldBindJSON(response); err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}
	authenticationData := response.Data.([]interface{})[0].(map[string]interface{})
	authenticationInfo := models.CreateAuthenticationInfo()

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{ErrorUnused: true, Result: authenticationInfo})
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error creating decoder: "+err.Error())
		return
	}

	err = decoder.Decode(authenticationData)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error reading authentication info: "+err.Error())
		return
	}

	success, returnNonce := CheckNonce(c.ClientIP(), authenticationInfo.Nonce)
	if !success {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Incorrect nonce value, expected '"+returnNonce+"'")
		return
	}

	opts := models.CreateResolutionOptions()
	resolutionMeta, issuerDocument, _ := did.Resolve(didString, opts) //TODO: still needs proper implementation once smart contract is ready
	if resolutionMeta.Error != "" {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed to resolve did '"+didString+"': "+resolutionMeta.Error)
		return
	}

	success, err = did.AuthenticateDocumentHolder(issuerDocument, authenticationInfo.Signature, authenticationInfo.Nonce)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error authenticating signature: "+err.Error())
		return
	}
	if !success {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Authentication failed")
		return
	}

	subjectData := response.Data.([]interface{})[1].(map[string]interface{})
	wifiAccessInfo := models.CreateWifiAccessInfo()

	decoder, err = mapstructure.NewDecoder(&mapstructure.DecoderConfig{ErrorUnused: true, Result: wifiAccessInfo})
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error creating decoder: "+err.Error())
		return
	}

	err = decoder.Decode(subjectData)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error reading subject info: "+err.Error())
		return
	}

	issuerPrivateKey, _ := crypto.GenerateKey() //TODO: modify to get actual issuer private key as opposed to arbitrary value

	newVC, err := credentials.CreateWifiAccessVC(issuerDocument, wifiAccessInfo, issuerPrivateKey)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error creating wifi access VC: "+err.Error())
		return
	}

	//TODO: uncomment once smart contract is completed + deployed
	/*err = contract.CreateVC(newVC)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error registering wifi access VC: "+err.Error())
		return
	}*/

	newNonce := UpdateNonce(c.ClientIP())

	ResponseSuccessWithMsgAndData(c, "new nonce is: '"+newNonce+"'", newVC)
}

func IssueMiningVCHandler(c *gin.Context) {
	didString := "did:metablox:" + c.Param("did")
	if !CheckIfValidIssuer(didString) {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Provided did is not a valid issuer")
		return
	}
	response := &ResponseData{}
	if err := c.ShouldBindJSON(response); err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}
	authenticationData := response.Data.([]interface{})[0].(map[string]interface{})
	authenticationInfo := models.CreateAuthenticationInfo()

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{ErrorUnused: true, Result: authenticationInfo})
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error creating decoder: "+err.Error())
		return
	}

	err = decoder.Decode(authenticationData)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error reading authentication info: "+err.Error())
		return
	}

	success, returnNonce := CheckNonce(c.ClientIP(), authenticationInfo.Nonce)
	if !success {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Incorrect nonce value, expected '"+returnNonce+"'")
		return
	}

	opts := models.CreateResolutionOptions()
	resolutionMeta, issuerDocument, _ := did.Resolve(didString, opts) //TODO: still needs proper implementation once smart contract is ready
	if resolutionMeta.Error != "" {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed to resolve did '"+didString+"': "+resolutionMeta.Error)
		return
	}

	success, err = did.AuthenticateDocumentHolder(issuerDocument, authenticationInfo.Signature, authenticationInfo.Nonce)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error authenticating signature: "+err.Error())
		return
	}
	if !success {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Authentication failed")
		return
	}

	subjectData := response.Data.([]interface{})[1].(map[string]interface{})
	miningLicenseInfo := models.CreateMiningLicenseInfo()

	decoder, err = mapstructure.NewDecoder(&mapstructure.DecoderConfig{ErrorUnused: true, Result: miningLicenseInfo})
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error creating decoder: "+err.Error())
		return
	}

	err = decoder.Decode(subjectData)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error reading subject info: "+err.Error())
		return
	}

	issuerPrivateKey, _ := crypto.GenerateKey() //TODO: modify to get actual issuer private key as opposed to arbitrary value

	newVC, err := credentials.CreateMiningLicenseVC(issuerDocument, miningLicenseInfo, issuerPrivateKey)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error creating mining license VC: "+err.Error())
		return
	}

	//TODO: uncomment once smart contract is completed + deployed
	/*err = contract.CreateVC(newVC)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error registering mining license VC: "+err.Error())
		return
	}*/

	newNonce := UpdateNonce(c.ClientIP())

	ResponseSuccessWithMsgAndData(c, "new nonce is: '"+newNonce+"'", newVC)
}

func RenewVCHandler(c *gin.Context) {
	didString := "did:metablox:" + c.Param("did")
	response := &ResponseData{}
	if err := c.ShouldBindJSON(response); err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}
	authenticationData := response.Data.([]interface{})[0].(map[string]interface{})
	authenticationInfo := models.CreateAuthenticationInfo()

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{ErrorUnused: true, Result: authenticationInfo})
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error creating decoder: "+err.Error())
		return
	}

	err = decoder.Decode(authenticationData)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error reading authentication info: "+err.Error())
		return
	}

	success, returnNonce := CheckNonce(c.ClientIP(), authenticationInfo.Nonce)
	if !success {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Incorrect nonce value, expected '"+returnNonce+"'")
		return
	}

	vpData := response.Data.([]interface{})[1].(map[string]interface{})
	vp := models.CreatePresentation()

	decoder, err = mapstructure.NewDecoder(&mapstructure.DecoderConfig{ErrorUnused: true, Result: vp})
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error creating decoder: "+err.Error())
		return
	}

	err = decoder.Decode(vpData)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error reading presentation: "+err.Error())
		return
	}

	success, returnNonce = CheckNonce(c.ClientIP(), vp.Proof.Nonce)
	if !success {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Incorrect nonce value, expected '"+returnNonce+"'")
		return
	}

	if vp.VerifiableCredential[0].Issuer != didString {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Provided did does not match issuer of credential")
		return
	}

	opts := models.CreateResolutionOptions()
	resolutionMeta, issuerDocument, _ := did.Resolve(didString, opts) //TODO: still needs proper implementation once smart contract is ready
	if resolutionMeta.Error != "" {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed to resolve did '"+didString+"': "+resolutionMeta.Error)
		return
	}

	success, err = did.AuthenticateDocumentHolder(issuerDocument, authenticationInfo.Signature, authenticationInfo.Nonce)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error authenticating signature: "+err.Error())
		return
	}
	if !success {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Authentication failed")
		return
	}

	success, err = presentations.VerifyVP(vp)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error verifying presentation: "+err.Error())
		return
	}

	if !success {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed to verify presentation")
		return
	}

	err = credentials.RenewVC(&vp.VerifiableCredential[0])
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed to renew credential: "+err.Error())
	}

	//TODO: uncomment once smart contract is completed + deployed
	/*err = contract.RenewVC(&wifiVP.VerifiableCredential[0])
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed to renew credential in registry: "+err.Error())
		return
	}*/

	newNonce := UpdateNonce(c.ClientIP())
	ResponseSuccessWithMsgAndData(c, "new nonce is: '"+newNonce+"'", vp.VerifiableCredential[0])
}

func RevokeVCHandler(c *gin.Context) {
	didString := "did:metablox:" + c.Param("did")
	response := &ResponseData{}
	if err := c.ShouldBindJSON(response); err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}
	authenticationData := response.Data.([]interface{})[0].(map[string]interface{})
	authenticationInfo := models.CreateAuthenticationInfo()

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{ErrorUnused: true, Result: authenticationInfo})
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error creating decoder: "+err.Error())
		return
	}

	err = decoder.Decode(authenticationData)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error reading authentication info: "+err.Error())
		return
	}

	success, returnNonce := CheckNonce(c.ClientIP(), authenticationInfo.Nonce)
	if !success {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Incorrect nonce value, expected '"+returnNonce+"'")
		return
	}

	vpData := response.Data.([]interface{})[1].(map[string]interface{})
	vp := models.CreatePresentation()

	decoder, err = mapstructure.NewDecoder(&mapstructure.DecoderConfig{ErrorUnused: true, Result: vp})
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error creating decoder: "+err.Error())
		return
	}

	err = decoder.Decode(vpData)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error reading presentation: "+err.Error())
		return
	}

	success, returnNonce = CheckNonce(c.ClientIP(), vp.Proof.Nonce)
	if !success {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Incorrect nonce value, expected '"+returnNonce+"'")
		return
	}

	if vp.VerifiableCredential[0].Issuer != didString {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Provided did does not match issuer of credential")
		return
	}

	opts := models.CreateResolutionOptions()
	resolutionMeta, issuerDocument, _ := did.Resolve(didString, opts) //TODO: still needs proper implementation once smart contract is ready
	if resolutionMeta.Error != "" {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed to resolve did '"+didString+"': "+resolutionMeta.Error)
		return
	}

	success, err = did.AuthenticateDocumentHolder(issuerDocument, authenticationInfo.Signature, authenticationInfo.Nonce)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error authenticating signature: "+err.Error())
		return
	}
	if !success {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Authentication failed")
		return
	}

	success, err = presentations.VerifyVP(vp)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error verifying presentation: "+err.Error())
		return
	}

	if !success {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed to verify presentation")
		return
	}

	err = credentials.RevokeVC(&vp.VerifiableCredential[0])
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed to revoke credential: "+err.Error())
	}

	//TODO: uncomment once smart contract is completed + deployed
	/*err = contract.RevokeVC(&wifiVP.VerifiableCredential[0])
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed to revoke credential in registry: "+err.Error())
		return
	}*/

	newNonce := UpdateNonce(c.ClientIP())
	ResponseSuccessWithMsgAndData(c, "new nonce is: '"+newNonce+"'", vp.VerifiableCredential[0])
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

	//TODO: Upload did document to registry now that it has been reviewed
	ResponseSuccessWithMsg(c, "DID document has been successfully uploaded to registry")
}

func GetMinerListHandler(c *gin.Context) {
	didString := "did:metablox:" + c.Param("did")
	response := &ResponseData{}
	if err := c.ShouldBindJSON(response); err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	authenticationData := response.Data.([]interface{})[0].(map[string]interface{})
	authenticationInfo := models.CreateAuthenticationInfo()
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{ErrorUnused: true, Result: &authenticationInfo})
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error creating decoder: "+err.Error())
		return
	}
	err = decoder.Decode(authenticationData)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error reading input: "+err.Error())
		return
	}

	success, returnNonce := CheckNonce(c.ClientIP(), authenticationInfo.Nonce)
	if !success {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Incorrect nonce value, expected '"+returnNonce+"'")
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

	newNonce := UpdateNonce(c.ClientIP())

	ResponseSuccessWithMsgAndData(c, "new nonce is: '"+newNonce+"'", minerList)
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
