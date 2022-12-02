package dao

import "github.com/MetaBloxIO/metablox-foundation-services/models"

func GetAllMinerInfo() ([]*models.MinerInfo, error) {
	var miners []*models.MinerInfo
	sqlStr := "select * from MinerInfo"
	rows, err := SqlDB.Queryx(sqlStr)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		miner := models.CreateMinerInfo()
		err = rows.StructScan(miner)
		if err != nil {
			return nil, err
		}
		miners = append(miners, miner)
	}
	return miners, nil
}

//
//func GetAllVirtualMinerInfo() ([]*models.MinerInfo, error) {
//	var miners []*models.MinerInfo
//	sqlStr := "select * from MinerInfo where IsVirtual = 1"
//	rows, err := SqlDB.Queryx(sqlStr)
//	if err != nil {
//		return nil, err
//	}
//
//	for rows.Next() {
//		miner := models.CreateMinerInfo()
//		err = rows.StructScan(miner)
//		if err != nil {
//			return nil, err
//		}
//		miners = append(miners, miner)
//	}
//	return miners, nil
//}
