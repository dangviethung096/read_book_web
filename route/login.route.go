package route

import (
	"net/http"
	"text/template"
)

type LoginPageData struct {
	Title string
}

func RouteLoginPage(w http.ResponseWriter, r *http.Request) {
	data := LoginPageData{
		Title: "Login",
	}

	tmpl, err := template.ParseFiles(("html/login.html"))
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
