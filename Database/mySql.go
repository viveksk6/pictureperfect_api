package Database

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
)

type SQLRepo struct {
	db *sql.DB
}

type Config struct {
	UserName string
	Password string
	Endpoint string
	Port string
	DbName string
}

var sqlr SQLRepo

func ConnectDb() {
	var err error
	file, _ := ioutil.ReadFile("configuration.json")
	var c Config
	_ = json.Unmarshal(file, &c)
	dataSource:= "" + c.UserName + ":" + c.Password + "@tcp(" + c.Endpoint + ":" + c.Port + ")/" + c.DbName
	sqlr.db, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err.Error())
	}
}

func GetMovieRepo() *SQLRepo {
	return  &SQLRepo{
		db: sqlr.db,
	}
}

func (s *SQLRepo)FetchAllMovies(pageNo string, pageSize string) []MovieDetails {
	var movieDetails []MovieDetails
	//sqlr.db := ConnectDb()
	//defer sqlr.db.Close()
	q := "SELECT movieId, title, summary, genre, img, language, certificate from movies limit "
	qstring := q + pageNo + "," + pageSize
	results, err := s.db.Query(qstring)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	for results.Next() {
		var md MovieDetails

		err = results.Scan(&md.MovieId, &md.Title,  &md.Summary, &md.Genre, &md.Img, &md.Language, &md.Certificate)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		movieDetails = append(movieDetails, md)
	}
	return movieDetails
}

func (s *SQLRepo)FetchMovieById(movieId string)  []MovieDetails{
	var movieDetails []MovieDetails
	//db:= ConnectDb()
	//defer sqlr.db.Close()
	q := "SELECT movieId, title, summary, genre, img, language, certificate from movies where movieId = "
	qstring := q + "'" + movieId + "'"
	results, err := s.db.Query(qstring)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	for results.Next() {
		var md MovieDetails

		err = results.Scan(&md.MovieId, &md.Title,  &md.Summary, &md.Genre, &md.Img, &md.Language, &md.Certificate)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		movieDetails = append(movieDetails, md)
	}
	return movieDetails
}
