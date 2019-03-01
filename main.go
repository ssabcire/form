package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/createuser", createUser)
	mux.HandleFunc("/deluser", delUser)

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("form.html"))
	err := t.ExecuteTemplate(w, "form.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	user := User{}
	//リクエストを解析し、フォームを取得。
	user.Name = r.PostFormValue("name")

	err := user.Create()
	if err != nil {
		log.Fatal("create err")
	}
	http.Redirect(w, r, "/", 302)
}

func delUser(w http.ResponseWriter, r *http.Request) {
	//削除する関数
	id, err := strconv.Atoi(r.PostFormValue("id"))
	if err != nil {
		log.Fatal("delのAtoi変換エラー")
		//エラー1. 未入力
		//エラー2. 削除したいID以外の文字677
		//エラー3. IDの範囲外
	}
	err = Delete(id)
	if err != nil {
		log.Fatal("IDデリートエラー")
	}
	http.Redirect(w, r, "/", 302)
}
