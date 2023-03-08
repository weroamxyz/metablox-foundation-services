package main

import (
	"github.com/MetaBloxIO/did-sdk-go"
	"github.com/MetaBloxIO/metablox-foundation-services/comm/log"
	"github.com/MetaBloxIO/metablox-foundation-services/conf"
	"github.com/MetaBloxIO/metablox-foundation-services/contract"
	"github.com/MetaBloxIO/metablox-foundation-services/controller"
	"github.com/MetaBloxIO/metablox-foundation-services/dao"
	"github.com/MetaBloxIO/metablox-foundation-services/router"
	"github.com/MetaBloxIO/metablox-foundation-services/service"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	pflag.Parse()

	err := conf.Init()
	if err != nil {
		logger.Error(err)
		return
	}

	loggerConf := &log.Config{}
	viper.UnmarshalKey("logger", loggerConf)
	err = log.Init(loggerConf)
	if err != nil {
		logger.Error(err)
		return
	}

	err = dao.InitSql()
	if err != nil {
		logger.Error(err)
		return
	}

	sqlConf := &dao.Config{}
	viper.UnmarshalKey("wifiDB", sqlConf)
	err = dao.InitWifiDB(sqlConf)
	if err != nil {
		logger.Error(err)
		return
	}

	err = contract.Init()
	if err != nil {
		logger.Error(err)
		return
	}

	controller.InitializeValues()

	err = did.Init(&did.Config{
		Passphrase: viper.GetString("metablox.wallet.passphrase"),
		Keystore:   viper.GetString("metablox.wallet.keystore"),
	})
	if err != nil {
		logger.Error(err)
		return
	}

	if err = service.InitEvent(); err != nil {
		logger.Error(err)
		return
	}

	router.Setup()
}
