package logService

import (
	log "github.com/sirupsen/logrus"
	"sm-access/config"
	"sm-access/src/models"
)

type LogService struct{}

func (l LogService) CreateLog(deviceLog models.Log) {
	db := config.Db

	err := db.Create(&deviceLog).Error

	if err != nil {
		log.Printf("Create log error: %s", err.Error())
	}
}
