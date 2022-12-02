package service

import (
	"encoding/json"
	"github.com/MetaBloxIO/metablox-foundation-services/dao"
	"github.com/MetaBloxIO/metablox-foundation-services/errval"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/MetaBloxIO/metablox-foundation-services/presentations"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/shopspring/decimal"
	logger "github.com/sirupsen/logrus"
	"time"
)

func WorkloadValidate(req *models.WorkloadDTO) error {
	// verify vp/vc
	//opts := models.CreateResolutionOptions()
	//resolutionMeta, issuerDocument, _ := did.Resolve(credentials.IssuerDID, opts)
	//if resolutionMeta.Error != "" {
	//	return errors.New(resolutionMeta.Error)
	//}

	success0, err0 := presentations.VerifyVP(req.Identity.Miner)
	if err0 != nil {
		logger.Warn(err0)
	}

	if !success0 {
		return errval.ErrVerifyPresent
		logger.Warn("Miner Verify failed", errval.ErrVerifyPresent)
	}

	success1, err1 := presentations.VerifyVP(req.Identity.Validator)
	if err1 != nil {
		logger.Warn(err1)
	}

	if !success1 {
		return errval.ErrVerifyPresent
		logger.Warn("Validator Verify failed", errval.ErrVerifyPresent)
	}

	info := models.MiningLicenseInfo{}

	credentials := req.Identity.Miner.VerifiableCredential
	if len(credentials) > 0 {
		for _, vc := range credentials {
			if slice.Contain[string](vc.Type, "VerifiableCredential") && slice.Contain(vc.Type, "MiningLicense") {
				m, ok := vc.CredentialSubject.(map[string]string)
				if ok {
					info.CredentialID = m["credentialID"]
					info.Model = m["model"]
					info.Serial = m["serial"]
					info.Name = m["name"]
				}
			}
		}
	}

	qosJson, _ := json.Marshal(req.Qos)
	tracksJson, _ := json.Marshal(req.Tracks)
	workload := &models.Workload{
		Miner:        req.Identity.Miner.Holder,
		Validator:    req.Identity.Validator.Holder,
		Qos:          string(qosJson),
		Tracks:       string(tracksJson),
		CredentialID: info.CredentialID,
		Model:        info.Model,
		Serial:       info.Serial,
		Name:         info.Name,
		CreateTime:   time.Now(),
	}

	if _, err := dao.InsertWorkload(workload); err != nil {
		return err
	}

	return nil
}

func GetWorkloadProfit(did string) (decimal.Decimal, error) {
	total, err := dao.GetProfitByDID(did)
	if err != nil {
		return decimal.Decimal{}, err
	}
	return total, nil
}
