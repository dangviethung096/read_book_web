package route

import (
	"net/http"
	"text/template"
)

type HomePageData struct {
	Title string
}

func RouteHomePage(w http.ResponseWriter, r *http.Request) {
	data := HomePageData{
		Title: "Home",
	}

	tmpl, err := template.ParseFiles(("html/index.html"))
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
