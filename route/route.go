package route

import "github.com/dangviethung096/core"

func Init() {
	// Login page
	core.RegisterPage("/login", core.Page{
		PageFiles: []string{"html/login.html"},
		Data: LoginPageData{
			Title: "Login",
		},
	})
	// Homepage
	core.RegisterPage("/", core.Page{
		PageFiles: []string{"html/index.html"},
		Data: HomePageData{
			Title: "Home",
		},
	})
}
