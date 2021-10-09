package main

import (
	"log"
	"net/http"
	"github.com/karthik0309/insta_rest_api/controllers"
)

func handleRequests(){
	http.HandleFunc("/users/",controllers.UserHandler)
	http.HandleFunc("/posts/",controllers.PostHandler)
	http.HandleFunc("/posts/users/",controllers.GetPostsByUserId)
	log.Fatal(http.ListenAndServe(":8081",nil))
}

func main(){
	handleRequests()
}