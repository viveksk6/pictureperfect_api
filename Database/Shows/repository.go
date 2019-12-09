package Shows

type Repository interface {
	FetchShows(movieId string, city string, pageNo string, pageSize string)  []ShowDetails
	InsertShow(show ShowDetails)
	DeleteShow(show ShowDetails)
	DeleteAllShows(movieId string)
}
