package main

import (
	"github.com/dangviethung096/core"
)

func main() {
	core.Init("core.config.yaml")

	// http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./html/assets"))))
	// http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./html/css"))))
	// http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./html/js"))))
	// controller.Init()

	// http.HandleFunc("/login", route.RouteLoginPage)
	// http.HandleFunc("/", route.RouteHomePage)

	core.Start()

}
