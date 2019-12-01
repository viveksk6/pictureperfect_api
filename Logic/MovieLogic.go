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