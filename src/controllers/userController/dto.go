package userController

import (
	"sm-access/src/models"
	"sm-access/src/service/queryService"
)

type User struct {
	models.BaseModel
	Username string `json:"username"`
}

type CreateOneRequest struct {
	Username      string `json:"username"`
	PlainPassword string `json:"plainPassword"`
}

type CreateOneResponse struct {
	User User `json:"user"`
}

type GetOneResponse struct {
	User User `json:"user"`
}

type GetManyResponse struct {
	Meta  queryService.Meta `json:"meta"`
	Users []User            `json:"users"`
}

type UpdateOneRequest struct {
	Username string `json:"username"`
}

type UpdateOneResponse struct {
	User User `json:"user"`
}

type DeleteOneResponse struct {
	Message string `json:"message"`
}
