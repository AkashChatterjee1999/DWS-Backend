package pgsql

import (
	"fmt"
	"github.com/DWS-Backend/provider/common/helpers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PgDBConnect() (db *gorm.DB, err error) {
	dbUser := helpers.EnvConfigs.DbUser
	dbHost := helpers.EnvConfigs.DbHost
	dbName := helpers.EnvConfigs.DbName
	dbPassword := helpers.EnvConfigs.DbPassword
	dbPort := helpers.EnvConfigs.DbPort
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=disable TimeZone=Asia/Kolkata",
			dbUser, dbPassword, dbName, dbPort, dbHost),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err == nil {
		helpers.Logger.Info(nil, "Successfully started pgsql data source")
	}
	return
}
