package route

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"harshaweb.com/restful/models"
	"log"
	"net/http"
	"time"
)

//Get all
func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var posts []models.Post

	//Connect to database
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://login:7388139606@cluster0.dmjprkg.mongodb.net/?retryWrites=true&w=majority"))
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
	collection := client.Database("restful").Collection("posts")

	//Find all
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	//Loop through
	for cursor.Next(ctx) {
		var post models.Post
		cursor.Decode(&post)
		posts = append(posts, post)
	}

	//Return
	json.NewEncoder(w).Encode(posts)
}

//Get by id
func GetPostById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Get params
	params := mux.Vars(r)

	//Connect to database
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://login:7388139606@cluster0.dmjprkg.mongodb.net/?retryWrites=true&w=majority"))
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
	collection := client.Database("restful").Collection("posts")

	//Find by id
	var post models.Post
	id, _ := primitive.ObjectIDFromHex(params["id"])
	err = collection.FindOne(ctx, models.Post{ID: id}).Decode(&post)
	if err != nil {
		log.Fatal(err)
	}

	//Return
	json.NewEncoder(w).Encode(post)
}

//Create
func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Connect to database
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://login:7388139606@cluster0.dmjprkg.mongodb.net/?retryWrites=true&w=majority"))
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
	collection := client.Database("restful").Collection("posts")

	//Create
	var post models.Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	result, err := collection.InsertOne(ctx, post)
	if err != nil {
		log.Fatal(err)
	}

	//Return
	json.NewEncoder(w).Encode(result)
}

//Update
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Get params
	params := mux.Vars(r)

	//Connect to database
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://login:7388139606@cluster0.dmjprkg.mongodb.net/?retryWrites=true&w=majority"))
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
	collection := client.Database("restful").Collection("posts")

	//Update
	var post models.Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}
	update := bson.M{"$set": post}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}

	//Return
	json.NewEncoder(w).Encode(result)
}

//Delete
func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Get params
	params := mux.Vars(r)

	//Connect to database
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://login:7388139606@cluster0.dmjprkg.mongodb.net/?retryWrites=true&w=majority"))
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
	collection := client.Database("restful").Collection("posts")

	//Delete
	id, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	//Return
	json.NewEncoder(w).Encode(result)
}
