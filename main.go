package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"

	"github.com/metabloxDID/did"
	"github.com/metabloxDID/models"
)

func main() {
	document, privKey, err := did.CreateDID()
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonDoc, err := did.DocumentToJson(document)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("created DID document: ", string(jsonDoc))

	options := &models.ResolutionOptions{}
	did.Resolve("bad:did", options)
	did.Resolve("bad:did:string", options)
	did.Resolve("did:ijdiej^&$:hbdsuhue", options)
	did.Resolve("did:valid::!@#$%^&*()", options)
	did.Resolve("did:valid:iuhienwd:", options)
	did.Resolve("did:metablox:jhbwehj", options)

	sampleMessage := "This message will be encrypted with a private key"
	x, y, err := ecdsa.Sign(rand.Reader, privKey, []byte(sampleMessage))
	if err != nil {
		fmt.Println("Failed to create signature: ", err)
		return
	}
	verificationResult, err := did.AuthenticateDocumentSubject(document, sampleMessage, x, y)
	if err != nil {
		fmt.Println(err)
	}

	if verificationResult {
		fmt.Println("Successfully verified document subject!")
	} else {
		fmt.Println("Failed to verify document subject")
	}
}
