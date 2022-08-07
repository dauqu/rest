package route

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"harshaweb.com/restful/models"
)

//Post request to login
func Login(w http.ResponseWriter, r *http.Request) {
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
	collection := client.Database("restful").Collection("users")

	//Check if user exists
	var user models.User
	err = collection.FindOne(ctx, bson.M{"username": r.FormValue("username")}).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	//Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.FormValue("password"))); err != nil {
		log.Fatal(err)
	}

	//Return
	json.NewEncoder(w).Encode(user)
}
