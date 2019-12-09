package Catalogue

type MovieDetails struct {
	ID   int    `json:"ID"`
	Title string `json:"title,omitempty"`
	Summary string `json:"summary,omitempty"`
	Genre string `json:"genre,omitempty"`
	Img string `json:"img,omitempty"`
	Language string `json:"language,omitempty"`
	Certificate string `json:"certificate,omitempty"`
}
