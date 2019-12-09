package Review

type ReviewDetails struct {
	MovieId int `json:"movieId"`
	MovieTitle string `json:"movie_title,omitempty"`
	UserId int `json:"userId,omitempty"`
	UserName string `json:"user_name,omitempty"`
	Rating float64 `json:"rating,omitempty"`
	Review string `json:"review,omitempty"`
}
