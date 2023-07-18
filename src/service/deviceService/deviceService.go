package deviceService

import (
	"sm-access/config"
	"sm-access/src/models"
	"time"
)

type Interface interface {
}

type DeviceService struct {
}

func (d DeviceService) GetDevice(deviceId string) (error, models.Device) {

	db := config.Db

	device := models.Device{}

	err := db.Unscoped().
		Where("ID = ?", deviceId).
		Where("deleted_at is NULL").
		First(&device).
		Error

	if err != nil {
		return err, device
	}

	return nil, device

}

func (d DeviceService) GetDevices(page int, per int) (err error, devices []models.Device, totalCount int64) {

	db := config.Db

	err = db.
		Unscoped().
		Where("deleted_at is NULL").
		Offset((page - 1) * per).
		Limit(per).
		Find(&devices).
		Count(&totalCount).
		Error

	if err != nil {
		return err, nil, 0
	}

	return err, devices, totalCount
}

func (d DeviceService) CreateDevice(device models.Device) (models.Device, error) {

	db := config.Db

	err := db.Create(&device).Error

	if err != nil {
		return models.Device{}, err
	}

	return device, nil
}

func (d DeviceService) DeleteDevice(deviceId string) error {
	db := config.Db
	t := time.Now()
	device := models.Device{BaseModel: models.BaseModel{DeletedAt: &t}}

	err := db.
		Where("ID = ?", deviceId).
		Updates(&device).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (d DeviceService) UpdateDevice(device models.Device, deviceId string) error {
	db := config.Db

	err := db.Unscoped().
		Where("id = ?", deviceId).
		Updates(&device).
		Error

	if err != nil {

		return err
	}

	return nil
}

func (d DeviceService) GetOrCreateDevice(device models.Device) (resp models.Device, err error) {
	err = config.Db.
		FirstOrCreate(
			&resp,
			models.Device{
				ProductName:  device.ProductName,
				SerialNumber: device.SerialNumber,
			}).Error

	if err != nil {
		return models.Device{}, err
	}

	return resp, nil
}
