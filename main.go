package main

import (
	"crypto/sha256"

	"github.com/dappley/go-dappley/crypto/keystore/secp256k1"
	"github.com/metabloxDID/credentials"
	"github.com/metabloxDID/did"
	"github.com/metabloxDID/key"
	"github.com/metabloxDID/models"
	"github.com/metabloxDID/routers"
	"github.com/metabloxDID/settings"
	logger "github.com/sirupsen/logrus"
)

func main() {
	err := settings.Init()
	if err != nil {
		logger.Error(err)
		return
	}

	_, fileName, err := key.GenerateNewPrivateKey()
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("dumped data into file: ", fileName)

	loadedPrivKey, err := key.LoadPrivateKey("privateKey1")
	if err != nil {
		logger.Error(err)
		return
	}

	document, err := did.CreateDID(loadedPrivKey)
	if err != nil {
		logger.Info(err)
		return
	}
	jsonDoc, err := did.DocumentToJson(document)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("Created DID document: ", string(jsonDoc))

	options := &models.ResolutionOptions{}
	did.Resolve("bad:did", options)
	did.Resolve("bad:did:string", options)
	did.Resolve("did:ijdiej^&$:hbdsuhue", options)
	did.Resolve("did:valid::!@#$%^&*()", options)
	did.Resolve("did:valid:iuhienwd:", options)
	did.Resolve("did:metablox:jhbwehj", options)

	privData, err := secp256k1.FromECDSAPrivateKey(loadedPrivKey)
	if err != nil {
		logger.Error(err)
		return
	}

	sampleMessage := "This message will be encrypted with a private key"
	hashedMessage := sha256.Sum256([]byte(sampleMessage))
	signature, err := secp256k1.Sign(hashedMessage[:], privData)
	if err != nil {
		logger.Error("Failed to create signature: ", err)
		return
	}
	verificationResult, err := did.AuthenticateDocumentSubject(document, hashedMessage[:], signature)
	if err != nil {
		logger.Error(err)
		return
	}

	if verificationResult {
		logger.Info("Successfully verified document subject!")
	} else {
		logger.Info("Failed to verify document subject")
	}

	sampleSubject := models.CreateSubjectInfo()
	sampleSubject.ID = document.ID
	sampleSubject.Type = make([]string, 0)
	sampleSubject.Type = append(sampleSubject.Type, "sampleType")
	sampleSubject.GivenName = "John"
	sampleSubject.FamilyName = "Jacobs"
	sampleSubject.Gender = "Male"
	sampleSubject.BirthCountry = "Canada"
	sampleSubject.BirthDate = "2022-03-22"

	document.ID = "did:metablox:sampleIssuer"

	sampleVC, err := credentials.CreateVC(document, sampleSubject, loadedPrivKey)
	if err != nil {
		logger.Error(err)
		return
	}
	jsonVC, err := credentials.VCToJson(sampleVC)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("Created VC: ", string(jsonVC))

	verificationResult, err = credentials.VerifyVC(sampleVC) //since did.Resolve isn't implemented fully yet, this is expected to fail
	if err != nil {
		logger.Error(err)
	}

	if verificationResult {
		logger.Info("Successfully verified credential!")
	} else {
		logger.Info("Failed to verify credential")
	}

	verificationResult, err = credentials.VerifyVCSecp256k1(sampleVC, document.VerificationMethod[0]) //can use this function to just test the verification without needed to use did.Resolve
	if err != nil {
		logger.Error(err)
	}

	if verificationResult {
		logger.Info("Successfully verified credential!")
	} else {
		logger.Info("Failed to verify credential")
	}
	routers.Setup()
}
