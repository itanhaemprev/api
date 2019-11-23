package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itanhaemprev/api2/routers"
)

func main() {
	r := gin.Default()
	routers.UsersRouter(r)
	r.Run(":8080")
}