package main

import (
	"read_book_webs/controller"
	"read_book_webs/route"

	"github.com/dangviethung096/core"
)

func main() {
	core.Init("core.config.yaml")
	// Static file
	core.RegisterFolder("/assets/", "/assets/", "./html/assets")
	core.RegisterFolder("/css/", "/css/", "./html/css")
	core.RegisterFolder("/js/", "/js/", "./html/js")
	// Controller
	controller.Init()
	// Route to page
	route.Init()

	core.Start()

}
