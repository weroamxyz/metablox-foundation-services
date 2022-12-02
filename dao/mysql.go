package dao

import (
	"fmt"

	"github.com/MetaBloxIO/metablox-foundation-services/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var SqlDB *sqlx.DB

func InitSql() error {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.dbname"),
	)

	SqlDB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		logger.Error("Failed to open database: " + err.Error())
		return err
	}

	//Set the maximum number of database connections

	SqlDB.SetConnMaxLifetime(100)

	//Set the maximum number of idle connections on the database

	SqlDB.SetMaxIdleConns(10)

	//Verify connection

	if err := SqlDB.Ping(); err != nil {
		logger.Error("open database fail: ", err)
		return err
	}
	logger.Info("connect success")
	return nil
}

func Close() {
	SqlDB.Close()
}

func UploadWifiAccessVC(vc models.VerifiableCredential) (int, error) {
	tx, err := SqlDB.Beginx()
	if err != nil {
		return 0, err
	}

	sqlStr := "insert into Credentials (Type, Issuer, IssuanceDate, ExpirationDate, Description, Revoked) values ('WifiAccess', :Issuer, :IssuanceDate, :ExpirationDate, :Description, 0)"
	result, err := tx.NamedExec(sqlStr, vc)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	newID, _ := result.LastInsertId()
	sqlStr = "insert into WifiAccessInfo (CredentialID, ID, Type) values (?,?,?)"
	wifiAccessInfo := vc.CredentialSubject.(models.WifiAccessInfo)
	_, err = tx.Exec(sqlStr, newID, wifiAccessInfo.ID, wifiAccessInfo.Type)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return int(newID), nil
}

func UploadMiningLicenseVC(vc models.VerifiableCredential) (int, error) {
	tx, err := SqlDB.Beginx()
	if err != nil {
		return 0, err
	}

	sqlStr := "insert into Credentials (Type, Issuer, IssuanceDate, ExpirationDate, Description, Revoked) values ('MiningLicense', :Issuer, :IssuanceDate, :ExpirationDate, :Description, 0)"
	result, err := tx.NamedExec(sqlStr, vc)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	newID, _ := result.LastInsertId()
	sqlStr = "insert into MiningLicenseInfo (CredentialID, ID, Name, Model, Serial) values (?,?,?,?,?)"
	miningLicenseInfo := vc.CredentialSubject.(models.MiningLicenseInfo)
	_, err = tx.Exec(sqlStr, newID, miningLicenseInfo.ID, miningLicenseInfo.Name, miningLicenseInfo.Model, miningLicenseInfo.Serial)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return int(newID), nil
}

func UploadStakingVC(vc models.VerifiableCredential) (int, error) {
	tx, err := SqlDB.Beginx()
	if err != nil {
		return 0, err
	}

	sqlStr := "insert into Credentials (Type, Issuer, IssuanceDate, ExpirationDate, Description, Revoked) values ('StakingVC', :Issuer, :IssuanceDate, :ExpirationDate, :Description, 0)"
	result, err := tx.NamedExec(sqlStr, vc)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	newID, _ := result.LastInsertId()
	sqlStr = "insert into StakingVCInfo (CredentialID, ID) values (?,?)"
	stakingInfo := vc.CredentialSubject.(models.StakingVCInfo)
	_, err = tx.Exec(sqlStr, newID, stakingInfo.ID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return int(newID), nil
}

func UpdateVCExpirationDate(id, expirationDate string) error {
	sqlStr := "update Credentials set ExpirationDate = ? where ID = ?"
	_, err := SqlDB.Exec(sqlStr, expirationDate, id)
	if err != nil {
		return err
	}
	return nil
}

func RevokeVC(id string) error {
	sqlStr := "update Credentials set Revoked = 1 where ID = ?"
	_, err := SqlDB.Exec(sqlStr, id)
	if err != nil {
		return err
	}
	return nil
}

func GetCredentialStatusByID(id string) (bool, error) {
	sqlStr := "select Revoked from Credentials where ID = ?"
	var revoked bool
	err := SqlDB.Get(&revoked, sqlStr, id)
	if err != nil {
		return false, err
	}

	return revoked, nil
}

func CheckWifiAccessForExistence(id string) (bool, error) {
	var count int
	sqlStr := "select count(*) from WifiAccessInfo where ID = ?"
	err := SqlDB.Get(&count, sqlStr, id)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func GetWifiAccessFromDB(id string) (*models.VerifiableCredential, error) {
	wifiInfo := models.CreateWifiAccessInfo()
	sqlStr := "select * from WifiAccessInfo where ID = ?"
	err := SqlDB.Get(wifiInfo, sqlStr, id)
	if err != nil {
		return nil, err
	}

	vc := models.CreateVerifiableCredential()
	sqlStr = "select ID, Issuer, IssuanceDate, ExpirationDate, Description, Revoked from Credentials where ID = ?"
	err = SqlDB.Get(vc, sqlStr, wifiInfo.CredentialID)
	if err != nil {
		return nil, err
	}

	vc.CredentialSubject = *wifiInfo
	return vc, nil
}

func CheckMiningLicenseForExistence(id string) (bool, error) {
	var count int
	sqlStr := "select count(*) from MiningLicenseInfo where ID = ?"
	err := SqlDB.Get(&count, sqlStr, id)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func GetMiningLicenseFromDB(id string) (*models.VerifiableCredential, error) {
	miningInfo := models.CreateMiningLicenseInfo()
	sqlStr := "select * from MiningLicenseInfo where ID = ?"
	err := SqlDB.Get(miningInfo, sqlStr, id)
	if err != nil {
		return nil, err
	}

	vc := models.CreateVerifiableCredential()
	sqlStr = "select ID, Issuer, IssuanceDate, ExpirationDate, Description, Revoked from Credentials where ID = ?"
	err = SqlDB.Get(vc, sqlStr, miningInfo.CredentialID)
	if err != nil {
		return nil, err
	}

	vc.CredentialSubject = *miningInfo
	return vc, nil
}

func CheckStakingVCForExistence(id string) (bool, error) {
	var count int
	sqlStr := "select count(*) from StakingVCInfo where ID = ?"
	err := SqlDB.Get(&count, sqlStr, id)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func GetStakingVCFromDB(id string) (*models.VerifiableCredential, error) {
	stakingInfo := models.CreateStakingVCInfo()
	sqlStr := "select * from StakingVCInfo where ID = ?"
	err := SqlDB.Get(stakingInfo, sqlStr, id)
	if err != nil {
		return nil, err
	}

	vc := models.CreateVerifiableCredential()
	sqlStr = "select ID, Issuer, IssuanceDate, ExpirationDate, Description, Revoked from Credentials where ID = ?"
	err = SqlDB.Get(vc, sqlStr, stakingInfo.CredentialID)
	if err != nil {
		return nil, err
	}

	vc.CredentialSubject = *stakingInfo
	return vc, nil
}

func GetMinerList() ([]models.MinerInfo, error) {
	sqlStr := "select * from MinerInfo"
	rows, err := SqlDB.Queryx(sqlStr)
	if err != nil {
		return nil, err
	}

	var miners []models.MinerInfo

	for rows.Next() {
		miner := models.CreateMinerInfo()
		err = rows.StructScan(miner)
		if err != nil {
			return nil, err
		}
		miners = append(miners, *miner)
	}
	return miners, nil
}
