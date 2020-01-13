package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/itanhaemprev/api/auth"
)

func authRouter(gin *gin.Engine) {
	authRoute := gin.Group("/auth")
	authRoute.POST("login", auth.MidleWare().LoginHandler)
	authRoute.GET("refresh_token", auth.MidleWare().RefreshHandler)
}
