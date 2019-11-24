package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/itanhaemprev/api/models"
	"net/http"
	"strconv"
)

func GetUsers(gin *gin.Context) {
	var user models.User
	page, err := strconv.ParseInt(gin.Query("page"), 10, 64)
	if err != nil {
		page = 0
	}
	limit, err := strconv.ParseInt(gin.Query("limit"), 10, 64)
	if err != nil {
		limit = 100
	}
	users, err := user.GetUsers(page, limit)
	if err != nil {
		gin.JSON(http.StatusBadRequest, err)
		return
	}
	gin.JSON(http.StatusOK, users)
}
