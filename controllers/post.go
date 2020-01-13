package controllers

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/itanhaemprev/api/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
	"github.com/itanhaemprev/api/models"
)

func GetPosts(c *gin.Context) {
	var post models.Post
	page, err := strconv.ParseInt(c.Query("page"), 10, 64)
	if err != nil {
		page = 0
	}
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		limit = 100
	}
	posts, err, total, totalPages := post.GetPosts(page, limit)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"posts": posts, "paginate": gin.H{"total": total, "total_pages": totalPages, "page": page, "limit": limit}})
}

func GetPost(c *gin.Context) {
	var post models.Post
	var err error
	id := html.EscapeString(c.Param("id"))
	if post.ID, err = primitive.ObjectIDFromHex(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := post.GetPost(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)

}

func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := post.CreatePost(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}

func UpdatePost(c *gin.Context) {
	var post models.Post
	var err error
	id := html.EscapeString(c.Param("id"))
	if post.ID, err = primitive.ObjectIDFromHex(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := post.UpdatePost(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}

//PartialUpdatePost return a user updated
func PartialUpdatePost(c *gin.Context) {
	var post models.Post
	var err error
	if err := json.NewDecoder(c.Request.Body).Decode(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := html.EscapeString(c.Param("id"))
	if post.ID, err = primitive.ObjectIDFromHex(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := post.UpdatePost(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func DeletePost(c *gin.Context) {
	var post models.Post
	var err error
	id := html.EscapeString(c.Param("id"))
	if post.ID, err = primitive.ObjectIDFromHex(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := post.DeletePost(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func UploadImage(c *gin.Context) {
	objectId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	file, _ := c.FormFile("file")
	os.Mkdir("images", 0755)
	splitname := strings.Split(file.Filename, ".")
	ext := splitname[len(splitname)-1]
	newName := fmt.Sprintf("images/%s.%s", c.Param("id"), ext)
	if err := c.SaveUploadedFile(file, newName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	imageFile, err := os.Open(newName)
	if err != nil {
		os.Remove(newName)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := utils.ImageReduce(imageFile); err != nil {
		os.Remove(newName)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var p models.Post

	p.ID = objectId
	p.Thumbnail = "http://localhost:8080/" + newName

	if err := p.UpdatePost(); err != nil {
		os.Remove(newName)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
