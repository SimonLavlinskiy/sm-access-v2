package models

type User struct{
	BaseModel
	Username string `json:"username"`
	Password string `json:"password"`
}