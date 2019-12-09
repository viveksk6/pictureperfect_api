package Shows

type ShowDetails struct {
	MovieId int `json:"movieId"`
	CineplexId string `json:"cineplexId"`
	CineplexName string `json:"cineplexName"`
	City string `json:"city,omitempty"`
	StartTime string `json:"startTime"`
}
