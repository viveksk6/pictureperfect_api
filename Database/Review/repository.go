package Review

type Repository interface {
	InsertRating(review ReviewDetails)
	FetchUserRating(movieId string, userId string)  ReviewDetails
	DeleteUserRating(movieId string, userId string)
	FetchRating(movieId string)  ReviewDetails
	DeleteUserRatings(userId string)
	DeleteMovieRatings(movieId string)
	FetchReview(movieId string, pageNo string, pageSize string) []ReviewDetails
	UpdateReview(review ReviewDetails)
	DeleteReview(movieId string, userId int)
}
