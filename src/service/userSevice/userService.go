package userService

import (
	"sm-access/config"
	"sm-access/src/models"
	"time"
)

type Interface interface {
}

type UserService struct {
}

func (d UserService) GetUser(userId string) (error, models.User) {

	db := config.Db

	user := models.User{}

	err := db.Unscoped().
		Where("ID = ?", userId).
		Where("deleted_at is NULL").
		First(&user).
		Error

	if err != nil {
		return err, user
	}

	return nil, user

}

func (d UserService) GetUsers(page int, per int) (err error, users []models.User, totalCount int64) {

	db := config.Db

	err = db.
		Unscoped().
		Offset((page - 1) * per).
		Limit(per).
		Find(&users).
		Count(&totalCount).
		Error

	if err != nil {
		return err, nil, 0
	}

	return err, users, totalCount
}

func (d UserService) CreateUser(user models.User) (models.User, error) {

	db := config.Db

	err := db.Create(&user).Error

	if err != nil {

		return models.User{}, err
	}

	return user, nil
}

func (d UserService) DeleteUser(userId string) error {
	db := config.Db
	t := time.Now()
	user := models.User{BaseModel: models.BaseModel{DeletedAt: &t}}

	err := db.
		Where("ID = ?", userId).
		Updates(&user).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (d UserService) UpdateUser(user models.User, userId string) error {
	db := config.Db

	err := db.Unscoped().
		Where("id = ?", userId).
		Updates(&user).
		Error

	if err != nil {

		return err
	}

	return nil
}
