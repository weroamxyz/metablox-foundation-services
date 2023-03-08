package controller

import (
	"github.com/MetaBloxIO/did-sdk-go"
)

// When parsing a credential in JSON, credential subject is read in as interface.
// Have to manually convert the interface to the appropriate subject type
func ConvertCredentialSubject(vc *did.VerifiableCredential) {
	for i := 0; i < len(vc.Type); i++ {
		subjectMap := vc.CredentialSubject.(map[string]interface{})
		switch vc.Type[i] {
		case did.TypeWifi:
			wifiInfo := did.CreateWifiAccessInfo()
			wifiInfo.ID = subjectMap["id"].(string)
			wifiInfo.Type = subjectMap["type"].(string)
			vc.CredentialSubject = *wifiInfo
			return
		case did.TypeMining:
			miningInfo := did.CreateMiningLicenseInfo()
			miningInfo.ID = subjectMap["id"].(string)
			miningInfo.Model = subjectMap["model"].(string)
			miningInfo.Name = subjectMap["name"].(string)
			miningInfo.Serial = subjectMap["serial"].(string)
			vc.CredentialSubject = *miningInfo
			return
		}
	}
}
