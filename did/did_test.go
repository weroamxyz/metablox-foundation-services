package did

import (
	"encoding/json"
	"testing"

	"github.com/metabloxDID/models"
	"github.com/stretchr/testify/assert"
)

const exampleDIDDocString = `{"@context":["https://w3id.org/did/v1","https://ns.did.ai/suites/secp256k1-2019/v1/"],"id":"did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo","created":"2022-03-31T12:53:19-07:00","updated":"2022-03-31T12:53:19-07:00","version":1,"verificationMethod":[{"id":"did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo#verification","type":"EcdsaSecp256k1VerificationKey2019","controller":"did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo","publicKeyMultibase":"zR4TQJaWaLA3vvYukULRJoxTsRmqCMsWuEJdDE8CJwRFCUijDGwCBP89xVcWdLRQaEM6b7wD294xCs8byy3CdDMYa"}],"authentication":"did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo#verification"}`

var invalidDIDMetadata = &models.ResolutionMetadata{Error: "invalid Did"}
var unknownMethodMetadata = &models.ResolutionMetadata{Error: "methodNotSupported"}
var emptyResolutionMetadata = &models.ResolutionMetadata{}
var emptyJSONRepresentationResolutionMetadata = &models.RepresentationResolutionMetadata{ContentType: "application/did+json"}
var emptyDocumentMetadata = &models.DocumentMetadata{}

func TestCreateDID(t *testing.T) {
	privKey := models.GenerateTestPrivKey()
	document, err := CreateDID(privKey)
	assert.Nil(t, err)

	exampleDocument := models.GenerateTestDIDDocument()
	assert.Equal(t, exampleDocument.Context, document.Context)
	assert.Equal(t, exampleDocument.ID, document.ID)
	//no point comparing create/update time, won't be equal
	assert.Equal(t, exampleDocument.Version, document.Version)
	assert.Equal(t, exampleDocument.VerificationMethod, document.VerificationMethod)
	assert.Equal(t, exampleDocument.Authentication, document.Authentication)
}

func TestConvertDocumentToJson(t *testing.T) {
	document := models.GenerateTestDIDDocument()
	jsonDoc, err := DocumentToJson(document)
	assert.Nil(t, err)
	assert.Equal(t, exampleDIDDocString, string(jsonDoc))
}

func TestResolveDID(t *testing.T) {
	options := &models.ResolutionOptions{}
	resolutionMeta, document, documentMeta := Resolve("bad:did", options) //missing final section
	assert.Equal(t, invalidDIDMetadata, resolutionMeta)
	assert.Nil(t, document)
	assert.Equal(t, emptyDocumentMetadata, documentMeta)

	resolutionMeta, document, documentMeta = Resolve("bad:did:string", options) //does not start with 'did'
	assert.Equal(t, invalidDIDMetadata, resolutionMeta)
	assert.Nil(t, document)
	assert.Equal(t, emptyDocumentMetadata, documentMeta)

	resolutionMeta, document, documentMeta = Resolve("did:ijdiej^&$:hbdsuhue", options) //includes invalid symbols in method
	assert.Equal(t, invalidDIDMetadata, resolutionMeta)
	assert.Nil(t, document)
	assert.Equal(t, emptyDocumentMetadata, documentMeta)

	resolutionMeta, document, documentMeta = Resolve("did:valid::!@#$%^&*()", options) //includes invalid symbols in identifier
	assert.Equal(t, invalidDIDMetadata, resolutionMeta)
	assert.Nil(t, document)
	assert.Equal(t, emptyDocumentMetadata, documentMeta)

	resolutionMeta, document, documentMeta = Resolve("did:valid:iuhienwd:", options) //identifier ends with ':'
	assert.Equal(t, invalidDIDMetadata, resolutionMeta)
	assert.Nil(t, document)
	assert.Equal(t, emptyDocumentMetadata, documentMeta)

	resolutionMeta, document, documentMeta = Resolve("did:valid:iuhienwd", options) //'valid' is an unknown method name
	assert.Equal(t, unknownMethodMetadata, resolutionMeta)
	assert.Nil(t, document)
	assert.Equal(t, emptyDocumentMetadata, documentMeta)

	resolutionMeta, document, documentMeta = Resolve("did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo", options) //resolvable did
	assert.Equal(t, emptyResolutionMetadata, resolutionMeta)
	exampleDocument := models.GenerateTestDIDDocument()
	assert.Equal(t, exampleDocument.Context, document.Context)
	assert.Equal(t, exampleDocument.ID, document.ID)
	//no point comparing create/update time, won't be equal
	assert.Equal(t, exampleDocument.Version, document.Version)
	assert.Equal(t, exampleDocument.VerificationMethod, document.VerificationMethod)
	assert.Equal(t, exampleDocument.Authentication, document.Authentication)
	assert.Nil(t, documentMeta)
}

func TestResolveDIDRepresentation(t *testing.T) {
	options := &models.RepresentationResolutionOptions{Accept: "application/did+json"}

	resolutionMeta, byteStream, documentMeta := ResolveRepresentation("did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo", options) //resolvable did
	assert.Equal(t, emptyJSONRepresentationResolutionMetadata, resolutionMeta)
	exampleDocument := models.GenerateTestDIDDocument()
	document := models.CreateDIDDocument()
	err := json.Unmarshal(byteStream, document)
	assert.Nil(t, err)

	assert.Equal(t, exampleDocument.Context, document.Context)
	assert.Equal(t, exampleDocument.ID, document.ID)
	//no point comparing create/update time, won't be equal
	assert.Equal(t, exampleDocument.Version, document.Version)
	assert.Equal(t, exampleDocument.VerificationMethod, document.VerificationMethod)
	assert.Equal(t, exampleDocument.Authentication, document.Authentication)
	assert.Nil(t, documentMeta)
}
