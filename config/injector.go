package config

import (
	"github.com/google/wire"
	"go-gin-web/app/controller"
	"go-gin-web/app/repository"
	"go-gin-web/app/service"
)

var db = wire.NewSet(ConnectToDB)

var userServiceSet = wire.NewSet(service.NewUserService,
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
)

var userRepoSet = wire.NewSet(repository.NewUserRepository,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
)

var userCtrlSet = wire.NewSet(controller.NewUserController,
	wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization, db, userCtrlSet, userServiceSet, userRepoSet)
	return nil
}
