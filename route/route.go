package route

import "github.com/dangviethung096/core"

func Init() {
	// Homepage
	core.RegisterPage("/", core.Page{
		PageFiles: []string{
			"./html/index.html",
			"./html/welcome_image.html",
		},
		Data: HomePageData{
			Title: "Wedding page",
			Groom: "Hiệp",
			Bride: "Hạnh",
			WelcomeImage: WelcomeImageData{
				Groom:       "Hiệp",
				Bride:       "Hạnh",
				Information: "Chúng tôi sẽ kết hôn vào ngày 14/02/2024",
				YoutubeLink: "https://www.youtube.com/watch?v=AQqshO7MiMI",
			},
			NavHome:    "Home",
			NavAbout:   "About",
			NavGallery: "Gallery",
			NavStory:   "Story",
			NavFamily:  "Family",
			NavEvent:   "Event",
			NavRSVP:    "RSVP",
			NavContact: "Contact",
		},
	})
}
