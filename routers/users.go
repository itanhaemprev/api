package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/itanhaemprev/api2/controllers"
)

func UsersRouter( gin *gin.Engine) {
	gin.GET("/users", controllers.GetUsers)
}