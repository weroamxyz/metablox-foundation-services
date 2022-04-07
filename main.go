package main

import (
	"github.com/metabloxDID/routers"
	"github.com/metabloxDID/settings"
	logger "github.com/sirupsen/logrus"
)

func main() {
	err := settings.Init()
	if err != nil {
		logger.Error(err)
		return
	}

	routers.Setup()
}
