package controllers

import "github.com/MetaBloxIO/metablox-foundation-services/models"

//When parsing a credential in JSON, credential subject is read in as interface.
//Have to manually convert the interface to the appropriate subject type
func ConvertCredentialSubject(vc *models.VerifiableCredential) {
	for i := 0; i < len(vc.Type); i++ {
		subjectMap := vc.CredentialSubject.(map[string]interface{})
		switch vc.Type[i] {
		case models.TypeWifi:
			wifiInfo := models.CreateWifiAccessInfo()
			wifiInfo.ID = subjectMap["id"].(string)
			wifiInfo.Type = subjectMap["type"].(string)
			vc.CredentialSubject = *wifiInfo
			return
		case models.TypeMining:
			miningInfo := models.CreateMiningLicenseInfo()
			miningInfo.ID = subjectMap["id"].(string)
			miningInfo.Model = subjectMap["model"].(string)
			miningInfo.Name = subjectMap["name"].(string)
			miningInfo.Serial = subjectMap["serial"].(string)
			vc.CredentialSubject = *miningInfo
			return
		case models.TypeStaking:
			stakingInfo := models.CreateStakingVCInfo()
			stakingInfo.ID = subjectMap["id"].(string)
			vc.CredentialSubject = *stakingInfo
			return
		}
	}
}
