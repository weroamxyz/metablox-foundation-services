package dao

import (
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	logger "github.com/sirupsen/logrus"
)

var wifiSqlDB *sqlx.DB

type Config struct {
	Host     string `json:"host,omitempty"`
	Port     string `json:"port,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	Dbname   string `json:"dbname,omitempty"`
}

func WifiDB() *sqlx.DB {
	return wifiSqlDB
}

func InitWifiDB(c *Config) error {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Dbname,
	)

	wifiSqlDB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		logger.Error("Failed to open database: " + err.Error())
		return err
	}

	//Set the maximum number of database connections
	wifiSqlDB.SetConnMaxLifetime(100)

	//Set the maximum number of idle connections on the database
	wifiSqlDB.SetMaxIdleConns(10)

	//Verify connection
	if err := wifiSqlDB.Ping(); err != nil {
		log.Error("open database fail: ", err)
		return err
	}
	logger.Info("connect success")
	return nil
}

func WIfiDBClose() {
	wifiSqlDB.Close()
}
