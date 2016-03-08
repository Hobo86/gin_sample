package api

import (
	"net/http"
	"strconv"
	. "time"

	"github.com/gin-gonic/gin"
	// "github.com/jinzhu/gorm"

	"models"
)

func PostHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		panic(err)
	}

	model := models.Default(c)
	post := model.GetPostById(id)

	c.JSON(http.StatusOK, gin.H{
		"title": "Post",
		"post":  post,
	})
}

func PostsHandler(c *gin.Context) {

	userIdStr := c.Param("userId")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		panic(err)
	}

	page, err := strconv.Atoi(c.Param("p"))
	if err != nil {
		panic(err)
	}
	size, err := strconv.Atoi(c.Param("s"))
	if err != nil {
		panic(err)
	}

	model := models.Default(c)
	posts := model.GetUserPostsByUserId(userId, page, size)

	c.JSON(http.StatusOK, gin.H{
		"title": "Post",
		"posts": posts,
	})
}
