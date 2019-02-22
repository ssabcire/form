package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/createuser", createUser)

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
	/*	user = User{}
		len := r.ConentLength
		body := make([]byte, len)
		r.Body.Read(body)
		var user User
		json.Unmarshal(body, &user)
	*/
	err := Create()
	if err != nil {
		log.Fatal("create err")
	}
	http.Redirect(w, r, "/", 302)
}
