package userController

import "sm-access/src/models"

type User struct {
	models.BaseModel
	Username string `json:"username"`
}

type CreateOneRequest struct {
	Username      string `json:"username"`
	PlainPassword string `json:"plainPassword"`
}

type CreateOneResponse struct {
	User User
}
