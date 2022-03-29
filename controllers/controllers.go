package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/metabloxDID/credentials"
	"github.com/metabloxDID/did"
	"github.com/metabloxDID/models"
	"github.com/mitchellh/mapstructure"
)

const CtxUserToken = "userToken"

func IssueVCHandler(c *gin.Context) {
	response := &ResponseData{}
	if err := c.ShouldBindJSON(response); err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}
	subjectData := response.Data.([]interface{})[0].(map[string]interface{})
	subjectInfo := models.CreateSubjectInfo()

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{ErrorUnused: true, Result: subjectInfo})
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error creating decoder: "+err.Error())
		return
	}

	err = decoder.Decode(subjectData)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error reading subject info: "+err.Error())
		return
	}

	opts := models.CreateResolutionOptions()
	issuerDID := "did:metablox:sampleIssuerDID"                       //TODO: modify to get actual issuer DID as opposed to hard-coded placeholder value
	resolutionMeta, issuerDocument, _ := did.Resolve(issuerDID, opts) //TODO: still needs proper implementation once smart contract is ready
	if resolutionMeta.Error != "" {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Failed to resolve did '"+subjectInfo.ID+"'")
		return
	}
	issuerPrivateData := []byte{66, 94, 211, 215, 5, 230, 103, 245, 103, 92, 207, 182, 241, 116, 121, 103, 52, 172, 68, 78, 93, 241, 37, 34, 220, 30, 122, 173, 224, 212, 11, 124} //TODO: modify to get actual issuer private key as opposed to arbitrary hard-coded value

	newVC, err := credentials.CreateVC(issuerDocument, subjectInfo, issuerPrivateData)
	if err != nil {
		ResponseErrorWithMsg(c, http.StatusNotAcceptable, "Error creating VC: "+err.Error())
		return
	}

	ResponseSuccess(c, newVC)
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
