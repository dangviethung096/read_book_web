package main

import (
	"fmt"
	"net/http"
	"read_book_webs/controller"
	"read_book_webs/route"
)

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./html/assets"))))

	http.HandleFunc("/login", route.RouteLoginPage)
	controller.Init()

	fmt.Println("Listen and serve 80")
	http.ListenAndServe(":80", nil)
}
