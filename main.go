package main

import (
	"flag"
	"fmt"
	"net/http"
	"read_book_webs/controller"
	"read_book_webs/route"
)

var port = flag.Int64("port", 8080, "listen port")

func main() {
	flag.Parse()
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./html/assets"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./html/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./html/js"))))
	controller.Init()

	http.HandleFunc("/login", route.RouteLoginPage)
	http.HandleFunc("/", route.RouteHomePage)

	fmt.Printf("Listen and serve %d\n", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		fmt.Printf("Listen port %d fail: %s", *port, err.Error())
	}
}
