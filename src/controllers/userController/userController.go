package userController

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	model "sm-access/src/models"
	userService "sm-access/src/service/userSevice"
)

var UserController *userController

type userController struct {
}

func NewUserController() *userController {
	return &userController{}
}

// CreateOne
// @Summary
// @Description
// @Accept json
// @Produce json
// @Tags User
// @Param body body CreateOneRequest true "body"
// @Success 200 {object} CreateOneResponse
// @Success 401 {object} models.ErrorResponse
// @Router /users [post]
func (controller userController) CreateOne(c *gin.Context) {
	request := &CreateOneRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := userService.UserService{}
	baseModel := model.BaseModel{
		ID: (uuid.New()).String(),
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.PlainPassword), bcrypt.DefaultCost)
	if err != nil {
		e := err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{ErrorMessage: &e, ErrorCode: http.StatusBadRequest})
		return
	}

	user := model.User{
		BaseModel: baseModel,
		Username:  request.Username,
		Password:  string(password),
	}

	user, err = service.CreateUser(user)
	if err != nil {
		e := err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{ErrorMessage: &e, ErrorCode: http.StatusBadRequest})
		return
	}

	response := CreateOneResponse{}
	temporaryVariable, _ := json.Marshal(user)
	err = json.Unmarshal(temporaryVariable, &response)
	c.AbortWithStatusJSON(http.StatusOK, user)
	return
}
