package controllers

import (
	"net/http"
	"sesi9/middleware"
	"sesi9/routers"

	"github.com/gin-gonic/gin"
)

type InPost struct {
	Post *routers.Post
}

func (in *InPost) GetPosts(c *gin.Context) {
	if !middleware.Auth(c) {
		return
	}

	posts, err := in.Post.GetPosts("/posts")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func (in *InPost) GetPostById(c *gin.Context) {
	if !middleware.Auth(c) {
		return
	}

	id := c.Param("id")
	post, err := in.Post.GetPostId(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}

func (in *InPost) CreateNewPost(c *gin.Context) {
	if !middleware.Auth(c) {
		return
	}

	var post routers.Post
	c.BindJSON(&post)

	err := in.Post.CreateNewPost(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}
