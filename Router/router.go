package Router

import (
	"PicturePerfect2/API"
	"github.com/gorilla/mux"
)

func (handler *Handlers)HandleIAMRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/login", handler.Logger(API.Login)).Methods("POST")
	myRouter.HandleFunc("/reset", handler.AuthenticateUser(API.ResetPassword)).Methods("POST")
}

func (handler *Handlers)HandleCatalogueRequests(catalogue *mux.Router) {
	catalogue.HandleFunc("", handler.CheckURLParams(API.ReturnAllMovies)).Methods("GET")
	catalogue.HandleFunc("", API.EnableCors).Methods("OPTIONS")
	catalogue.HandleFunc("", handler.AuthenticateAdmin(API.PostMovie)).Methods("POST")
	catalogue.HandleFunc("/{movieId}", handler.AuthenticateUser(API.ReturnSingleMovie)).Methods("GET")
	catalogue.HandleFunc("/{movieId}", API.EnableCors).Methods("OPTIONS")
	catalogue.HandleFunc("/{movieId}", handler.AuthenticateAdmin(API.PutMovie)).Methods("PUT")
	catalogue.HandleFunc("/{movieId}", handler.AuthenticateAdmin(API.PatchMovie)).Methods("PATCH")
}

func (handler *Handlers)HandleRatingRequests(rating *mux.Router) {
	rating.HandleFunc("/movie/{movieId}/user/{userId}", handler.AuthenticateUser(API.PostRating)).Methods("PUT")
	rating.HandleFunc("/movie/{movieId}/user/{userId}", handler.AuthenticateUser(API.GetUserRating)).Methods("GET")
	rating.HandleFunc("/movie/{movieId}/user/{userId}", handler.AuthenticateUser(API.DeleteUserRating)).Methods("DELETE")
	rating.HandleFunc("/movie/{movieId}", handler.Logger(API.GetRating)).Methods("GET")
	rating.HandleFunc("/movie/{movieId}", handler.Logger(API.DeleteMovieRatings)).Methods("DELETE")
	rating.HandleFunc("/user/{userId}", handler.AuthenticateAdmin(API.DeleteUserRatings)).Methods("DELETE")
}

func (handler *Handlers)HandleReviewRequests(review *mux.Router) {
	review.HandleFunc("/{movieId}", handler.AuthenticateUser(handler.CheckURLParams(API.GetReview))).Methods("GET")
	review.HandleFunc("/{movieId}", handler.AuthenticateAdmin(API.DeleteReview)).Methods("DELETE")
	review.HandleFunc("/{movieId}", handler.AuthenticateUser(API.PutReview)).Methods("PUT")
}

func (handler *Handlers)HandleShowRequests(shows *mux.Router) {
	shows.HandleFunc("/{city}/movie/{movieId}",handler.AuthenticateUser(API.GetShows)).Methods("GET")
	shows.HandleFunc("/{city}/movie/{movieId}",handler.AuthenticateAdmin(API.PutShow)).Methods("PUT")
	shows.HandleFunc("/cineplex/{cineplexId}/movie/{movieId}",handler.AuthenticateAdmin(API.DeleteShow)).Methods("DELETE")
	shows.HandleFunc("/movie/{movieId}",handler.AuthenticateAdmin(API.DeleteAllShows)).Methods("DELETE")
}

func (handler *Handlers)HandleRequests(myRouter *mux.Router) {

	myRouter.HandleFunc("/", handler.Logger(API.HomePage))

	handler.HandleIAMRequests(myRouter)

	catalogue:= myRouter.PathPrefix("/movies/catalogue").Subrouter()
	handler.HandleCatalogueRequests(catalogue)

	rating:= myRouter.PathPrefix("/movies/rating").Subrouter()
	handler.HandleRatingRequests(rating)

	review:= myRouter.PathPrefix("/movies/review").Subrouter()
	handler.HandleReviewRequests(review)

	shows:= myRouter.PathPrefix("/movies/shows").Subrouter()
	handler.HandleShowRequests(shows)
}
