package Router

import (
	"PicturePerfect2/API"
	"github.com/gorilla/mux"
)

func (handler *Handlers)HandleIAMRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/login", handler.Logger(API.Login))
}

func (handler *Handlers)HandleCatalogueRequests(catalogue *mux.Router) {
	catalogue.HandleFunc("", handler.CheckURLParams(API.ReturnAllMovies)).Methods("GET")
	catalogue.HandleFunc("", API.EnableCors).Methods("OPTIONS")
	catalogue.HandleFunc("", handler.AuthenticateAdmin(API.PostMovie)).Methods("POST")
	catalogue.HandleFunc("/{movieId}", handler.AuthenticateUser(API.ReturnSingleMovie)).Methods("GET")
	catalogue.HandleFunc("/{movieId}", API.EnableCors).Methods("OPTIONS")
}

func (handler *Handlers)HandleRatingRequests(rating *mux.Router) {
	rating.HandleFunc("/movie/{movieId}/user/{userId}", handler.AuthenticateUser(API.PostRating)).Methods("POST")
	rating.HandleFunc("/movie/{movieId}/user/{userId}", handler.AuthenticateUser(API.GetUserRating)).Methods("GET")
	rating.HandleFunc("/movie/{movieId}/user/{userId}", handler.AuthenticateUser(API.DeleteUserRating)).Methods("DELETE")
	rating.HandleFunc("/movie/{movieId}", handler.Logger(API.GetRating)).Methods("GET")
}

func (handler *Handlers)HandleRequests(myRouter *mux.Router) {

	myRouter.HandleFunc("/", handler.Logger(API.HomePage))

	handler.HandleIAMRequests(myRouter)

	catalogue:= myRouter.PathPrefix("/movies/catalogue").Subrouter()
	handler.HandleCatalogueRequests(catalogue)

	rating:= myRouter.PathPrefix("/movies/rating").Subrouter()
	handler.HandleRatingRequests(rating)


}
