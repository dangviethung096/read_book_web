package route

import "github.com/dangviethung096/core"

func Init() {
	var groomName = "Hiệp"
	var brideName = "Hạnh"
	// Homepage
	core.RegisterPage("/", core.Page{
		PageFiles: []string{
			"./html/index.html",
			"./html/carousel.html",
			"./html/nav_bar.html",
			"./html/about.html",
			"./html/story.html",
			"./html/gallery.html",
			"./html/event.html",
		},
		Data: HomePageData{
			Title: "Wedding page",
			Carousel: Carousel{
				GroomName:   groomName,
				BrideName:   brideName,
				Information: "Chúng tôi sẽ kết hôn",
				YoutubeLink: "https://www.youtube.com/watch?v=AQqshO7MiMI",
			},
			NavBar: NavBar{
				GroomName:  groomName,
				BrideName:  brideName,
				NavHome:    "Home",
				NavAbout:   "About",
				NavGallery: "Gallery",
				NavStory:   "Story",
				NavFamily:  "Family",
				NavEvent:   "Event",
				NavRSVP:    "RSVP",
				NavContact: "Contact",
			},
			About:            About{},
			Story:            Story{},
			Gallery:          Gallery{},
			Event:            Event{},
			FriendsAndFamily: FriendsAndFamily{},
			RSVP:             RSVP{},
			Footer:           Footer{},
		},
	})
}
