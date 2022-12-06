package dao

import "github.com/MetaBloxIO/metablox-foundation-services/models"

func GetWifiUserInfo(username string) (*models.WifiUserInfo, error) {
	userInfo := &models.WifiUserInfo{}
	sqlStr := "select * from radcheck where username = ?"
	err := WifiDB().Unsafe().Get(userInfo, sqlStr, username)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

func InsertWifiUserInfo(userInfo *models.WifiUserInfo) (int, error) {
	tx, err := WifiDB().Beginx()
	if err != nil {
		return 0, err
	}

	sqlStr := "insert into radcheck (username,attribute,op,value) values (:username,'Cleartext-Password',':=',:value)"
	result, err := tx.NamedExec(sqlStr, userInfo)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	newID, _ := result.LastInsertId()

	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return int(newID), nil
}
