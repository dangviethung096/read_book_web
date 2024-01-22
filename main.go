package main

import (
	"fmt"
	"net/http"
	"read_book_webs/controller"
	"read_book_webs/route"
)

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./html/assets"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./html/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./html/js"))))
	controller.Init()

	http.HandleFunc("/login", route.RouteLoginPage)
	http.HandleFunc("/", route.RouteHomePage)

	fmt.Println("Listen and serve 80")
	http.ListenAndServe(":80", nil)
}
