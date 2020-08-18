package controllers

import (
	"html"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/itanhaemprev/api/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUsers(c *gin.Context) {
	var user models.User

	page, err := strconv.ParseInt(c.Query("page"), 10, 64)
	if err != nil {
		page = 0
	}
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		limit = 100
	}
	users, err := user.GetUsers(page, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, users)
}
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := user.CreateUser(user)
	if err != nil {
		var merr mongo.WriteException
		merr = err.(mongo.WriteException)
		errCode := merr.WriteErrors[0].Code
		log.Println(errCode)
		c.JSON(http.StatusBadRequest, gin.H{"error": merr.WriteErrors})
		return
	}
	c.JSON(http.StatusOK, user)
}
func GetUser(c *gin.Context) {
	var user models.User
	id := html.EscapeString(c.Param("id"))
	users, err := user.GetUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, users)
}
