package Router

import (
	"PicturePerfect2/API"
	"github.com/gorilla/mux"
)

func (handler *Handlers)HandleRequests(myRouter *mux.Router) {

	catalogue:= myRouter.PathPrefix("/movies/catalogue").Subrouter()
	myRouter.HandleFunc("/", handler.Logger(API.HomePage))
	catalogue.HandleFunc("", handler.CheckURLParams(API.ReturnAllMovies)).Methods("GET")
	catalogue.HandleFunc("", API.EnableCors).Methods("OPTIONS")
	catalogue.HandleFunc("/{movieId}", handler.Logger(API.ReturnSingleMovie)).Methods("GET")
	catalogue.HandleFunc("/{movieId}", API.EnableCors).Methods("OPTIONS")
	rating:= myRouter.PathPrefix("/movies/rating").Subrouter()
	rating.HandleFunc("/movie/{movieId}/user/{userId}", handler.Logger(API.PostRating)).Methods("POST")
	rating.HandleFunc("/movie/{movieId}/user/{userId}", handler.Logger(API.GetUserRating)).Methods("GET")
	rating.HandleFunc("/movie/{movieId}/user/{userId}", handler.Logger(API.DeleteUserRating)).Methods("DELETE")
	rating.HandleFunc("/movie/{movieId}", handler.Logger(API.GetRating)).Methods("GET")

}
