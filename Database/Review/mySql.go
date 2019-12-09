package Review

import (
	"fmt"
	"strconv"
)

func (s *SQLRepo)InsertRating(review ReviewDetails){
	q:= " "
	results, err := s.db.Exec("INSERT into review values (?,?,?,?) on duplicate key update  rating = ?",review.MovieId,review.UserId, q,review.Review,review.Rating)
	if err!= nil {
		panic(err.Error())
	}
	fmt.Print(results.LastInsertId())

}

func (s *SQLRepo)FetchUserRating(movieId string, userId string)  ReviewDetails{
	//q:="Select * from review where"
	results,err:=s.db.Query("Select * from review where movieId = ? and userId = ?",movieId,userId)
	if err!= nil {
		panic(err.Error())
	}
	var review ReviewDetails
	for results.Next() {
		err = results.Scan(&review.MovieId, &review.UserId, &review.Review, &review.Rating)
		if err != nil {
			panic(err.Error())
		}
	}
	return review
}

func (s *SQLRepo)DeleteUserRating(movieId string, userId string)  {

	_,err:=s.db.Exec("UPDATE review set rating = -1 where movieId = ? and userId = ?",movieId,userId)
	if err!= nil {
		panic(err.Error())
	}

}

func (s *SQLRepo)FetchRating(movieId string)  ReviewDetails{

	results,err:=s.db.Query("Select avg(rating) from review where movieId = ? and rating >= 0",movieId)
	if err!= nil {
		panic(err.Error())
	}

	var review ReviewDetails

	for results.Next() {
		err = results.Scan(&review.Rating)
		if err != nil {
			panic(err.Error())
		}
	}
	review.MovieId,err = strconv.Atoi(movieId)
	if err!=nil {
		panic(err.Error())
	}
	return review
}

func (s *SQLRepo)DeleteUserRatings(userId string) {

	_,err:=s.db.Exec("UPDATE review set rating = -1 where userId = ?",userId)
	if err!= nil {
		panic(err.Error())
	}
}

func (s *SQLRepo)DeleteMovieRatings(movieId string) {

	_,err:=s.db.Exec("UPDATE review set rating = -1 where movieId = ?",movieId)
	if err!= nil {
		panic(err.Error())
	}
}

func (s *SQLRepo)FetchReview(movieId string, pageNo string, pageSize string)  []ReviewDetails{

	q:=""
	results,err:=s.db.Query("SELECT review.movieId, review.review,users.userId, users.name FROM `review` JOIN `users` ON users.userId = review.userId where review.movieId = ? and review != ? limit ?,?",movieId,q, pageNo, pageSize)
	if err!= nil {
		panic(err.Error())
	}

	var AllReviews []ReviewDetails
	var review ReviewDetails

	for results.Next() {
		err = results.Scan(&review.MovieId, &review.Review, &review.UserId, &review.UserName)
		if err != nil {
			panic(err.Error())
		}
		AllReviews=append(AllReviews,review)
	}
	return AllReviews
}

func (s *SQLRepo)UpdateReview(review ReviewDetails){
	results, err := s.db.Exec("INSERT into review values (?,?,?,-1) on duplicate key update  review = ?",review.MovieId,review.UserId,review.Review,review.Review)
	if err!= nil {
		panic(err.Error())
	}
	fmt.Print(results.LastInsertId())
}

func (s *SQLRepo)DeleteReview(movieId string, userId int){
	_,err:=s.db.Exec("UPDATE review set review = null where movieId = ? and userId = ?",movieId, userId)
	if err!= nil {
		panic(err.Error())
	}
}
