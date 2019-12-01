package Database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"strconv"
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

func (s *SQLRepo)InsertRating(review ReviewDetails){
	mId:= strconv.Itoa(review.MovieId)
	uId:= strconv.Itoa(review.UserId)
	rating:= strconv.FormatFloat(review.Rating, 'f', 6, 32)
	q:= "Insert into review values(" + mId + "," + uId + "," + "'" + review.Review + "'" + "," + rating+")"
	results, err := s.db.Exec(q)
	if err!= nil {
		panic(err.Error())
	}
	fmt.Print(results.LastInsertId())

}

func (s *SQLRepo)FetchUserRating(movieId string, userId string)  ReviewDetails{
	//q:="Select * from review where"
	results,err:=s.db.Query("Select * from review where movieId = ? and userId = ?",movieId,userId)
	if err!= nil {
		panic(err.Error())
	}
	var review ReviewDetails
	for results.Next() {
		err = results.Scan(&review.MovieId, &review.UserId, &review.Review, &review.Rating)
		if err != nil {
			panic(err.Error())
		}
	}
	return review
}

func (s *SQLRepo)DeleteUserRating(movieId string, userId string)  {

	_,err:=s.db.Exec("DELETE from review where movieId = ? and userId = ?",movieId,userId)
	if err!= nil {
		panic(err.Error())
	}

}

func (s *SQLRepo)FetchRating(movieId string)  ReviewDetails{

	results,err:=s.db.Query("Select avg(rating) from review where movieId = ? ",movieId)
	if err!= nil {
		panic(err.Error())
	}

	var review ReviewDetails

	for results.Next() {
		err = results.Scan(&review.Rating)
		if err != nil {
			panic(err.Error())
		}
	}
	review.MovieId,err = strconv.Atoi(movieId)
	if err!=nil {
		panic(err.Error())
	}
	return review
}
