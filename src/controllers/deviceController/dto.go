package deviceController

import (
	"time"
)

type Device struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Imei      string    `json:"imei"`
	Type      string    `json:"type"`
	OSVersion string    `json:"os_version"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ErrorResponse struct {
	ErrorCode    int     `json:"errorCode"`
	ErrorMessage *string `json:"errorMessage"`
}

type GetDeviceRequest struct {
	Data string `json:"data"`
}

type GetManyResponse struct {
	Devices []Device `json:"devices"`
}

type GetOneResponse struct {
	Device Device `json:"device"`
}

type CreateOneRequest struct {
	Name string `json:"name" validate:"required"`
	Imei string `json:"imei" validate:"required"`
}

type CreateOneResponse struct {
	Device struct {
		Id        string    `json:"id"`
		Name      string    `json:"name"`
		Imei      string    `json:"imei"`
		Type      string    `json:"type"`
		OSVersion string    `json:"os_version"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"device"`
}

type UpdateOneRequest struct {
	Name string `json:"name"`
	Imei string `json:"imei"`
}

type UpdateOneResponse struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Imei      string    `json:"imei"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteOneResponse struct {
	Message string `json:"message"`
}
