package dao

import (
	"database/sql"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
)
import "github.com/shopspring/decimal"

func GetProfitByDID(id string) (decimal.Decimal, error) {
	sqlStr := "select ifnull(sum(profit),0) from workload_record where did = ?"
	var total decimal.Decimal
	err := SqlDB.Get(&total, sqlStr, id)

	if err == sql.ErrNoRows {
		return decimal.Zero, nil
	}

	if err != nil {
		return decimal.Zero, err
	}

	return total, nil
}

func InsertWorkload(workload *models.Workload) (int64, error) {
	tx, err := SqlDB.Beginx()
	if err != nil {
		return 0, err
	}
	sqlStr := "insert into workload_record (miner,validator,qos,tracks,credential_id,model,serial,name,create_time) values (:Miner,:Validator,:Qos,:Tracks,:CredentialId,:Model,:Serial,:Name,:CreateTime)"
	result, err := tx.NamedExec(sqlStr, workload)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	newID, _ := result.LastInsertId()

	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return newID, nil
}
