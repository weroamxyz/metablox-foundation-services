package service

import (
	"database/sql"
	"github.com/MetaBloxIO/did-sdk-go"
	"github.com/MetaBloxIO/metablox-foundation-services/contract"
	"github.com/MetaBloxIO/metablox-foundation-services/dao"
	"github.com/MetaBloxIO/metablox-foundation-services/errval"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

func GetWifiUserInfo(vp *did.VerifiablePresentation) (*models.WifiUserInfo, error) {
	var (
		userInfo *models.WifiUserInfo
	)

	success, err := did.VerifyVP(vp, contract.GetRegistry())
	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	if !success {
		logger.Warn("Miner Verify failed", errval.ErrVerifyPresent)
		return nil, errval.ErrVerifyPresent
	}

	didSplit, _ := did.PrepareDID(vp.Holder)

	// check if exist
	userInfo, err = dao.GetWifiUserInfo(didSplit[2])
	if err == nil {
		userInfo.Username = userInfo.Username + "@metablox.io"
		return userInfo, nil
	}

	if err != sql.ErrNoRows {
		return nil, err
	}

	userInfo = &models.WifiUserInfo{
		Username: didSplit[2],
		Password: uuid.New().String(),
	}

	if _, err = dao.InsertWifiUserInfo(userInfo); err != nil {
		return nil, err
	}

	userInfo.Username = userInfo.Username + "@metablox.io"
	return userInfo, nil
}
