package main

import (
	"fmt"

	"github.com/gorilla/mux"
	// "harshaweb.com/restful/config"
	"harshaweb.com/restful/routes"
	"net/http"
)

func todoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func main() {
	app := mux.NewRouter()

	app.HandleFunc("/", todoIndex)
	app.HandleFunc("/post", route.GetAllPosts).Methods("GET")
	app.HandleFunc("/post/{id}", route.GetPostById).Methods("GET")
	app.HandleFunc("/post", route.CreatePost).Methods("POST")
	app.HandleFunc("/post/{id}", route.UpdatePost).Methods("PUT")
	app.HandleFunc("/post/{id}", route.DeletePost).Methods("DELETE")

	//User routes
	app.HandleFunc("/users", route.GetAllUsers).Methods("GET")
	app.HandleFunc("/user/{id}", route.GetUserById).Methods("GET")
	app.HandleFunc("/user", route.CreateUser).Methods("POST")
	app.HandleFunc("/user/{id}", route.UpdateUser).Methods("PUT")
	app.HandleFunc("/user/{id}", route.DeleteUser).Methods("DELETE")

	//Login
	app.HandleFunc("/login", route.Login).Methods("POST")

	http.ListenAndServe(":8000", app)

}
