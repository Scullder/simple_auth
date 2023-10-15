package handlers

import (
	"html/template"
	"net/http"

	"github.com/Kwynto/gosession"
)

func Root(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("web/static/index.html")
	tmpl.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", 303)
	}

	session := gosession.Start(&w, r)

	var users = map[string]string{
		"email@email.com": "12345",
	}

	email, password := r.FormValue("email"), r.FormValue("password")

	if users[email] == password {
		session.Set("is_login", true)
		http.Redirect(w, r, "/hello", 303)
	}

	http.Redirect(w, r, "/", 303)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/hello", 303)
	}

	session := gosession.Start(&w, r)
	session.Remove("is_login")

	http.Redirect(w, r, "/", 303)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("web/static/hello.html")
	tmpl.Execute(w, nil)
}
