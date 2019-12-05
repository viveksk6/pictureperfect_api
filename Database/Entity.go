package Database

type MovieDetails struct {
	ID   int    `json:"ID"`
	Title string `json:"title"`
	Summary string `json:"summary"`
	Genre string `json:"genre"`
	Img string `json:"img"`
	Language string `json:"language"`
	Certificate string `json:"certificate"`
}

type ReviewDetails struct {
	MovieId int `json:"movieId"`
	UserId int `json:"userId"`
	Rating float64 `json:"rating"`
	Review string `json:"review"`
}

//func (mv *MovieDetails) QueryDb(pageNo string, pageSize string, movieDetails []MovieDetails) []MovieDetails
