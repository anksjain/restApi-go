package runner

import (
	"fmt"
	"net/http"
	"restsample/config"
	"restsample/router"
	"restsample/services"
	"time"
)

func Start() {
	services := services.InitServices() // intialize services
	routes := router.Init(services)     // setting up router
	s := &http.Server{
		Addr:         ":" + config.GetAppPort(),
		Handler:      routes,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Server is running on Port: ", config.GetAppPort())
	s.ListenAndServe()
}
