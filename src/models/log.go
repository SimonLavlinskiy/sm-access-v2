package models

type Log struct {
	BaseModel
	DeviceId string `json:"device_id"`
	Status   string `json:"status"`
	UserId   string `json:"user_id"`
}
