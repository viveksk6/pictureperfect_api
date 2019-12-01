package main

import (
	"PicturePerfect2/Database"
	"PicturePerfect2/Router"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)



func main() {
	fmt.Println("Go MySQL Server")
	Database.ConnectDb()
	Logger := log.New(os.Stdout, "General Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
	handler:=Router.NewHandlers(Logger)
	myRouter := mux.NewRouter().StrictSlash(true)
	handler.HandleRequests(myRouter)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
