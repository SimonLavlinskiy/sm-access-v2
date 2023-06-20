package models

type Device struct {
	BaseModel
	Name string `json:"name"`
	Imei string `json:"imei"`
}
