package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Article struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	// log.Fatal(http.ListenAndServe(":8080", nil))
	err := http.ListenAndServe(":3130", nil)
	if err != nil {
		fmt.Println("error is :", err)
		return
	}

}

func main() {
	fmt.Println("server started")
	Articles = []Article{
		Article{Id: 100, Title: "ibm", Content: "ibm", Author: "ibm"},
	}
	handleRequests()

}
