package main

import (
	"read_book_webs/controller"
	"read_book_webs/route"

	"github.com/dangviethung096/core"
)

func main() {
	core.Init("core.config.yaml")
	// Static file
	core.RegisterFolder("/css/", "/css/", "./html/css")
	core.RegisterFolder("/img/", "/img/", "./html/img")
	core.RegisterFolder("/js/", "/js/", "./html/js")
	core.RegisterFolder("/lib/", "/lib/", "./html/lib")
	core.RegisterFolder("/scss/", "/scss/", "./html/scss")
	core.RegisterFolder("/fonts/", "/fonts/", "./html/fonts")
	// Controller
	controller.Init()
	// Route to page
	route.Init()

	core.Start()

}
