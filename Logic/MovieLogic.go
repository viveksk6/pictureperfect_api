package Logic

import (
	"PicturePerfect2/Database"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	Role string `json:"role"`
	jwt.StandardClaims
}

var JwtKey = []byte("clumio-pictureperfect")

func GenerateJWT(userId string, password string) (string, time.Time,int) {
	db:= Database.GetMovieRepo()
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

func GetAllMovies(pageNo string, pageSize string)  []Database.MovieDetails{
	db:= Database.GetMovieRepo()
	movieDetails:=db.FetchAllMovies(pageNo,pageSize)
	return movieDetails
}

func GetMovieById(movieId string)  []Database.MovieDetails{
	db:= Database.GetMovieRepo()
	movieDetails:=db.FetchMovieById(movieId)
	return movieDetails
}

func AddMovie(movie Database.MovieDetails){
	db:=Database.GetMovieRepo()
	db.InsertMovie(movie)
}

func AddRating(review Database.ReviewDetails)  {
	db:=Database.GetMovieRepo()
	db.InsertRating(review)
}

func ReturnUserRating(movieId string, userId string) Database.ReviewDetails  {
	db:=Database.GetMovieRepo()
	rating:=db.FetchUserRating(movieId, userId)
	return  rating
}

func DeleteUserRating(movieId string, userId string)   {
	db:=Database.GetMovieRepo()
	db.DeleteUserRating(movieId, userId)

}

func ReturnRating(movieId string) Database.ReviewDetails  {
	db:=Database.GetMovieRepo()
	rating:=db.FetchRating(movieId)
	return  rating
}