package API

import (
	"PicturePerfect2/Database/Catalogue"
	"PicturePerfect2/Database/Review"
	"PicturePerfect2/Database/Shows"
	"PicturePerfect2/Logic"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Credentials struct {
	UserId int `json:"userId"`
	Username string `json:"username"`
	Password string `json:"password"`
	NewPassword string `json:"new_password"`
}

func EnableCors(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func HomePage(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func Login(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("Endpoint Hit: Login")
	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	var creds Credentials
	json.Unmarshal(data, &creds)
	tokenString, expirationTime, er:=Logic.GenerateJWT(creds.Username,creds.Password)
	if er==http.StatusUnauthorized{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if er == http.StatusInternalServerError{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func ResetPassword(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: ResetPassword")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var creds Credentials
	json.Unmarshal(data, &creds)
	res:=Logic.ResetPassword(creds.UserId, creds.Password, creds.NewPassword)
	if res != ""{
		fmt.Fprint(w,res)
	}
}

func PostMovie(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: Add Movie")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var movie Catalogue.MovieDetails
	data, err := ioutil.ReadAll(r.Body)
	//fmt.Printf()
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &movie)
	Logic.AddMovie(movie)
}

func PutMovie(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: Put Movie")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var movie Catalogue.MovieDetails
	data, err := ioutil.ReadAll(r.Body)
	//fmt.Printf()
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &movie)
	vars := mux.Vars(r)
	movie.ID,err= strconv.Atoi(vars["movieId"])
	if err != nil {
		panic(err)
	}
	Logic.UpdateMovie(movie)
}

func PatchMovie(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: Patch Movie")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	data, err := ioutil.ReadAll(r.Body)
	var movie map[string]string
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &movie)
	vars := mux.Vars(r)
	movieId,err := strconv.Atoi(vars["movieId"])
	if err != nil {
		panic(err)
	}
	Logic.UpdateMovieAttribute(movieId,movie)
}

func ReturnAllMovies(w http.ResponseWriter, r *http.Request){

	fmt.Println("Endpoint Hit: returnAllMovies")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	queryValues := r.URL.Query()
	pageSize := queryValues.Get("pageSize")
	pageNo := queryValues.Get("pageNo")
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
	var review Review.ReviewDetails
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

func DeleteUserRatings(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Endpoint Hit: DeleteUserRatings")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	Logic.DeleteUserRatings(vars["userId"])
}

func DeleteMovieRatings(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Endpoint Hit: DeleteMovieRatings")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	Logic.DeleteMovieRatings(vars["movieId"])
}

func GetRating(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Endpoint Hit: GetRating")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	rating:= Logic.ReturnRating(vars["movieId"])
	json.NewEncoder(w).Encode(rating)
}

func GetReview(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Endpoint Hit: GetReview")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	queryValues := r.URL.Query()
	pageSize := queryValues.Get("pageSize")
	pageNo := queryValues.Get("pageNo")
	vars := mux.Vars(r)
	review:= Logic.ReturnReview(vars["movieId"], pageNo, pageSize)
	json.NewEncoder(w).Encode(review)
}

func PutReview(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: PutReview")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	data, err := ioutil.ReadAll(r.Body)
	var review Review.ReviewDetails
	//fmt.Printf()
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &review)
	vars := mux.Vars(r)
	review.MovieId, err = strconv.Atoi(vars["movieId"])
	if err != nil {
		panic(err)
	}
	Logic.UpdateReview(review)
}

func DeleteReview(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: DeleteReview")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	data, err := ioutil.ReadAll(r.Body)
	var review Review.ReviewDetails
	//fmt.Printf()
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &review)
	vars := mux.Vars(r)
	Logic.DeleteReview(vars["movieId"], review.UserId)
}

func GetShows(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Endpoint Hit: GetShows")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	queryValues := r.URL.Query()
	pageSize := queryValues.Get("pageSize")
	pageNo := queryValues.Get("pageNo")
	vars := mux.Vars(r)
	review:= Logic.ReturnShows(vars["movieId"],vars["city"], pageNo, pageSize)
	json.NewEncoder(w).Encode(review)
}

func PutShow(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: PutShow")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var show Shows.ShowDetails
	json.Unmarshal(data, &show)
	vars := mux.Vars(r)
	show.MovieId, err = strconv.Atoi(vars["movieId"])
	if err != nil {
		panic(err)
	}
	show.City = vars["city"]
	Logic.AddShow(show)
}

func DeleteShow(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: DeleteShow")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var show Shows.ShowDetails
	json.Unmarshal(data,&show)
	vars := mux.Vars(r)
	show.MovieId,err = strconv.Atoi(vars["movieId"])
	if err!=nil{
		panic(err.Error())
	}
	show.CineplexId = vars["cineplexId"]
	Logic.DeleteShow(show)
}

func DeleteAllShows(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: DeleteAllShows")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	movieId:=vars["movieId"]
	Logic.DeleteAllShows(movieId)
}