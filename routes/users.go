package route

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"harshaweb.com/restful/models"
)

//Get all users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []models.User

	//Connect to database
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://login:" + "7388139606" + "@cluster0.dmjprkg.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	//Get collection
	collection := client.Database("restful").Collection("users")

	//Find all
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	//Loop through
	for cursor.Next(ctx) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}

	//Return
	json.NewEncoder(w).Encode(users)
}

//Get by id
func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Get params
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	//Connect to database
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://login:" + "7388139606" + "@cluster0.dmjprkg.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	//Get collection
	collection := client.Database("restful").Collection("users")

	//Find by id
	var user models.User
	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	//Return
	json.NewEncoder(w).Encode(user)
}

//Create
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Connect to database
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://login:" + "7388139606" + "@cluster0.dmjprkg.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	//Get collection
	collection := client.Database("restful").Collection("users")

	//Create
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	//Hash password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	user.Password = string(hashedPassword)

	//Insert
	result, _ := collection.InsertOne(ctx, user)
	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"message": "User created",
			"status":  "success",
			"data":    result,
		},
	)
}

//Update
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Get params
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	//Connect to database
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://login:" + "7388139606" + "@cluster0.dmjprkg.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	//Get collection
	collection := client.Database("restful").Collection("users")

	//Update
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	result, _ := collection.ReplaceOne(ctx, bson.M{"_id": id}, user)
	json.NewEncoder(w).Encode(result)
}

//Delete
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Get params
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	//Connect to database
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://login:" + "7388139606" + "@cluster0.dmjprkg.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	//Get collection
	collection := client.Database("restful").Collection("users")

	//Delete
	result, _ := collection.DeleteOne(ctx, bson.M{"_id": id})
	json.NewEncoder(w).Encode(result)
}
