package Logic

import (
	"PicturePerfect2/Database/Catalogue"
	"PicturePerfect2/Database/IAM"
	"PicturePerfect2/Database/Review"
	"PicturePerfect2/Database/Shows"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	Role string `json:"role"`
	jwt.StandardClaims
}

var JwtKey = []byte("clumio-pictureperfect")

func GenerateJWT(userId string, password string) (string, time.Time,int) {
	db:= IAM.GetMovieRepo()
	rows, role:= db.ValidateCred(userId,password)
	if rows==0{
		return "",time.Time{}, http.StatusUnauthorized
	}
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: userId,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "",time.Time{},http.StatusInternalServerError
	}
	return tokenString, expirationTime, 0
}

func ResetPassword(userId int, password string, newPassword string) string {
	db:= IAM.GetMovieRepo()
	res:= db.ResetPassword(userId,password,newPassword)
	return  res
}

func GetAllMovies(pageNo string, pageSize string)  []Catalogue.MovieDetails{
	db:= Catalogue.GetMovieRepo()
	pageNoInt, err:= strconv.Atoi(pageNo)
	if err != nil {
		panic(err.Error())
	}
	pageSizeInt, err:= strconv.Atoi(pageSize)
	if err != nil {
		panic(err.Error())
	}
	pageNo = strconv.Itoa(pageSizeInt*(pageNoInt-1))
	movieDetails:=db.FetchAllMovies(pageNo,pageSize)
	return movieDetails
}

func GetMovieById(movieId string)  []Catalogue.MovieDetails{
	db:= Catalogue.GetMovieRepo()
	movieDetails:=db.FetchMovieById(movieId)
	return movieDetails
}

func AddMovie(movie Catalogue.MovieDetails){
	db:= Catalogue.GetMovieRepo()
	db.InsertMovie(movie)
}

func UpdateMovie(movie Catalogue.MovieDetails){
	db:= Catalogue.GetMovieRepo()
	db.UpdateMovie(movie)
}

func UpdateMovieAttribute(movieId int, movieDetail map[string]string){
	db:= Catalogue.GetMovieRepo()
	var attribute, value_attribute  string
	for key, value:= range movieDetail{
		attribute = key
		value_attribute = value
	}
	db.UpdateMovieAttribute(movieId, attribute, value_attribute)
}

func AddRating(review Review.ReviewDetails)  {
	db:= Review.GetMovieRepo()
	db.InsertRating(review)
}

func ReturnUserRating(movieId string, userId string) Review.ReviewDetails  {
	db:= Review.GetMovieRepo()
	rating:=db.FetchUserRating(movieId, userId)
	return  rating
}

func DeleteUserRating(movieId string, userId string)   {
	db:= Review.GetMovieRepo()
	db.DeleteUserRating(movieId, userId)

}

func DeleteUserRatings(userId string)   {
	db:= Review.GetMovieRepo()
	db.DeleteUserRatings(userId)

}

func DeleteMovieRatings(movieId string)   {
	db:= Review.GetMovieRepo()
	db.DeleteMovieRatings(movieId)

}

func ReturnRating(movieId string) Review.ReviewDetails  {
	db:= Review.GetMovieRepo()
	rating:=db.FetchRating(movieId)
	return  rating
}

func ReturnReview(movieId string, pageNo string, pageSize string) []Review.ReviewDetails  {
	db:= Review.GetMovieRepo()
	pageNoInt, err:= strconv.Atoi(pageNo)
	if err != nil {
		panic(err.Error())
	}
	pageSizeInt, err:= strconv.Atoi(pageSize)
	if err != nil {
		panic(err.Error())
	}
	pageNo = strconv.Itoa(pageSizeInt*(pageNoInt-1))
	review:=db.FetchReview(movieId, pageNo, pageSize)
	return  review
}

func UpdateReview(review Review.ReviewDetails){
	db:= Review.GetMovieRepo()
	db.UpdateReview(review)
}

func DeleteReview(movieId string, userId int){
	db:= Review.GetMovieRepo()
	db.DeleteReview(movieId,userId)
}

func ReturnShows(movieId string,city string, pageNo string, pageSize string) []Shows.ShowDetails  {
	db:= Shows.GetMovieRepo()
	pageNoInt, err:= strconv.Atoi(pageNo)
	if err != nil {
		panic(err.Error())
	}
	pageSizeInt, err:= strconv.Atoi(pageSize)
	if err != nil {
		panic(err.Error())
	}
	pageNo = strconv.Itoa(pageSizeInt*(pageNoInt-1))
	shows:=db.FetchShows(movieId,city, pageNo, pageSize)
	return  shows
}

func AddShow(show Shows.ShowDetails)  {
	db:= Shows.GetMovieRepo()
	db.InsertShow(show)
}

func DeleteShow(show Shows.ShowDetails){
	db:= Shows.GetMovieRepo()
	db.DeleteShow(show)
}

func DeleteAllShows(movieId string){
	db:= Shows.GetMovieRepo()
	db.DeleteAllShows(movieId)
}