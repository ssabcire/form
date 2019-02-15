package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
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

func main() {
	err := Create()
	if err != nil {
		log.Fatal("Create error!")
	}
	fmt.Println("Createできたよ!")

	user, err := GetPost(1)
	if err != nil {
		log.Fatal("GetPost error!")
	}
	fmt.Print("ID1の値を取得したよ:")
	fmt.Println(user)
}

func Create() (err error) {
	_, err = Db.Exec("INSERT INTO tm (name) VALUES ('nyaa')")
	fmt.Println("test")
	return
}

func GetPost(id int) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, name FROM tm WHERE id = ?", id).Scan(&user.Id, &user.Name)
	return
}