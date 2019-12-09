package Catalogue



func (s *SQLRepo)FetchAllMovies(pageNo string, pageSize string) []MovieDetails {
	var movieDetails []MovieDetails
	//defer sqlr.db.Close()
	results, err := s.db.Query("SELECT ID, title, img from movies limit ?,?",pageNo,pageSize)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	for results.Next() {
		var md MovieDetails

		err = results.Scan(&md.ID, &md.Title, &md.Img)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		movieDetails = append(movieDetails, md)
	}
	return movieDetails
}

func (s *SQLRepo)FetchMovieById(movieId string)  []MovieDetails{
	var movieDetails []MovieDetails
	results, err := s.db.Query("SELECT ID, title, summary, genre, img, language, certificate from movies where ID= ?",movieId)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	for results.Next() {
		var md MovieDetails

		err = results.Scan(&md.ID, &md.Title,  &md.Summary, &md.Genre, &md.Img, &md.Language, &md.Certificate)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		movieDetails = append(movieDetails, md)
	}
	return movieDetails
}

func (s *SQLRepo)InsertMovie(movie MovieDetails)  {
	_, err := s.db.Exec("INSERT into movies (title, summary, genre, img, language, certificate) values (?, ?, ?, ?, ?, ?)", movie.Title, movie.Summary, movie.Genre, movie.Img, movie.Language, movie.Certificate)
	if err!= nil {
		panic(err.Error())
	}
}

func (s *SQLRepo)UpdateMovie(movie MovieDetails)  {
	results, err := s.db.Exec("UPDATE movies set title = ?, summary = ?, genre = ?, img = ?, language = ?, certificate = ? where ID = ?", movie.Title, movie.Summary, movie.Genre, movie.Img, movie.Language, movie.Certificate, movie.ID)
	if err!= nil {
		panic(err.Error())
	}
	i := 0
	var i64 int64
	i64 = int64(i)
	rows, err:= results.RowsAffected()
	if err!= nil {
		panic(err.Error())
	}
	if rows == i64{
		return
	}
}

func (s *SQLRepo)UpdateMovieAttribute(movieId int, column string, value string)  {
	results, err := s.db.Exec("UPDATE movies set "+ column + " = ? where ID = ?",  value,movieId)
	if err!= nil {
		panic(err.Error())
	}
	i := 0
	var i64 int64
	i64 = int64(i)
	rows, err:= results.RowsAffected()
	if err!= nil {
		panic(err.Error())
	}
	if rows == i64{
		return
	}
}
