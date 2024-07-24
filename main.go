package main

import (
	"app/controllers"
	"app/middleware"
	"app/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	godotenv.Load()

	utils.InitDB()
	defer utils.CloseDB()

	router := mux.NewRouter()

	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/login", controllers.AuthorizeUser).Methods("POST")
	router.HandleFunc("/product/{name}", middleware.AuthMiddleware(controllers.GetProduct)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
