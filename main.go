package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/itanhaemprev/api/routers"
)

func main() {
	r := gin.Default()

	//load midlewares
	r.Use(cors.Default())
	r.Static("/images", "./images")
	r.StaticFile("favicon.ico", "./logo.png")
	routers.APIRouter(r)
	r.Run(":8080")
}
