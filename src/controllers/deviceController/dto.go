package deviceController

import (
	"sm-access/src/models"
	"time"
)

type Device struct {
	models.Device
}

type GetOneRequest struct {
	Data string `json:"data"`
}

type GetManyResponse struct {
	Devices []struct {
		Id          string    `json:"id"`
		Name        string    `json:"name"`
		Imei        string    `json:"imei"`
		Type        string    `json:"type"`
		OSVersion   string    `json:"os_version"`
		IsConnected bool      `json:"is_connected"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
}

type GetOneResponse struct {
	Device Device `json:"device"`
}

type CreateOneRequest struct {
	Name      string `json:"name" validate:"required"`
	Imei      string `json:"imei" validate:"required"`
	Type      string `json:"type" validate:"required"`
	OSVersion string `json:"os_version"`
}

type CreateOneResponse struct {
	Device Device `json:"device"`
}

type UpdateOneRequest struct {
	Name      string `json:"name"`
	Imei      string `json:"imei"`
	Type      string `json:"type"`
	OSVersion string `json:"os_version"`
}

type UpdateOneResponse struct {
	Device Device `json:"device"`
}

type DeleteOneResponse struct {
	Message string `json:"message"`
}
