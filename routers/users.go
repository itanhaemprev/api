package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/itanhaemprev/api/controllers"
)

//UsersRouter is a router for user
func UsersRouter(gin *gin.Engine) {
	gin.GET("/users", controllers.GetUsers)
	gin.GET("/users/:id", controllers.GetUser)
	gin.POST("/users", controllers.CreateUser)
}
