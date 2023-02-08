package dao

import (
	"fmt"

	"github.com/MetaBloxIO/metablox-foundation-services/conf"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/jmoiron/sqlx"
)

func TestDBInit() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		"tester",
		"testing",
		"127.0.0.1",
		"3306",
		"foundationservicetest",
	)
	SqlDB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	return nil
}

func ConnectToTestDB() error {
	err := conf.Init()
	if err != nil {
		return err
	}

	err = TestDBInit()
	if err != nil {
		return err
	}
	return nil
}

func CreateTestCredentialsTable() error {
	sqlStr := `CREATE TABLE foundationservicetest.Credentials (
		ID int NOT NULL AUTO_INCREMENT,
		Type varchar(100) NOT NULL,
		Issuer varchar(100) NOT NULL,
		IssuanceDate timestamp NOT NULL,
		ExpirationDate timestamp NOT NULL,
		Description varchar(100) NOT NULL,
		Revoked tinyint NOT NULL,
		PRIMARY KEY (ID)
	  )`

	_, err := SqlDB.Exec(sqlStr)

	if err != nil {
		return err
	}

	return nil

}

func DeleteTestCredentialsTable() {
	sqlStr := "DROP TABLE foundationservicetest.Credentials;"
	SqlDB.Exec(sqlStr)
}

func InsertSampleIntoCredentials(vc *models.VerifiableCredential, vcType string) error {
	sqlStr := "insert into Credentials (Type, Issuer, IssuanceDate, ExpirationDate, Description, Revoked) values (?,?,?,?,?,?);"
	_, err := SqlDB.Exec(sqlStr, vcType, vc.Issuer, vc.IssuanceDate, vc.ExpirationDate, vc.Description, vc.Revoked)
	return err
}

func RetrieveSampleFromCredentials(id string) (*models.VerifiableCredential, error) {
	sqlStr := "select ID, Issuer, IssuanceDate, ExpirationDate, Description, Revoked from Credentials where ID = ?;"
	result := models.CreateVerifiableCredential()
	err := SqlDB.Get(result, sqlStr, id)
	return result, err
}

func CreateTestMiningLicenseTable() error {
	sqlStr := `CREATE TABLE MiningLicenseInfo (
		CredentialID int NOT NULL,
		ID varchar(100) NOT NULL,
		Name varchar(100) NOT NULL,
		Model varchar(100) NOT NULL,
		Serial varchar(100) NOT NULL,
		PRIMARY KEY (CredentialID)
	  )`

	_, err := SqlDB.Exec(sqlStr)

	if err != nil {
		return err
	}

	return nil
}

func DeleteTestMiningLicenseTable() {
	sqlStr := "DROP TABLE foundationservicetest.MiningLicenseInfo;"
	SqlDB.Exec(sqlStr)
}

func RetrieveSampleFromMiningLicenseInfo(id string) (*models.MiningLicenseInfo, error) {
	sqlStr := "select * from MiningLicenseInfo where ID = ?;"
	result := models.CreateMiningLicenseInfo()
	err := SqlDB.Get(result, sqlStr, id)
	return result, err
}

func CreateTestWifiAccessTable() error {
	sqlStr := `CREATE TABLE WifiAccessInfo (
		CredentialID int NOT NULL,
		ID varchar(100) NOT NULL,
		Type enum('User','Validator') NOT NULL,
		PRIMARY KEY (CredentialID)
	  )`

	_, err := SqlDB.Exec(sqlStr)

	if err != nil {
		return err
	}

	return nil
}

func DeleteTestWifiAccessTable() {
	sqlStr := "DROP TABLE foundationservicetest.WifiAccessInfo;"
	SqlDB.Exec(sqlStr)
}

func RetrieveSampleFromWifiAccessInfo(id string) (*models.WifiAccessInfo, error) {
	sqlStr := "select * from WifiAccessInfo where ID = ?;"
	result := models.CreateWifiAccessInfo()
	err := SqlDB.Get(result, sqlStr, id)
	return result, err
}
