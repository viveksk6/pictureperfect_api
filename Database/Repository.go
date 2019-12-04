package Database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


type Repository interface {
	ConnectDb() *sql.DB
	ValidateCred(userId string, password string) (int,string)
	FetchAllMovies(pageNo string, pageSize string) []MovieDetails
	FetchMovieById(movieId string, movieDetails []MovieDetails)  []MovieDetails
	InsertMovie(movie MovieDetails)
	InsertRating(review ReviewDetails)
	FetchUserRating(movieId string, userId string)  ReviewDetails
	DeleteUserRating(movieId string, userId string)
	FetchRating(movieId string)  ReviewDetails
}
