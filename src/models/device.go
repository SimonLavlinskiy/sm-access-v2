package models

type Device struct {
	BaseModel
	Name        	string `json:"name"`
	Imei        	string `json:"imei"`
	SerialNumber	string `json:"serial_number"`
	ProductName 	string `json:"product_name"`
	Type			string `json:"type"`
	OSVersion		string `json:"os_version"`
	IsConnected 	bool   `json:"is_connected"`
	// type (IOS or Android), OSversion, isConnected (bool), + CRUD for USER
}
