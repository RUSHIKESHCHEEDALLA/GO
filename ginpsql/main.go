package main

import (
	"net/http"
	"strconv"
	"time"

	blogs "example.com/hello/Desktop/task1/blogspsql"
	"example.com/hello/Desktop/task1/db"
	"github.com/gin-gonic/gin"
)

func getBlogs(c *gin.Context) {

	blogs, err := blogs.GetAllBlogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch blogs",
		})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

func getDataById(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	context.JSON(http.StatusCreated, id)
	data, err := blogs.GetBlogByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"404": "could not fetch event id  "})
		return
	}

	context.JSON(http.StatusFound, data)

}

func postData(context *gin.Context) {
	var blog blogs.Blog
	err := context.ShouldBindJSON(&blog)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"404": "error Blog not found"})
		return
	}
	if blog.Title == "" || blog.Content == "" || blog.Author == "" {
		context.JSON(http.StatusBadRequest, gin.H{"400": "Missing fields"})
		return
	}

	// println(Blogs.GetID() - 1)

	blog.CreatedAt = time.Now()
	blog.UpdatedAt = time.Now()
	blog.Save()
	// blog.Id = Blogs.GetID() - 1
	context.JSON(http.StatusCreated, blog)

}

func updateData(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	context.JSON(http.StatusCreated, id)

	var updatedblog blogs.Blog
	err := context.ShouldBindJSON(&updatedblog)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"500": "internal server error"})
		return
	}

	if updatedblog.Title == "" || updatedblog.Content == "" || updatedblog.Author == "" {
		context.JSON(http.StatusBadRequest, gin.H{"400": "Missing fields"})
		return
	}

	updatedblog.Id = id
	updatedblog.UpdatedAt = time.Now()

	data, err := blogs.UpdateBlogByID(updatedblog)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"404": "error Blog not found"})
		return
	}
	context.JSON(http.StatusFound, data)

}

func deleteData(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	context.JSON(http.StatusCreated, id)

	err := blogs.DeleteBlogById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"404": "error Blog not found"})
		return
	}
	context.JSON(http.StatusFound, gin.H{"200": "blog delted succesfully"})

}

func main() {

	db.InitDB()

	router := gin.Default()

	router.GET("/blog", getBlogs)
	router.GET("/blog/:id", getDataById)
	router.POST("/blog", postData)
	router.PUT("/blog/:id", updateData)
	router.DELETE("/blog/:id", deleteData)

	router.Run(":3031")
}
