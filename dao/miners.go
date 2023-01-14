package dao

import (
	"github.com/Masterminds/squirrel"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
)

func GetMinerInfo(dto *models.MinerDetailReqDTO) (*models.MinerInfoDTO, error) {

	sqlExp := squirrel.Select("*").From("MinerInfo").Limit(1)

	if dto.BSSID != "" {
		sqlExp = sqlExp.Where("find_in_set(?,bssid)", dto.BSSID)
	}

	sql, args, err := sqlExp.ToSql()
	if err != nil {
		return nil, err
	}
	miner := models.CreateMinerInfo()
	err = SqlDB.Get(miner, sql, args...)
	if err != nil {
		return nil, err
	}

	amount, err := SelectTotalRewardsByDID(miner.DID)
	if err != nil {
		return nil, err
	}

	m := &models.MinerInfoDTO{
		ID:             miner.ID,
		Name:           miner.Name,
		SSID:           miner.SSID,
		BSSID:          miner.BSSID,
		Longitude:      miner.Longitude,
		Latitude:       miner.Latitude,
		Availability:   miner.OnlineStatus,
		RewardEarned:   amount,
		MiningPower:    miner.MiningPower,
		IsMinable:      miner.IsMinable,
		DID:            miner.DID,
		DeviceName:     miner.DeviceName,
		Address:        miner.Location,
		SignalStrength: miner.SignalStrength,
		CreateTime:     miner.CreateTime,
	}

	return m, nil
}
