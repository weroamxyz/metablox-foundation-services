package dao

import "database/sql"

func GetProfitByDID(id string) (bool, error) {
	sqlStr := "select ifnull(sum(),0) from work_proof where did = ?"
	var revoked bool
	err := SqlDB.Get(&revoked, sqlStr, id)

	if err == sql.ErrNoRows {
		return false, err
	}

	if err != nil {
		return false, err
	}

	return revoked, nil
}
