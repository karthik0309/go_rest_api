package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/karthik0309/insta_rest_api/config"
	"github.com/karthik0309/insta_rest_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var postCollection = config.GetCollections("posts")
var contex = context.Background()

func PostHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method{
		case http.MethodPost:
			CreatePost(w,r)
		case http.MethodGet:
			id:= r.URL.Query().Get("id")
			if(id==""){
				GetPosts(w,r)
			}else{
				GetPostById(w,r)
				fmt.Printf("i am here")
			}
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	var err error
	w.Header().Set("Content-Type","application/json")
	json.NewDecoder(r.Body).Decode(&post)

	if(post.Caption=="" || post.ImageURL==""){
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"Enter all the fields"}`))
		return
	}
	post.CreatedAt=time.Now().GoString()
	_,err = postCollection.InsertOne(ctx,post)
	if(err!=nil){
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"`+err.Error()+`"}`))
		return
	}

	json.NewEncoder(w).Encode(post)
}


func GetPosts(w http.ResponseWriter, r *http.Request){
	var posts []models.Post
	var err error
	var cursor *mongo.Cursor


	query := r.URL.Query()
	limit :=query["limit"]
	findOptions := options.Find()
	if(len(limit)>0){
		intLimit,_ :=strconv.ParseInt(limit[0], 6, 12)
		findOptions.SetLimit(intLimit)
	}
	w.Header().Add("Content-Type","application/json")

	cursor,err = postCollection.Find(ctx,bson.M{},findOptions)

	if(err!=nil){
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"`+err.Error()+`"}`))
		return
	}

	defer cursor.Close(ctx)

	for(cursor.Next(ctx)){
		var post models.Post
		cursor.Decode(&post)
		posts = append(posts, post)
	}

	if err:= cursor.Err() ; err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"`+err.Error()+`"}`))
		return
	}

	json.NewEncoder(w).Encode(posts)
}



func GetPostById(w http.ResponseWriter, r *http.Request){
	var err error
	var post models.Post
	w.Header().Add("Content-Type","application/json")

	query := r.URL.Query()
	reqId :=query["id"][0]
	id,err := primitive.ObjectIDFromHex(reqId)
	fmt.Printf(reqId)
	err = postCollection.FindOne(ctx,bson.M{"_id":id}).Decode(&post)

	if(err!=nil){
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"`+err.Error()+`"}`))
		return
	}

	json.NewEncoder(w).Encode(post)
}

func GetPostsByUserId(w http.ResponseWriter, r *http.Request){
	var posts []models.Post
	var err error
	var cursor *mongo.Cursor

	query := r.URL.Query()
	reqId :=query["id"][0]
	limit :=query["limit"]
	findOptions := options.Find()
	if(len(limit)>0){
		intLimit,_ :=strconv.ParseInt(limit[0], 6, 12)
		findOptions.SetLimit(intLimit)
	}
	id,err := primitive.ObjectIDFromHex(reqId)

	w.Header().Add("Content-Type","application/json")

	cursor,err = postCollection.Find(ctx,bson.M{"UID":id},findOptions)
	

	if(err!=nil){
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"`+err.Error()+`"}`))
		return
	}

	defer cursor.Close(ctx)

	for(cursor.Next(ctx)){
		var post models.Post
		cursor.Decode(&post)
		posts = append(posts, post)
	}

	if err:= cursor.Err() ; err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"`+err.Error()+`"}`))
		return
	}

	json.NewEncoder(w).Encode(posts)
}
