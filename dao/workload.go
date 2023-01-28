package dao

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/MetaBloxIO/metablox-foundation-services/comm/sqlutil"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/jmoiron/sqlx"
	"time"
)
import "github.com/shopspring/decimal"

func SelectWorkloadList(req *models.WorkloadDTOReq) ([]*models.WorkloadRecord, error) {

	var (
		list = make([]*models.WorkloadRecord, 0)
	)

	sqlData := squirrel.Select("validator,miner,date(create_time) bizDate").From("workload_record").OrderBy("id desc")

	if req.Did != "" {
		sqlData = sqlData.Where("did=?", req.Did)
	}

	if req.BizDate != "" {
		sqlData = sqlData.Where("create_time>?", req.BizDate)
	}

	sqlData = sqlData.GroupBy("bizDate").OrderBy("bizDate desc")

	sql, args, err := sqlData.ToSql()
	if err != nil {
		return list, err
	}
	var rows *sqlx.Rows
	if rows, err = SqlDB.Queryx(sql, args...); err != nil {
		return list, err
	}
	defer rows.Close()

	list, err = sqlutil.ParseList[models.WorkloadRecord](rows)
	if err != nil {
		return nil, err
	}
	return list, nil
}

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

func SelectCountWorkloadByDIDAndValidator(miner string, validator string, bizDate time.Time) (int, error) {
	sqlStr := `select count(*) from workload_record where miner =? and validator = ? and date(create_time) = date(?)`
	var count int
	if err := SqlDB.Get(&count, sqlStr, miner, validator, bizDate); err != nil {
		return 0, err
	}
	return count, nil
}

func InsertWorkload(workload *models.WorkloadRecord) (int64, error) {
	tx, err := SqlDB.Beginx()
	if err != nil {
		return 0, err
	}
	sqlStr := "insert into workload_record (miner,validator,qos,tracks,credential_id,model,serial,name,create_time) values (:miner,:validator,:qos,:tracks,:credential_id,:model,:serial,:name,:create_time)"
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
