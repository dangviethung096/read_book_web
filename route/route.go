package route

import "github.com/dangviethung096/core"

func Init() {
	// Homepage
	core.RegisterPage("/", core.Page{
		PageFiles: []string{"./html/index.html"},
		Data: HomePageData{
			Title:            "Wedding page",
			Groom:            "Hiệp",
			Bride:            "Hạnh",
			FirstInformation: "Chúng tôi sẽ kết hôn vào ngày 14/02/2023",
		},
	})
}
