package userController

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"math"
	"net/http"
	model "sm-access/src/models"
	"sm-access/src/service/queryService"
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
	temporaryVariable, err := json.Marshal(user)
	err = json.Unmarshal(temporaryVariable, &response.User)
	c.AbortWithStatusJSON(http.StatusOK, response.User)
	return
}

// GetOne
// @Summary
// @Description
// @Accept json
// @Produce json
// @Tags User
// @Param id path string true "id"
// @Success 200 {object} GetOneResponse
// @Success 401 {object} models.ErrorResponse
// @Router /users/{id} [get]
func (controller userController) GetOne(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		e := "id parameter not specified"
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, model.ErrorResponse{ErrorMessage: &e, ErrorCode: http.StatusUnprocessableEntity})
		return
	}

	service := userService.UserService{}
	err, user := service.GetUser(id)
	if err != nil {
		e := err.Error()
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, model.ErrorResponse{ErrorMessage: &e, ErrorCode: http.StatusUnprocessableEntity})
		return
	}

	response := GetOneResponse{}
	temporaryVariable, _ := json.Marshal(user)
	err = json.Unmarshal(temporaryVariable, &response.User)
	c.AbortWithStatusJSON(http.StatusOK, response.User)
	return
}

// GetMany
// @Summary
// @Description
// @Accept json
// @Produce json
// @Tags User
// @Param per query int true "limit element per page" default(30)
// @Param page query int true "page number" default(1)
// @Param filter[id][eq] query string false "фильтр по полному совпадению"
// @Success 200 {object} GetManyResponse
// @Success 401 {object} models.ErrorResponse
// @Router /users [get]
func (controller userController) GetMany(c *gin.Context) {

	service := userService.UserService{}
	q, meta := queryService.GetQueries(c)

	err, users, totalCount := service.GetUsers(q.Page, q.Per)
	meta.TotalPages = int(math.Ceil(float64(totalCount) / float64(meta.Per)))
	meta.TotalItems = int(totalCount)

	if err != nil {
		e := err.Error()
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, model.ErrorResponse{ErrorMessage: &e, ErrorCode: http.StatusUnprocessableEntity})
		return
	}

	response := GetManyResponse{}
	temp, _ := json.Marshal(users)
	err = json.Unmarshal(temp, &response.Users)
	response.Meta = meta
	c.AbortWithStatusJSON(http.StatusOK, response)
	return
}

// UpdateOne
// @Summary
// @Description
// @Accept json
// @Produce json
// @Tags User
// @Param id path string true "id"
// @Param body body UpdateOneRequest true "body"
// @Success 200 {object} UpdateOneResponse
// @Success 401 {object} models.ErrorResponse
// @Router /users/{id} [put]
func (controller userController) UpdateOne(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		e := "id parameter not specified"
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, model.ErrorResponse{ErrorMessage: &e, ErrorCode: http.StatusUnprocessableEntity})
		return
	}

	service := userService.UserService{}
	err, _ := service.GetUser(id)
	if err != nil {
		e := err.Error()
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, model.ErrorResponse{ErrorMessage: &e, ErrorCode: http.StatusUnprocessableEntity})
		return
	}

	request := &UpdateOneRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		e := err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{ErrorMessage: &e, ErrorCode: http.StatusBadRequest})
		return
	}

	userData, err := json.Marshal(request)
	if err != nil {
		e := err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{ErrorMessage: &e, ErrorCode: http.StatusBadRequest})
		return
	}
	var updateUser model.User
	err = json.Unmarshal(userData, &updateUser)
	if err != nil {
		e := err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{ErrorMessage: &e, ErrorCode: http.StatusBadRequest})
		return
	}

	err = service.UpdateUser(updateUser, id)
	if err != nil {
		e := err.Error()
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, model.ErrorResponse{ErrorMessage: &e, ErrorCode: http.StatusUnprocessableEntity})
		return
	}

	_, user := service.GetUser(id)

	response := UpdateOneResponse{}
	temporaryVariable, _ := json.Marshal(user)
	err = json.Unmarshal(temporaryVariable, &response.User)
	c.AbortWithStatusJSON(http.StatusOK, response.User)
	return
}

// DeleteOne
// @Summary
// @Description
// @Accept json
// @Produce json
// @Tags User
// @Param id path string true "id"
// @Success 200 {object} DeleteOneResponse
// @Success 204 {object} models.ErrorResponse
// @Router /users/{id} [delete]
func (controller userController) DeleteOne(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		e := "id parameter not specified"
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, model.ErrorResponse{ErrorMessage: &e, ErrorCode: http.StatusUnprocessableEntity})
		return
	}

	service := userService.UserService{}
	err, _ := service.GetUser(id)
	if err != nil {
		e := err.Error()
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, model.ErrorResponse{ErrorMessage: &e, ErrorCode: http.StatusUnprocessableEntity})
		return
	}

	err = service.DeleteUser(id)
	if err != nil {
		e := err.Error()
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, model.ErrorResponse{ErrorMessage: &e, ErrorCode: http.StatusUnprocessableEntity})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, DeleteOneResponse{Message: "Ok"})
	return
}