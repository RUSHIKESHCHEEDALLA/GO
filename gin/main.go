package main

import (
	"net/http"
	"strconv"

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
		context.JSON(http.StatusInternalServerError, "error ")
		return
	}

	blog.Save()
	// println(Blogs.GetID() - 1)
	blog.Id = Blogs.GetID() - 1
	context.JSON(http.StatusCreated, blog)

}

func getDataById(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	context.JSON(http.StatusCreated, id)
	data, found := Blogs.GetBlogByID(id)
	if !found {
		context.JSON(http.StatusNotFound, "error Blog not found")
		return
	}

	context.JSON(http.StatusFound, data)

}

func updateData(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	context.JSON(http.StatusCreated, id)
	var updatedblog Blogs.Blog
	err := context.ShouldBindJSON(&updatedblog)

	if err != nil {
		context.JSON(http.StatusInternalServerError, "error ")
		return
	}
	updatedblog.Id = id

	data, found := Blogs.UpdateBlogByID(id, updatedblog)
	if !found {
		context.JSON(http.StatusNotFound, "error Blog not found")
		return
	}
	context.JSON(http.StatusFound, data)

}

func deleteData(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	context.JSON(http.StatusCreated, id)

	found := Blogs.DeleteBlogById(id)
	if !found {
		context.JSON(http.StatusNotFound, "error Blog not found")
		return
	}
	context.JSON(http.StatusFound, "blog deleted ")

}

func main() {
	ser := gin.Default()
	ser.GET("/getblogs", getData)
	ser.GET("/getblogs/:id", getDataById)
	ser.POST("/postblogs", postData)
	ser.PUT("/updateblogs/:id", updateData)
	ser.DELETE("/deleteblogs/:id", deleteData)
	ser.Run(":8081")
}
