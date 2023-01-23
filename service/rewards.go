package service

import (
	"github.com/MetaBloxIO/metablox-foundation-services/dao"
	"time"
)

func CheckRewardRecordByCreateTimeAndValidator(validatorDID string, time time.Time) (bool, error) {
	return dao.CheckRewardsByDIDAndValidator(validatorDID, &time)
}
