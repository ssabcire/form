package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "tm:tm55@/tm")
	if err != nil {
		panic(err)
	}
}

func Create() (err error) {
	_, err = Db.Exec("INSERT INTO tm (name) VALUES ('yama')")
	return
}

func GetPost(id int) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, name FROM tm WHERE id = ?", id).Scan(&user.Id, &user.Name)
	return
}
