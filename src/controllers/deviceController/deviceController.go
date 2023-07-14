package deviceController

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"math"
	"net/http"
	model "sm-access/src/models"
	"sm-access/src/service/deviceService"
)

var DeviceController *deviceController

type deviceController struct {
}

func NewDeviceController() *deviceController {
	return &deviceController{}
}

// GetMany
// @Summary
// @Description
// @Accept json
// @Produce json
// @Tags Device
// @Param per query int true "limit element per page" default(30)
// @Param page query int true "page number" default(1)
// @Param filter[id][eq] query string false "фильтр по полному совпадению"
// @Param filter[name][includes] query string false "фильтр по неполному совпадению"
// @Success 200 {object} GetManyResponse
// @Success 401 {object} ErrorResponse
// @Router /devices [get]
func (controller deviceController) GetMany(c *gin.Context) {

	service := deviceService.DeviceService{}

	q, meta := GetQueries(c)

	err, device, totalCount := service.GetDevices(q.Page, q.Per)
	meta.TotalPages = int(math.Ceil(float64(totalCount) / float64(meta.Per)))
	if err != nil {
		e := err.Error()
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, ErrorResponse{ErrorMessage: &e, ErrorCode: http.StatusUnprocessableEntity})
		return
	}

	response := GetManyResponse{}
	temp, _ := json.Marshal(device)
	err = json.Unmarshal(temp, &response.Devices)

	c.AbortWithStatusJSON(http.StatusOK, GetManyResponse{Devices: response.Devices})
	return
}

// GetOne
// @Summary
// @Description
// @Accept json
// @Produce json
// @Tags Device
// @Param id path string true "id"
// @Success 200 {object} GetOneResponse
// @Success 401 {object} ErrorResponse
// @Router /devices/{id} [get]
func (controller deviceController) GetOne(c *gin.Context) {

	param := c.Param("id")

	if param == "" {
		e := "id parameter not specified"
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, ErrorResponse{ErrorMessage: &e, ErrorCode: http.StatusUnprocessableEntity})
		return
	}

	service := deviceService.DeviceService{}

	err, device := service.GetDevice(param)

	if err != nil {
		e := err.Error()
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, ErrorResponse{ErrorMessage: &e, ErrorCode: http.StatusUnprocessableEntity})
		return
	}
	response := GetOneResponse{}
	temporaryVariable, _ := json.Marshal(device)
	err = json.Unmarshal(temporaryVariable, &response.Device)

	c.AbortWithStatusJSON(http.StatusOK, GetOneResponse{Device: response.Device})
	return
}

// CreateOne
// @Summary
// @Description
// @Accept json
// @Produce json
// @Tags Device
// @Param body body CreateOneRequest true "body"
// @Success 200 {object} CreateOneResponse
// @Success 401 {object} ErrorResponse
// @Router /devices [post]
func (controller deviceController) CreateOne(c *gin.Context) {

	request := &CreateOneRequest{}

	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service := deviceService.DeviceService{}
	baseModel := model.BaseModel{
		ID: (uuid.New()).String(),
	}

	device := model.Device{
		BaseModel: 		baseModel,
		Name:      		request.Name,
		Imei:      		request.Imei,
		Type:      		request.Type,
		OSVersion: 		request.OSVersion,
		IsConnected: 	false,
	}

	device, err := service.CreateDevice(device)
	if err != nil {
		e := err.Error()
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, ErrorResponse{ErrorMessage: &e, ErrorCode: http.StatusUnprocessableEntity})
		return
	}

	response := CreateOneResponse{}
	temporaryVariable, _ := json.Marshal(device)
	err = json.Unmarshal(temporaryVariable, &response.Device)
	c.AbortWithStatusJSON(http.StatusOK, CreateOneResponse{Device: response.Device})
	return
}

// UpdateOne
// @Summary
// @Description
// @Accept json
// @Produce json
// @Tags Device
// @Param id path string true "id"
// @Param body body UpdateOneRequest true "body"
// @Success 200 {object} UpdateOneResponse
// @Success 401 {object} ErrorResponse
// @Router /devices/{id} [put]
func (controller deviceController) UpdateOne(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		e := "id parameter not specified"
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, ErrorResponse{ErrorMessage: &e, ErrorCode: http.StatusUnprocessableEntity})
		return
	}

	service := deviceService.DeviceService{}

	err, _ := service.GetDevice(id)
	if err != nil {
		e := err.Error()
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, ErrorResponse{ErrorMessage: &e, ErrorCode: http.StatusUnprocessableEntity})
		return
	}

	request := &UpdateOneRequest{}

	updateDevice := model.Device{
		Name: request.Name,
		Imei: request.Imei,
	}

	err = service.UpdateDevice(updateDevice, id)
	if err != nil {
		e := err.Error()
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, ErrorResponse{ErrorMessage: &e, ErrorCode: http.StatusUnprocessableEntity})
		return
	}

	_, device := service.GetDevice(id)

	response := UpdateOneResponse{}
	temporaryVariable, _ := json.Marshal(device)
	err = json.Unmarshal(temporaryVariable, &response)
	c.AbortWithStatusJSON(http.StatusOK, UpdateOneResponse{Id: device.ID, Name: device.Name, Imei: device.Imei})
	return

}

// DeleteOne
// @Summary
// @Description
// @Accept json
// @Produce json
// @Tags Device
// @Param id path string true "id"
// @Success 200 {object} DeleteOneResponse
// @Success 204 {object} ErrorResponse
// @Router /devices/{id} [delete]
func (controller deviceController) DeleteOne(c *gin.Context) {
	service := deviceService.DeviceService{}

	id := c.Param("id")

	if id == "" {
		e := "id parameter not specified"
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, ErrorResponse{ErrorMessage: &e, ErrorCode: http.StatusUnprocessableEntity})
		return
	}

	err, _ := service.GetDevice(id)
	if err != nil {
		e := err.Error()
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, ErrorResponse{ErrorMessage: &e, ErrorCode: http.StatusUnprocessableEntity})
		return
	}

	err = service.DeleteDevice(id)
	if err != nil {
		e := err.Error()
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, ErrorResponse{ErrorMessage: &e, ErrorCode: http.StatusUnprocessableEntity})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, DeleteOneResponse{Message: "Ok"})
	return
}
