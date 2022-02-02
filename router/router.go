package router

import (
	"restsample/middleware"
	"restsample/services"

	"github.com/gin-gonic/gin"
)

func Init(dependencies services.Services) *gin.Engine {
	router := gin.Default()

	// middleware if need
	router.Use(middleware.CustomMiddleware())

	// adding user controller to routes
	userRoutes(dependencies.UserApi(), router)

	// setting handler to each routes
	for _, x := range routes {
		router.Handle(x.method, x.endPoint, x.handler)
	}
	return router
}
