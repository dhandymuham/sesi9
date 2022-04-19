package main

import (
	"sesi9/controllers"
	"sesi9/routers"

	"github.com/gin-gonic/gin"
)

var (
	port = ":9000"
)

func main() {
	p := routers.Post{}
	controllerPost := controllers.InPost{
		Post: &p,
	}

	router := gin.Default()

	router.GET("/posts", controllerPost.GetPosts)
	router.GET("/posts/:id", controllerPost.GetPostById)
	router.POST("/posts", controllerPost.CreateNewPost)

	router.Run(port)
}
