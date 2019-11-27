package Logic

import "PicturePerfect2/Database"

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
