package main

import (
	"github.com/MetaBloxIO/metablox-foundation-services/contract"
	"github.com/MetaBloxIO/metablox-foundation-services/controllers"
	"github.com/MetaBloxIO/metablox-foundation-services/dao"
	"github.com/MetaBloxIO/metablox-foundation-services/routers"
	"github.com/MetaBloxIO/metablox-foundation-services/settings"
	logger "github.com/sirupsen/logrus"
)

func main() {
	err := settings.Init()
	if err != nil {
		logger.Error(err)
		return
	}

	err = dao.InitSql()
	if err != nil {
		logger.Error(err)
		return
	}

	err = contract.Init()
	if err != nil {
		logger.Error(err)
		return
	}

	err = controllers.InitializeValues()
	if err != nil {
		logger.Error(err)
		return
	}

	routers.Setup()
}
