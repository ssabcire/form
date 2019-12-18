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

//------以下CRUDを行う関数------------------
func (user *User) Create() (err error) {
	//SQLをただ実行するだけ。行を返さない
	_, err = Db.Exec("INSERT INTO tm (name) VALUES (?)", user.Name)
	if err != nil {
		return
	}
	//SQLを実行し、1行だけ値を返す。(QueryRowの説明)
	err = Db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&user.Id)
	return
}

func Delete(id int) (err error) {
	_, err = Db.Exec("DELETE FROM tm WHERE id = ?", id)
	return
}

//すべてのユーザーを得る
func Users() (users []User, err error) {
	rows, err := Db.Query("SELECT * FROM tm")
	if err != nil {
		return
	}
	for rows.Next() {
		conv := User{}
		if err = rows.Scan(&conv.Id, &conv.Name); err != nil {
			return
		}
		users = append(users, conv)
	}
	rows.Close()
	return
}

func (user *User) Update() (err error) {
	Db.Exec("UPDATE tm SET name = ? WHERE id = ?", user.Name, user.Id)
	return
}

//----------- 以下はテスト用の関数 ---
func DeleteUserAll() (err error) {
	_, err = Db.Exec("DELETE FROM tm")
	if err != nil {
		return
	}
	_, err = Db.Exec("UPDATE tm SET id = (@i := @i +1)")
	if err != nil {
		return
	}
	_, err = Db.Exec("ALTER TABLE tm AUTO_INCREMENT = 1")
	if err != nil {
		return
	}
	return
}

func (u *User) UpdateCheck() (err error) {
	// idを渡して、名前を取得する
	err = Db.QueryRow("SELECT name from tm WHERE id = ?", u.Id).Scan(&u.Name)
	if err != nil {
		return
	}
	return
}
