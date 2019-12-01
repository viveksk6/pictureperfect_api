package API

import (
	"PicturePerfect2/Database"
	"PicturePerfect2/Logic"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
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

	fmt.Println("Endpoint Hit: returnAllMovies")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	queryValues := r.URL.Query()
	pageSize := queryValues.Get("pageSize")
	pageNo := queryValues.Get("pageNo")
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
	json.NewEncoder(w).Encode(movieDetails)
}

func ReturnSingleMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	movieId:= vars["movieId"]
	movieDetails:=Logic.GetMovieById(movieId)
	fmt.Println("Endpoint Hit: returnSingleMovie")
	json.NewEncoder(w).Encode(movieDetails)
}

func PostRating(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: PostRating")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var review Database.ReviewDetails
	data, err := ioutil.ReadAll(r.Body)
	//fmt.Printf()
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &review)
	vars := mux.Vars(r)
	review.MovieId, err = strconv.Atoi(vars["movieId"])
	if err!= nil {
		panic(err.Error())
	}
	review.UserId, err = strconv.Atoi(vars["userId"])
	if err!= nil {
		panic(err.Error())
	}
	//a:=fmt.Sprint(review.Rating)
	//fmt.Printf(a)
	Logic.AddRating(review)
}

func GetUserRating(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Endpoint Hit: GetUserRating")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	rating:= Logic.ReturnUserRating(vars["movieId"], vars["userId"])
	json.NewEncoder(w).Encode(rating)
}

func DeleteUserRating(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Endpoint Hit: DeleteUserRating")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	Logic.DeleteUserRating(vars["movieId"], vars["userId"])
}

func GetRating(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Endpoint Hit: GetRating")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	rating:= Logic.ReturnRating(vars["movieId"])
	json.NewEncoder(w).Encode(rating)
}