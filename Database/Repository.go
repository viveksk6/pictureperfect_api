package Database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


type Repository interface {
	FetchAllMovies(pageNo string, pageSize string) []MovieDetails
	FetchMovieById(movieId string, movieDetails []MovieDetails)  []MovieDetails
	ConnectDb() *sql.DB
}
