package Blogs

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type Blog struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

var filePath = "blogs.json"
var blogs = []Blog{}

func loadBlogs() {
	file, _ := os.Open(filePath)
	defer file.Close()

	data, _ := io.ReadAll(file)
	json.Unmarshal(data, &blogs)
}

func saveBlogsToFile() {
	jsonData, _ := json.MarshalIndent(blogs, "", "  ")

	file, _ := os.Create(filePath)
	defer file.Close()

	_, err := file.Write(jsonData)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetID() int {
	maxID := 0
	for _, blog := range blogs {
		if blog.Id > maxID {
			maxID = blog.Id
		}
	}
	return maxID + 1
}

func (b Blog) Save() {
	loadBlogs()

	b.Id = GetID()
	blogs = append(blogs, b)
	saveBlogsToFile()
}

func GetAllBlogs() []Blog {
	loadBlogs()
	return blogs
}

func GetBlogByID(id int) (Blog, bool) {
	loadBlogs()

	for _, blog := range blogs {
		if blog.Id == id {
			return blog, true
		}
	}

	return Blog{}, false
}

func UpdateBlogByID(id int, updated Blog) (Blog, bool) {
	loadBlogs()

	for i, blog := range blogs {
		if blog.Id == id {
			blogs[i].Title = updated.Title
			blogs[i].Content = updated.Content
			blogs[i].Author = updated.Author
			blogs[i].UpdatedAt = time.Now()

			saveBlogsToFile()
			return blogs[i], true
		}
	}

	return Blog{}, false
}

func DeleteBlogById(id int) bool {
	loadBlogs()
	for i, blog := range blogs {
		if blog.Id == id {
			blogs = append(blogs[:i], blogs[i+1:]...)
			saveBlogsToFile()
			return true
		}

	}
	return false
}
