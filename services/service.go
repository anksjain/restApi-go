package services

import (
	"restsample/api"
	"restsample/instance"
	"restsample/repository"
)

type services struct {
	userApi api.IUserController
}

type Services interface {
	UserApi() api.IUserController
}

func (svc *services) UserApi() api.IUserController {
	return svc.userApi
}

// service layer setting all services
func InitServices() Services {
	db := instance.GetDb()
	userController := api.NewUserController(
		repository.NewUserRepo(db),
	)
	return &services{
		userApi: userController,
	}
}
