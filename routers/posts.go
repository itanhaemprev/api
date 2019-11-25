package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/itanhaemprev/api/controllers"
)

func PostsRouter(gin *gin.Engine) {
	gin.GET("/posts", controllers.GetPosts)
	gin.POST("/posts", controllers.CreatePost)
}
