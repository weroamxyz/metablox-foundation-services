package dao

import (
	"github.com/Masterminds/squirrel"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/jmoiron/sqlx"
	logger "github.com/sirupsen/logrus"
)

func SelectRewardRecordPage(req *models.AppRewardsPageReqDTO) ([]*models.AppRewardsPageDTO, int64, error) {
	var (
		list  = make([]*models.AppRewardsPageDTO, 0)
		count int64
	)

	sqlData := squirrel.Select("id,biz_date,rewards,is_withdrawn").From("rewards_record").OrderBy("id desc")
	sqlCount := squirrel.Select("count(*)").From("rewards_record")

	if req.Did != "" {
		sqlData = sqlData.Where("did=?", req.Did)
		sqlCount = sqlCount.Where("did=?", req.Did)
	}

	if req.BizDate != "" {
		sqlData = sqlData.Where("date(biz_date)=date(?)", req.BizDate)
		sqlCount = sqlCount.Where("date(biz_date)=date(?)", req.BizDate)
	}

	if req.UserType != "" {
		sqlData = sqlData.Where("user_type=?", req.UserType)
		sqlCount = sqlCount.Where("user_type=?", req.UserType)
	}

	sqlCountStr, sqlCountArgs, err := sqlCount.ToSql()
	if err != nil {
		return list, count, err
	}

	if err = SqlDB.QueryRow(sqlCountStr, sqlCountArgs...).Scan(&count); err != nil {
		return list, count, err
	}

	if count == 0 {
		return list, count, err
	}

	offset, limit := req.Page.PageInfo()
	sqlData = sqlData.OrderBy("id desc").Limit(limit).Offset(offset)

	sql, args, err := sqlData.ToSql()
	if err != nil {
		return list, count, err
	}
	var rows *sqlx.Rows
	if rows, err = SqlDB.Queryx(sql, args...); err != nil {
		return list, count, err
	}
	defer rows.Close()

	for rows.Next() {
		record := &models.AppRewardsPageDTO{}
		err = rows.StructScan(&record)
		if err != nil {
			logger.Warn("Scan failed: ", err)
			return make([]*models.AppRewardsPageDTO, 0), 0, err
		}
		list = append(list, record)
	}
	return list, count, nil

}

func SelectAppTotalRewards(dto *models.AppTotalRewardsReqDTO) (*models.AppTotalRewardsDTO, error) {

	sqlData := squirrel.Select(`
			round(ifnull(sum(if(is_withdrawn=0,rewards,0)),0),2) as rewardsBalance,
			round(ifnull(sum(if(is_withdrawn=1,rewards,0)),0),2) as totalWithdrawn,
			max(if(is_withdrawn=1,withdrawal_time,null)) as latestWithdrawalTime
			`).From("rewards_record")

	// query conditions
	if dto.Did != "" {
		sqlData = sqlData.Where("did = ?", dto.Did)
	}
	if dto.UserType != "" {
		sqlData = sqlData.Where("user_type = ?", dto.UserType)
	}

	sqlStr, args, err := sqlData.ToSql()
	if err != nil {
		return nil, err
	}

	data := &models.AppTotalRewardsDTO{}
	if err = SqlDB.Get(data, sqlStr, args...); err != nil {
		return nil, err
	}
	return data, nil
}
