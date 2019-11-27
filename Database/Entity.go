package Database

type MovieDetails struct {
	MovieId   int    `json:"movieId"`
	Title string `json:"title"`
	Summary string `json:"summary"`
	Genre string `json:"genre"`
	Img string `json:"img"`
	Language string `json:"language"`
	Certificate string `json:"certificate"`
}

//func (mv *MovieDetails) QueryDb(pageNo string, pageSize string, movieDetails []MovieDetails) []MovieDetails
