package controllers

import (
	"sm-access/src/controllers/deviceController"
	"sm-access/src/controllers/userController"
)

func InitControllers() {
	deviceController.DeviceController = deviceController.NewDeviceController()
	userController.UserController = userController.NewUserController()
}
