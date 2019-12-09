package Database

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	UserName string
	Password string
	Endpoint string
	Port string
	DbName string
}

type SQLRepo struct {
	db *sql.DB
}

var sqlr SQLRepo

func ConnectDb() {
	var err error
	file, _ := ioutil.ReadFile("configuration.json")
	var c Config
	_ = json.Unmarshal(file, &c)
	dataSource:= "" + c.UserName + ":" + c.Password + "@tcp(" + c.Endpoint + ":" + c.Port + ")/" + c.DbName
	sqlr.db, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err.Error())
	}
}

func ReturnDB() *sql.DB {
	return sqlr.db
}
