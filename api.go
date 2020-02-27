package main

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	client         *mongo.Client
	dbHandler      = new(DbHandler)
	dbName         = "MyDatabase"
	userCollection = "User"
)

func init() {

	client = dbHandler.connectWithMongoDB()
}

func getUserByJSON(response http.ResponseWriter, request *http.Request) {

	var users []User
	collection := client.Database(dbName).Collection(userCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message" : "` + err.Error() + `}`))
		return
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user User
		cursor.Decode(&user)
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message" : "` + err.Error() + `}`))
		return
	}

	json.NewEncoder(response).Encode(users)
}

func getUserByXML(response http.ResponseWriter, request *http.Request) {

	var users []User
	collection := client.Database(dbName).Collection(userCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message" : "` + err.Error() + `}`))
		return
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user User
		cursor.Decode(&user)
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message" : "` + err.Error() + `}`))
		return
	}

	result, _ := xml.MarshalIndent(users, "", " ")

	response.Write(result)

}

// HandlerRequest ....
func HandlerRequest() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/json/users", getUserByJSON)
	router.HandleFunc("/api/xml/users", getUserByXML)

	log.Panic(http.ListenAndServe(":1818", router))

}

func main() {
	fmt.Println("MongoDB API Started ")
	HandlerRequest()
}
