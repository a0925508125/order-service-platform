package router

import (
	"order-service-platform/service/api-gateway/router/controller"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	// Gin router
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(204) // No Content
	})
	return r
}

func SetupRouter(r *gin.Engine) {
	controllerCenter := controller.NewController()
	noLogin := r.Group("/v1")

	noLogin.POST("/order", controllerCenter.Order)
	noLogin.GET("/order", controllerCenter.GetOrder)
}
