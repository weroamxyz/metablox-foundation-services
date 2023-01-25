package service

import (
	"github.com/MetaBloxIO/metablox-foundation-services/dao"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/pkg/errors"
	logger "github.com/sirupsen/logrus"
	"strings"
	"time"
)

const FreeWifiType = "free"
const OpenroamingWifiType = "openroaming"

func HandleHeartBeat(beat *models.HeartbeatInfo) error {
	// Verify parameters
	if beat == nil || len(beat.RadioStatus) == 0 {
		return errors.New("empty beats")
	}

	if beat.DID == "" {
		return errors.New("empty did")
	}

	bssids := make([]string, 0)
	ssid := ""

	for _, v := range beat.RadioStatus {

		if v.Disabled {
			logger.Debug("beat disabled,ignored")
			continue
		}
		if v.Type == FreeWifiType {
			continue
		}

		if v.Type != OpenroamingWifiType {
			logger.Error("beat type invalid")
			return errors.New("beat type invalid")
		}

		if v.Bssid == "" {
			logger.Error("empty bssid")
			return errors.New("empty bssid")
		}

		if v.Ssid == "" {
			logger.Error("empty ssid")
			return errors.New("empty ssid")
		}

		if ssid == "" {
			ssid = v.Ssid
		}

		bssids = append(bssids, v.Bssid)

	}

	now := time.Now()
	// merge data
	minerInfo := &models.MinerInfo{
		SSID:           ssid,
		BSSID:          strings.Join(bssids, ","),
		OnlineStatus:   true,
		MiningPower:    1,
		IsMinable:      true,
		DID:            beat.DID,
		CreateTime:     &now,
		SignalStrength: "Strong",
	}

	// 更新miner信息
	err := dao.InsertOrUpdateMinerInfo(minerInfo)

	return err
}
