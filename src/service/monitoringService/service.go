package monitoringService

import (
	"fmt"
	"os"
	"sm-access/src/models"
	"sm-access/src/service/logService"
	"sm-access/src/service/monitoringService/usbService"
	"sm-access/src/service/socketService"
	"strconv"
	"time"
)

var (
	lastSync map[string]models.Device
)

func UsbMonitoring(server *socketService.Server) {
	for {
		devices := usbService.GetConnectedDevices()

		lastSync = CheckDevices(server, lastSync, devices)

		timeout, _ := strconv.Atoi(os.Getenv("MONITORING_TIMEOUT"))
		time.Sleep(time.Duration(timeout) * time.Second)
	}
}

func CheckDevices(server *socketService.Server, lastSync map[string]models.Device, conDevices []models.Device) map[string]models.Device {
	var deviceLogService logService.LogService
	currentSync, newConnectedDevices := map[string]models.Device{}, map[string]models.Device{}

	for _, d := range conDevices {
		currentSync[d.SerialNumber] = d
		newConnectedDevices[d.SerialNumber] = d
	}

	// Находим новые подключенные девайсы
	for key, _ := range lastSync {
		delete(newConnectedDevices, key)
	}

	// Если остались какие либо девайсы, значит они были подключены
	for _, device := range newConnectedDevices {
		server.SendToSocket([]byte(fmt.Sprintf("%s: true", device.ID)))

		// Логируем подключение девайса
		deviceLogService.CreateLog(models.Log{DeviceId: device.ID, Status: "connected"})
	}

	// Находим отключенные девайсы
	for key, _ := range currentSync {
		delete(lastSync, key)
	}

	// Если остались какие либо девайсы, значит они были отключены
	for _, device := range lastSync {
		server.SendToSocket([]byte(fmt.Sprintf("%s: false", device.ID)))

		// Логируем отключение девайса
		deviceLogService.CreateLog(models.Log{DeviceId: device.ID, Status: "disconnected"})
	}

	return currentSync
}
