package controllers

import (
	"sm-access/src/controllers/deviceController"
)

func InitControllers() {
	deviceController.DeviceController = deviceController.NewDeviceController()
}
