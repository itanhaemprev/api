package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/itanhaemprev/api/auth"
	"github.com/itanhaemprev/api/controllers"
)

func postsRouter(gin *gin.Engine) {
	posts := gin.Group("/posts")
	posts.GET("", controllers.GetPosts)
	posts.GET("/:id", controllers.GetPost)
	posts.Use(auth.MidleWare().MiddlewareFunc())
	{
		posts.POST("", controllers.CreatePost)
		posts.PUT("/:id", controllers.UpdatePost)
		posts.PATCH("/:id", controllers.PartialUpdatePost)
		posts.DELETE("/:id", controllers.DeletePost)
		posts.POST("/:id/thumbnail", controllers.UploadImage)
	}
}
