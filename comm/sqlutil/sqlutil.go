package sqlutil

import (
	"github.com/jmoiron/sqlx"
)

func ParseList[T interface{}](rows *sqlx.Rows) (list []*T, err error) {
	for rows.Next() {
		t := new(T)
		err = rows.StructScan(&t)
		if err != nil {
			return make([]*T, 0), err
		}
		list = append(list, t)
	}
	return list, nil
}
