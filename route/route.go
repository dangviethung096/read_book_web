package route

import "github.com/dangviethung096/core"

func Init() {
	var groomName = "Hiệp"
	var brideName = "Hạnh"
	var homeTitle = "Home"
	var aboutTitle = "About"
	var galleryTitle = "Gallery"
	var storyTitle = "Story"
	var familyTitle = "Family"
	var eventTitle = "Event"
	var rSVPTitle = "RSVP"
	var contactTitle = "Contact"
	var groomTitle = "Chú rể"
	var brideTitle = "Cô dâu"
	var storyLineTitle = "Hanh & Hiep's Story"
	var storyContents = []StoryContent{}
	var imageLinks = []string{
		"img/carousel-1.jpg",
	}

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
			"./html/friends_and_family.html",
			"./html/rsvp.html",
			"./html/footer.html",
		},
		Data: HomePageData{
			Title: "Wedding page",
			Carousel: Carousel{
				GroomName:   groomName,
				BrideName:   brideName,
				Information: "Chúng tôi sẽ kết hôn",
				YoutubeLink: "https://www.youtube.com/watch?v=AQqshO7MiMI",
				ImageLinks:  imageLinks,
			},
			NavBar: NavBar{
				GroomName:  groomName,
				BrideName:  brideName,
				NavHome:    homeTitle,
				NavAbout:   aboutTitle,
				NavGallery: galleryTitle,
				NavStory:   storyTitle,
				NavFamily:  familyTitle,
				NavEvent:   eventTitle,
				NavRSVP:    rSVPTitle,
				NavContact: contactTitle,
			},
			VideoModal: VideoModal{},
			About: About{
				AboutTitle: aboutTitle,
				GroomTitle: groomTitle,
				BrideTitle: brideTitle,
				GroomAbout: "",
				BrideAbout: "",
				GroomName:  groomName,
				BrideName:  brideName,
			},
			Story: Story{
				StoryTitle:     storyTitle,
				StoryLineTitle: storyLineTitle,
				StoryContents:  storyContents,
			},
			Gallery:          Gallery{},
			Event:            Event{},
			FriendsAndFamily: FriendsAndFamily{},
			RSVP:             RSVP{},
			Footer:           Footer{},
		},
	})
}
