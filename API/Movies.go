package API

import (

	"PicturePerfect2/Logic"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func EnableCors(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func HomePage(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")


	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func ReturnAllMovies(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Access-Control-Allow-Origin", "*")
	queryValues := r.URL.Query()
	pageSize := queryValues.Get("pageSize")
	pageNo := queryValues.Get("pageNo")



	if pageNo!="" && pageSize!="" {
		pageNoInt, err:= strconv.Atoi(pageNo)
		if err != nil {
			panic(err.Error())
		}
		pageSizeInt, err:= strconv.Atoi(pageSize)
		if err != nil {
			panic(err.Error())
		}
		pageNo = strconv.Itoa(pageSizeInt*(pageNoInt-1))
		movieDetails:=Logic.GetAllMovies(pageNo, pageSize)
		fmt.Println("Endpoint Hit: returnAllMovies")
		json.NewEncoder(w).Encode(movieDetails)
	} else {
		fmt.Fprintf(w, "Invalid query parameters")
	}


}

func ReturnSingleMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	movieId:= vars["movieId"]
	movieDetails:=Logic.GetMovieById(movieId)
	fmt.Println("Endpoint Hit: returnSingleMovie")
	json.NewEncoder(w).Encode(movieDetails)
}