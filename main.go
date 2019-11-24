package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/itanhaemprev/api/routers"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	routers.UsersRouter(r)
	r.Run(":8080")
}