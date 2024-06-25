package route

import (
	"github.com/gin-gonic/gin"
	"go-gin-web/config"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		user := api.Group("/user")
		user.GET("", init.UserCtrl.FindAll)
		user.POST("", init.UserCtrl.Add)
		user.GET("/:userID", init.UserCtrl.FindById)
		user.PUT("/:userID", init.UserCtrl.Update)
		user.DELETE("/:userID", init.UserCtrl.DeleteById)
	}

	return router
}
