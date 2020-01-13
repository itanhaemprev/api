package routers

import "github.com/gin-gonic/gin"

//APIRouter load all routers for api
func APIRouter(r *gin.Engine) {
	authRouter(r)
	usersRouter(r)
	postsRouter(r)
}
