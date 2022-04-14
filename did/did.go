package did

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/metabloxDID/models"
	"github.com/mr-tron/base58"
	"github.com/multiformats/go-multibase"
)

func GenerateDIDString(privKey *ecdsa.PrivateKey) string {
	privData := crypto.FromECDSA(privKey)

	hash := sha256.New()
	hash.Write(privData)
	hashData := hash.Sum(nil)
	didString := base58.Encode(hashData)
	returnString := "did:metablox:" + didString
	return returnString
}

func CreateDID(privKey *ecdsa.PrivateKey) (*models.DIDDocument, error) {

	document := new(models.DIDDocument)

	var err error
	document.ID = GenerateDIDString(privKey)
	if err != nil {
		return nil, err
	}
	document.Context = make([]string, 0)
	document.Context = append(document.Context, "https://w3id.org/did/v1")
	document.Context = append(document.Context, "https://ns.did.ai/suites/secp256k1-2019/v1/")
	document.Created = time.Now().Format(time.RFC3339)
	document.Updated = document.Created
	document.Version = 1

	pubData := crypto.FromECDSAPub(&privKey.PublicKey)

	VM := models.VerificationMethod{}
	VM.ID = document.ID + "#verification"
	VM.MultibaseKey, err = multibase.Encode(multibase.Base58BTC, pubData)
	if err != nil {
		return nil, err
	}
	VM.Controller = document.ID
	VM.MethodType = "EcdsaSecp256k1VerificationKey2019"

	document.VerificationMethod = append(document.VerificationMethod, VM)
	document.Authentication = VM.ID

	//once blockchain is implemented, will also need to upload the document to the blockchain

	return document, nil
}

func DocumentToJson(document *models.DIDDocument) ([]byte, error) {
	jsonDoc, err := json.Marshal(document)
	if err != nil {
		return nil, err
	}
	return jsonDoc, nil
}

func JsonToDocument(jsonDoc []byte) (*models.DIDDocument, error) {
	document := models.CreateDIDDocument()
	err := json.Unmarshal(jsonDoc, document)
	if err != nil {
		return nil, err
	}
	return document, nil
}

//check format of DID string
func PrepareDID(did string) ([]string, bool) {
	splitString := strings.Split(did, ":")
	if len(splitString) < 3 {
		fmt.Println("Not enough sections in DID")
		return nil, false
	}
	prefix := splitString[0]
	if prefix != "did" {
		fmt.Println("First section of DID was '" + prefix + "' instead of 'did'")
		return nil, false
	}

	methodName := splitString[1]
	methodNamePattern := `^[A-Za-z0-9]+$`
	match, _ := regexp.MatchString(methodNamePattern, methodName)
	if !match {
		fmt.Println("Method name '" + methodName + "' is formatted incorrectly")
		return nil, false
	}

	identifierPattern := `^([a-zA-Z0-9\._\-]*(%[0-9A-Fa-f][0-9A-Fa-f])*)*$`
	identifierExp, _ := regexp.Compile(identifierPattern)

	for i := 2; i < len(splitString); i++ {
		identifierSection := splitString[i]
		match = identifierExp.MatchString(identifierSection)
		if !match {
			fmt.Println("Identifier section '" + identifierSection + "' is formatted incorrectly")
			return nil, false
		}

		if i == len(splitString)-1 && len(identifierSection) == 0 {
			fmt.Println("Final portion of identifier is empty")
			return nil, false
		}
	}
	return splitString, true
}

func Resolve(did string, options *models.ResolutionOptions) (*models.ResolutionMetadata, *models.DIDDocument, *models.DocumentMetadata) {
	splitDID, valid := PrepareDID(did)
	if !valid {
		return &models.ResolutionMetadata{Error: "invalid Did"}, nil, &models.DocumentMetadata{}
	}

	methodName := splitDID[1]
	switch methodName {
	case "metablox":
		//remove placeholder values after blockchain is implemented
		placeholderDoc := models.GenerateTestDIDDocument()
		placeholderHash := sha256.Sum256(ConvertDocToBytes(*placeholderDoc))
		return MetabloxRead(splitDID[2], options, placeholderDoc, placeholderHash)
	default:
		fmt.Println("Unable to resolve unknown method type '" + methodName + "'")
		return &models.ResolutionMetadata{Error: "methodNotSupported"}, nil, &models.DocumentMetadata{}
	}
}

func ResolveRepresentation(did string, options *models.RepresentationResolutionOptions) (*models.RepresentationResolutionMetadata, []byte, *models.DocumentMetadata) {
	splitDID, valid := PrepareDID(did)
	if !valid {
		return &models.RepresentationResolutionMetadata{Error: "invalid Did"}, nil, &models.DocumentMetadata{}
	}

	methodName := splitDID[1]
	switch methodName {
	case "metablox":
		placeholderDoc := models.GenerateTestDIDDocument()
		placeholderHash := sha256.Sum256(ConvertDocToBytes(*placeholderDoc))
		return MetabloxReadRepresentation(splitDID[2], options, placeholderDoc, placeholderHash)
	default:
		fmt.Println("Unable to resolve unknown method type '" + methodName + "'")
		return &models.RepresentationResolutionMetadata{Error: "methodNotSupported"}, nil, &models.DocumentMetadata{}
	}
}

func AuthenticateDocumentSubject(document *models.DIDDocument, message, signature []byte) (bool, error) {
	//The subject of the document is the person who has the private key matching the public key in the Authentication verification method

	//Get Authentication VM
	var authenticationMethod *models.VerificationMethod
	for _, vm := range document.VerificationMethod {
		if vm.ID == document.Authentication {
			authenticationMethod = &vm
			break
		}
	}
	if authenticationMethod == nil {
		return false, errors.New("Failed to find authentication method '" + document.Authentication + "'")
	}

	switch authenticationMethod.MethodType {
	case "EcdsaSecp256k1VerificationKey2019":
		_, pubData, err := multibase.Decode(authenticationMethod.MultibaseKey)
		if err != nil {
			return false, err
		}

		result := crypto.VerifySignature(pubData, message, signature)
		return result, nil

	default:
		return false, errors.New("Unable to resolve unknown verification method type '" + authenticationMethod.MethodType + "'")
	}
}

func MetabloxRead(identifier string, options *models.ResolutionOptions, placeholderDocument *models.DIDDocument, placeholderHash [32]byte) (*models.ResolutionMetadata, *models.DIDDocument, *models.DocumentMetadata) {
	//Once blockchain is implemented, this function will use the identifier to retrieve the matching document from the blockchain (if it exists).
	//Some mechanism should likely be in place to ensure the document was not modified during the transfer, ex. an encrypted hash.

	//Temporarily treating the document provided in placeholderDocument as the document we've retrieved from the blockchain, and placeholderHash as the hash generated by the blockchain.
	//These inputs should be removed after the blockchain is implemented

	generatedDocument := placeholderDocument
	generatedHash := placeholderHash

	docID, success := PrepareDID(generatedDocument.ID)
	if !success {
		return &models.ResolutionMetadata{Error: "document DID is invalid"}, nil, nil
	}

	if docID[2] != identifier {
		return &models.ResolutionMetadata{Error: "generated document DID does not match provided DID"}, nil, nil
	}

	comparisonHash := sha256.Sum256(ConvertDocToBytes(*generatedDocument))
	if comparisonHash != generatedHash {
		return &models.ResolutionMetadata{Error: "document failed hash check"}, nil, nil
	}
	return &models.ResolutionMetadata{}, generatedDocument, nil
}

func MetabloxReadRepresentation(identifier string, options *models.RepresentationResolutionOptions, placeholderDocument *models.DIDDocument, placeholderHash [32]byte) (*models.RepresentationResolutionMetadata, []byte, *models.DocumentMetadata) {
	//Should be similar to MetabloxRead, but returns the document in a specific representation format.
	//Representation type is included in options and returned in resolution metadata
	readOptions := models.CreateResolutionOptions()
	readResolutionMeta, document, readDocumentMeta := MetabloxRead(identifier, readOptions, placeholderDocument, placeholderHash)
	if readResolutionMeta.Error != "" {
		return &models.RepresentationResolutionMetadata{Error: readResolutionMeta.Error}, nil, nil
	}

	switch options.Accept {
	case "application/did+json":
		fallthrough
	default: //default to JSON format if options.Accept is empty/invalid
		byteStream, err := json.Marshal(document)
		if err != nil {
			return &models.RepresentationResolutionMetadata{Error: "failed to convert document into JSON"}, nil, nil
		}
		return &models.RepresentationResolutionMetadata{ContentType: "application/did+json"}, byteStream, readDocumentMeta
	}
}

func ConvertDocToBytes(doc models.DIDDocument) []byte {
	var convertedBytes []byte
	for _, item := range doc.Context {
		convertedBytes = bytes.Join([][]byte{convertedBytes, []byte(item)}, []byte{})
	}

	convertedBytes = bytes.Join([][]byte{convertedBytes, []byte(doc.ID), []byte(doc.Created), []byte(doc.Updated), []byte(strconv.Itoa(doc.Version))}, []byte{})
	for _, item := range doc.VerificationMethod {
		convertedBytes = bytes.Join([][]byte{convertedBytes, ConvertVMToBytes(item)}, []byte{})
	}

	convertedBytes = bytes.Join([][]byte{convertedBytes, []byte(doc.Authentication)}, []byte{})

	for _, item := range doc.Service {
		convertedBytes = bytes.Join([][]byte{convertedBytes, ConvertServiceToBytes(item)}, []byte{})
	}
	return convertedBytes
}

func ConvertVMToBytes(vm models.VerificationMethod) []byte {
	var convertedBytes []byte

	convertedBytes = bytes.Join([][]byte{convertedBytes, []byte(vm.ID), []byte(vm.MethodType), []byte(vm.Controller), []byte(vm.MultibaseKey)}, []byte{})
	return convertedBytes
}

func ConvertServiceToBytes(service models.Service) []byte {
	var convertedBytes []byte

	convertedBytes = bytes.Join([][]byte{convertedBytes, []byte(service.ID), []byte(service.Type), []byte(service.ServiceEndpoint)}, []byte{})
	return convertedBytes
}
