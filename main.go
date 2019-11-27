package main

import (
	"PicturePerfect2/API"
	"PicturePerfect2/Database"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"

)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", API.HomePage)
	myRouter.HandleFunc("/all", API.ReturnAllMovies).Methods("GET")
	myRouter.HandleFunc("/all", API.EnableCors).Methods("OPTIONS")
	myRouter.HandleFunc("/all/{movieId}", API.ReturnSingleMovie).Methods("GET")
	myRouter.HandleFunc("/all/{movieId}", API.EnableCors).Methods("OPTIONS")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Go MySQL Server")
	Database.ConnectDb()
	handleRequests()
}
