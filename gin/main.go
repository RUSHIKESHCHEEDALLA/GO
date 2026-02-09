package main

import (
	"fmt"
	"net/http"

	"example.com/hello/Desktop/task1/Blogs"
	"github.com/gin-gonic/gin"
)

func getData(context *gin.Context) {

	blogs := Blogs.GetAllBlogs()

	context.JSON(http.StatusOK, blogs)
}

func postData(context *gin.Context) {
	var blog Blogs.Blog
	err := context.ShouldBindJSON(&blog)

	if err != nil {
		fmt.Println("error is :", err)
		return
	}

	blog.Save()

	context.JSON(http.StatusCreated, blog)

}

func main() {
	ser := gin.Default()
	ser.GET("/posts", getData)
	ser.POST("/posts", postData)
	ser.Run(":8081")
}
