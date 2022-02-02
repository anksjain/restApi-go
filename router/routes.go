package router

import (
	"restsample/api"

	"github.com/gin-gonic/gin"
)

type route struct {
	endPoint string
	method   string
	handler  gin.HandlerFunc
}

var routes []route

// config routers for user controller
func userRoutes(controller api.IUserController, g *gin.Engine) {

	// get user details based on id
	routes = append(routes, route{
		endPoint: "/user/:id",
		method:   "GET",
		handler:  controller.GetUser(),
	})

	// get users details for given ids
	routes = append(routes, route{
		endPoint: "/users",
		method:   "POST",
		handler:  controller.GetAllUsers(),
	})

}
