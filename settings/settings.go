package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	confPath string
	port     string
)

func init() {
	pflag.StringVarP(&confPath, "conf", "c", "", "config file path")
}

func Init() (err error) {

	viper.SetConfigType("yaml")
	// the "config.yaml" in config/ folder has higher priority,then ""config.yaml"" in root folder
	viper.AddConfigPath("config/")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.ReadInConfig()

	if confPath != "" {
		fmt.Println("loading external configuration...")
		viper.SetConfigFile(confPath)
		err = viper.MergeInConfig()
		if err != nil {
			return err
		}
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		logger.Info("config file has been changed")
	})
	return nil
}
