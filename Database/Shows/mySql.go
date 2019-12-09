package Shows

import "fmt"

func (s *SQLRepo)FetchShows(movieId string, city string, pageNo string, pageSize string)  []ShowDetails{

	results,err:=s.db.Query("SELECT screening.movieId, screening.cineplexId, cinemas.cineplexName, screening.startTime FROM `screening` JOIN `cinemas` ON screening.cineplexId = cinemas.cineplexId where screening.movieId = ? and cinemas.city = ? limit ?,?",movieId,city, pageNo, pageSize)
	if err!= nil {
		panic(err.Error())
	}

	var AllShows []ShowDetails
	var show ShowDetails

	for results.Next() {
		err = results.Scan(&show.MovieId, &show.CineplexId, &show.CineplexName, &show.StartTime)
		if err != nil {
			panic(err.Error())
		}
		AllShows=append(AllShows,show)
	}
	return AllShows
}


func (s *SQLRepo)InsertShow(show ShowDetails){

	results, err := s.db.Query("Select cineplexId from cinemas where cineplexName = ? and city = ? ",show.CineplexName, show.City)
	if err!= nil {
		panic(err.Error())
	}
	var id int
	for results.Next(){
		results.Scan(&id)
	}
	updateResult, err := s.db.Exec("INSERT into screening values (?,?,?) ",show.MovieId,id,show.StartTime)
	if err!= nil {
		panic(err.Error())
	}
	fmt.Print(updateResult.LastInsertId())
}

func (s *SQLRepo)DeleteShow(show ShowDetails) {

	_,err:=s.db.Exec("DELETE from screening where movieId = ? and cineplexId = ? and startTime = ?",show.MovieId,show.CineplexId,show.StartTime)
	if err!= nil {
		panic(err.Error())
	}
}

func (s *SQLRepo)DeleteAllShows(movieId string) {

	_,err:=s.db.Exec("DELETE from screening where movieId = ? ",movieId)
	if err!= nil {
		panic(err.Error())
	}
}
