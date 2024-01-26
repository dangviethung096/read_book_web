package route

import "github.com/dangviethung096/core"

func Init() {
	// Homepage
	core.RegisterPage("/", core.Page{
		PageFiles: []string{"html/index.html"},
		Data: HomePageData{
			Title: "Home",
		},
	})
}
