package Shows

import (
	"PicturePerfect2/Database"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type SQLRepo struct {
	db *sql.DB
}

var sqlr SQLRepo

func ConnectDb() {
	sqlr.db = Database.ReturnDB()
}

func GetMovieRepo() *SQLRepo {
	return  &SQLRepo{
		db: sqlr.db,
	}
}
