package service

import (
	"fmt"
	"github.com/MetaBloxIO/metablox-foundation-services/comm/consts"
	"github.com/MetaBloxIO/metablox-foundation-services/comm/event"
	"github.com/MetaBloxIO/metablox-foundation-services/dao"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/shopspring/decimal"
	logger "github.com/sirupsen/logrus"
	"math/rand"
)

func InitEvent() error {
	err := event.Subscribe(event.WorkloadValidated, RewardToAppUser)
	if err != nil {
		return err
	}
	return nil
}

func RewardToAppUser(wl *models.WorkloadRecord) {
	if wl == nil {
		logger.Error("workload is nil")
		return
	}

	count, err := dao.SelectCountWorkloadByDIDAndValidator(wl.Miner, wl.Validator, wl.CreateTime.UTC())
	if err != nil {
		logger.Error(err.Error())
		return
	}

	if count > 1 {
		logger.Info("rewards exists")
		return
	}

	//flag, err := CheckRewardRecordByCreateTimeAndValidator(wl.Validator, wl.CreateTime.UTC())
	//if err != nil {
	//	logger.Error(err.Error())
	//}

	//if flag {
	//	logger.Info("the user has reward today")
	//	return
	//}

	now := wl.CreateTime.UTC()

	rewards := getRewards(consts.AppRewardsPerDay, consts.AppRewardsPerDay)
	record := &models.RewardsRecord{
		BizDate:        &now,
		Did:            wl.Validator,
		UserType:       consts.AppUser,
		Rewards:        rewards,
		IsWithdrawn:    false,
		CreateTime:     &now,
		WithdrawalTime: nil,
	}

	if _, err = dao.InsertRewards(record); err != nil {
		logger.Error(fmt.Sprintf("App rewards=%s did=%s, rewards collection error", record.Rewards, record.Did), err.Error())
	}

}

// getRewards get rewards according to the user type
func getRewards(min, max float64) decimal.Decimal {
	f := min + rand.Float64()*(max-min)
	return decimal.NewFromFloat(f)
}
