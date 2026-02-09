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

	_, err := io.WriteString(file, string(jsonData))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (b Blog) Save() {
	loadBlogs()

	b.Id = len(blogs) + 1
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()

	blogs = append(blogs, b)
	saveBlogsToFile()
}

func GetAllBlogs() []Blog {
	loadBlogs()
	return blogs
}
