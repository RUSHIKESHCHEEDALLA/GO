package main

import (
	"net/http"
	"strconv"
	"time"

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
	blog.CreatedAt = time.Now()
	blog.UpdatedAt = time.Now()
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
	updatedblog.UpdatedAt = time.Now()

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
		context.JSON(http.StatusNotFound, gin.H{"404": "error Blog not found"})
		return
	}
	context.JSON(http.StatusFound, gin.H{"201": "blog delted succesfully"})

}

func main() {
	ser := gin.Default()
	ser.GET("/get", getData)
	ser.GET("/get/:id", getDataById)
	ser.POST("/post", postData)
	ser.PUT("/update/:id", updateData)
	ser.DELETE("/delete/:id", deleteData)
	ser.Run(":8081")
}
