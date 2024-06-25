package controller

import (
	"github.com/gin-gonic/gin"
	"go-gin-web/app/service"
)

type UserController interface {
	FindAll(c *gin.Context)
	FindById(c *gin.Context)
	Add(c *gin.Context)
	Update(c *gin.Context)
	DeleteById(c *gin.Context)
}

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(usrService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{userService: usrService}
}

func (u UserControllerImpl) FindAll(c *gin.Context) {
	u.userService.FindAll(c)
}

func (u UserControllerImpl) FindById(c *gin.Context) {
	u.userService.FindById(c)
}

func (u UserControllerImpl) Add(c *gin.Context) {
	u.userService.Add(c)
}

func (u UserControllerImpl) Update(c *gin.Context) {
	u.userService.Update(c)
}

func (u UserControllerImpl) DeleteById(c *gin.Context) {
	u.userService.DeleteById(c)
}
