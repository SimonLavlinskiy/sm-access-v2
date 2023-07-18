package webServer

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
	"sm-access/docs"
	"sm-access/src/controllers/deviceController"
	"sm-access/src/controllers/userController"
)

func InitServer() *gin.Engine {

	switch os.Getenv("APP_ENV_MODE") {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	g := gin.Default()

	g.Use(Options)

	docs.SwaggerInfo.BasePath = "/api/v1"

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := g.Group("/api")
	{
		v1 := api.Group("v1")
		{
			device := v1.Group("devices")
			{
				device.GET(":id", deviceController.DeviceController.GetOne)
				device.GET("", deviceController.DeviceController.GetMany)
				device.POST("", deviceController.DeviceController.CreateOne)
				device.PUT(":id", deviceController.DeviceController.UpdateOne)
				device.DELETE(":id", deviceController.DeviceController.DeleteOne)
			}
			user := v1.Group("users")
			{
				user.POST("", userController.UserController.CreateOne)
			}
		}

	}

	return g
}

func Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		// FIXIT
		// TODO
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		c.Header("Allow", "GET,POST,OPTIONS")
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, pragma, origin, content-type, accept, cache-control")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(http.StatusOK)
	}
}
