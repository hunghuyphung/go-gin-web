package config

import (
	"go-gin-web/app/controller"
	"go-gin-web/app/repository"
	"go-gin-web/app/service"
)

type Initialization struct {
	userRepo repository.UserRepository
	userSvc  service.UserService
	UserCtrl controller.UserController
}

func NewInitialization(userRepo repository.UserRepository, userService service.UserService, userCtrl controller.UserController) *Initialization {
	return &Initialization{
		userRepo: userRepo,
		userSvc:  userService,
		UserCtrl: userCtrl,
	}
}
