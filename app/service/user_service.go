package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-gin-web/app/constant"
	"go-gin-web/app/entity"
	"go-gin-web/app/pkg"
	"go-gin-web/app/repository"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

type UserService interface {
	FindAll(c *gin.Context)
	FindById(c *gin.Context)
	Add(c *gin.Context)
	Update(c *gin.Context)
	DeleteById(c *gin.Context)
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepository: userRepository}
}

func (u UserServiceImpl) FindAll(c *gin.Context) {
	log.Info("find all users")
	users, err := u.userRepository.FindAll()
	if err != nil {
		log.Info("find all users fail")
	}

	c.JSON(http.StatusOK, pkg.BuildResponse("00", constant.Success, users))
}

func (u UserServiceImpl) FindById(c *gin.Context) {
	log.Info("find user by id")
	userId, _ := strconv.Atoi(c.Param("id"))

	user, err := u.userRepository.FindById(userId)
	if err != nil {
		log.Info("find user by id fail")
	}

	c.JSON(http.StatusOK, pkg.BuildResponse("00", constant.Success, user))
}

func (u UserServiceImpl) Add(c *gin.Context) {
	log.Info("add user")
	var user entity.User

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Info("mapping addUserRequest fail")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 15)
	user.Password = string(hash)

	user, err := u.userRepository.Save(&user)
	if err != nil {
		log.Info("Save user fail")
	}

	c.JSON(http.StatusOK, pkg.BuildResponse("00", constant.Success, user))
}
