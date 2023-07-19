package deviceController

import (
	"sm-access/src/models"
	"sm-access/src/service/queryService"
)

type Device struct {
	models.Device
}

type GetManyResponse struct {
	Meta queryService.Meta `json:"meta"`
	Devices []Device `json:"devices"`
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
