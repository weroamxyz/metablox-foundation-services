package dao

import (
	"encoding/json"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/MetaBloxIO/metablox-foundation-services/comm/consts"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

func SelectNearbyMinersList(dto *models.MinersDTO) ([]*models.MinersWithDistanceDTO, error) {

	if dto.Latitude.IsZero() || dto.Longitude.IsZero() {
		return nil, errors.New("both longitude and latitude are required")
	}

	// max 30km
	if dto.Distance.IsZero() || dto.Distance.GreaterThan(decimal.NewFromFloat(consts.MaxDistance)) {
		dto.Distance = decimal.NewFromFloat(consts.MaxDistance)
	}
	bytes, _ := json.Marshal(dto)
	fmt.Println(string(bytes))

	sql := squirrel.Select(` *,unix_timestamp(CreateTime) createTime,
	ROUND(
    IFNULL(6378.138 * 2 * ASIN(
      SQRT(
        POW(
          SIN(
            (
              ` + dto.Latitude.String() + ` * PI() / 180 - Latitude * PI() / 180
            ) / 2
          ), 2
        ) + COS(` + dto.Latitude.String() + ` * PI() / 180) * COS(Latitude * PI() / 180) * POW(
          SIN(
            (
              ` + dto.Longitude.String() + `* PI() / 180 - Longitude * PI() / 180
            ) / 2
          ), 2
        )
      )
    ),0),2) AS distance`).From("MinerInfo").OrderBy(" distance ASC")

	sql = sql.Having("distance<=?", dto.Distance)
	var list = make([]*models.MinersWithDistanceDTO, 0)
	var rows *sqlx.Rows

	sqlStr, args, err := sql.ToSql()
	if err != nil {
		return list, err
	}
	rows, err = SqlDB.Queryx(sqlStr, args...)

	if err != nil {
		return list, err
	}
	defer rows.Close()

	for rows.Next() {
		record := &models.MinersWithDistanceDTO{}
		err = rows.StructScan(&record)
		if err != nil {
			return nil, err
		}
		list = append(list, record)
	}

	return list, nil
}
