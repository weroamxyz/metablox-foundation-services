package daily

import (
	"time"

	"github.com/metabloxDID/controllers"
	logger "github.com/sirupsen/logrus"
)

func StartDailyTimer() error {
	loc, err := time.LoadLocation("Etc/GMT+8")
	if err != nil {
		logger.Error("Failed to load location: ", err)
		return err
	}
	go func() {
		for {
			CleanupNonces()
			currentTime := time.Now().In(loc)
			nextMidnight := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 24, 0, 0, 0, loc)
			dailyTimer := time.NewTimer(nextMidnight.Sub(currentTime))
			<-dailyTimer.C
		}
	}()
	return nil
}

func CleanupNonces() {
	for ip, _ := range controllers.NonceLookup {
		delete(controllers.NonceLookup, ip)
	}
}
