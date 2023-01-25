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
		BSSID:          dto.BSSID,
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

func InsertOrUpdateMinerInfo(info *models.MinerInfo) error {

	sqlStr := `INSERT INTO minerinfo ( SSID, BSSID, OnlineStatus, MiningPower, DID, IsVirtual, SignalStrength, CreateTime, IsMinable )
			VALUES( ?, ?, ?,?, ?, ?, ?, ?, ? ) 
			ON DUPLICATE KEY UPDATE SSID =VALUES( SSID ),BSSID =VALUES( BSSID ),OnlineStatus =VALUES( OnlineStatus ),
	                        MiningPower =VALUES( MiningPower ),IsVirtual =VALUES( IsVirtual ),
	                        SignalStrength =VALUES( SignalStrength ),IsMinable =VALUES(IsMinable)`

	_, err := SqlDB.Exec(sqlStr, info.SSID, info.BSSID, info.OnlineStatus, info.MiningPower, info.DID, info.IsVirtual, info.SignalStrength, info.CreateTime, info.IsMinable)
	return err

	return nil
}
