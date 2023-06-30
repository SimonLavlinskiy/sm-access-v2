package usbService

import (
	"github.com/google/gousb"
	log "github.com/sirupsen/logrus"
	"sm-access/src/models"
	"sm-access/src/service/deviceService"
)

func GetConnectedDevices() (connectedDevices []models.Device) {
	var service deviceService.DeviceService

	ctx := gousb.NewContext()
	defer ctx.Close()

	devs, err := ctx.OpenDevices(func(device *gousb.DeviceDesc) bool {
		return true
	})

	if err != nil {
		log.Printf("error OpenDevices(): %v", err)
	}

	for _, d := range devs {
		defer d.Close()

		var device models.Device

		device.ProductName, err = d.Product()
		if err != nil {
			device.ProductName = "undefined"
		}

		device.SerialNumber, err = d.SerialNumber()
		if err != nil {
			device.SerialNumber = "undefined"
		}

		device, err = service.GetOrCreateDevice(device)

		connectedDevices = append(connectedDevices, device)
	}

	return connectedDevices
}
