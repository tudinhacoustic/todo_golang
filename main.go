package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}
func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	// Loop over all of our Articles
	// if the article.Id equals the key we pass in
	// return the article encoded as JSON

	// fmt.Fprintf(w, "Key: "+key)

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}
func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response contraining the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	// fmt.Println(reqBody)
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	// fmt.Fprintf(w, "%+v", string(reqBody))
	json.NewEncoder(w).Encode(article)
}
func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}
func updateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	for index, article2 := range Articles {
		if article2.Id == id {
			Articles[index] = article
		}
	}
	fmt.Println(Articles)
	json.NewEncoder(w).Encode(article)
}
func handleRequests() {
	// http.HandleFunc("/", homePage)
	// // add our articles route and map it to our
	// // returnAllArticles function like so
	// http.HandleFunc("/articles", returnAllArticles)
	// log.Fatal(http.ListenAndServe(":9090", nil))

	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	// NOTE: Ordeing is important here! This has to be defined
	// the other `/article` endpoint.
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)

	log.Fatal(http.ListenAndServe(":9090", myRouter))
}
func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}
