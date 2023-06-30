package models

type Device struct {
	BaseModel
	Name         string `json:"name"`
	Imei         string `json:"imei"`
	SerialNumber string `json:"serial_number"`
	ProductName  string `json:"product_name"`
}
