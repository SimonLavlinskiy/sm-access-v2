package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"sm-access/src/models"
)

var Db *gorm.DB
var err error

func initDatabase() bool {

	mysqlUserName, _ := os.LookupEnv("MYSQL_USER")
	mysqlPassword, _ := os.LookupEnv("MYSQL_PASSWORD")
	mysqlHost, _ := os.LookupEnv("MYSQL_HOST")
	mysqlPort, _ := os.LookupEnv("MYSQL_PORT")
	mysqlDbName, _ := os.LookupEnv("MYSQL_DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", mysqlUserName, mysqlPassword, mysqlHost, mysqlPort, mysqlDbName)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	Db = Db.Debug()

	if err != nil {
		log.Fatal("DB initialization failed")
		return false
	}

	return true
}

func initMigrations() {
	if Db == nil {
		log.Panicf("Database connection not established\n")
	}

	err := Db.AutoMigrate(
		models.Device{},
		models.User{},
		models.Log{},
	)
	if err != nil {
		log.Panicf("migration error: %v", err)
	}
}
