package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"fmt"
	"strconv"
	"github.com/karthik0309/insta_rest_api/config"
	"github.com/karthik0309/insta_rest_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var collection = config.GetCollections("users")
var ctx = context.Background()

func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", fmt.Errorf("failed to hash password: %w", err)
    }
    return string(hashedPassword), nil
}

func UserHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method{
		case http.MethodPost:
			CreateUser(w,r)
		case http.MethodGet:
			id:= r.URL.Query().Get("id")
			if(id==""){
				GetUsers(w,r)
			}else{
				GetUserById(w,r)
			}
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var err error
	w.Header().Set("Content-Type","application/json")
	json.NewDecoder(r.Body).Decode(&user)

	if(user.Name=="" || user.Email=="" || user.HashedPassword==""){
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"Enter all the fields"}`))
		return
	}

	hashPassword,passErr:=HashPassword(user.HashedPassword)
	if(passErr!=nil){
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"`+err.Error()+`"}`))
		return
	}

	user.HashedPassword=hashPassword

	_,err = collection.InsertOne(ctx,user)
	if(err!=nil){
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"`+err.Error()+`"}`))
		return
	}

	user.HashedPassword=""
	json.NewEncoder(w).Encode(user)
}


func GetUsers(w http.ResponseWriter, r *http.Request){
	var users []models.User
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

	cursor,err = collection.Find(ctx,bson.M{},findOptions)

	if(err!=nil){
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"`+err.Error()+`"}`))
		return
	}

	defer cursor.Close(ctx)

	for(cursor.Next(ctx)){
		var user models.User
		cursor.Decode(&user)
		user.HashedPassword=""
		users = append(users, user)
	}

	if err:= cursor.Err() ; err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"`+err.Error()+`"}`))
		return
	}


	json.NewEncoder(w).Encode(users)
}



func GetUserById(w http.ResponseWriter, r *http.Request){
	var err error
	var user models.User
	w.Header().Add("Content-Type","application/json")

	query := r.URL.Query()
	reqId :=query["id"][0]
	id,err := primitive.ObjectIDFromHex(reqId)
	err = collection.FindOne(ctx,bson.M{"_id":id}).Decode(&user)

	if(err!=nil){
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"`+err.Error()+`"}`))
		return
	}

	user.HashedPassword=""
	json.NewEncoder(w).Encode(user)
}
