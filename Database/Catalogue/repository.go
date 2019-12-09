package Catalogue

type Repository interface {
	FetchAllMovies(pageNo string, pageSize string) []MovieDetails
	FetchMovieById(movieId string) []MovieDetails
	InsertMovie(movie MovieDetails)
	UpdateMovie(movie MovieDetails)
	UpdateMovieAttribute(movieId int, column string, value string)
}
