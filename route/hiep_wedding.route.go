package route

import "github.com/dangviethung096/core"

func hiepWedding() {
	var groomName = "Hiệp"
	var brideName = "Hạnh"
	var groomFullName = "Nguyễn Thái Hiệp"
	var brideFullName = "Phạm Thị Hạnh"
	var homeTitle = "Trang chủ"
	var aboutTitle = "Về chúng tớ"
	var galleryTitle = "Ảnh"
	var storyTitle = "Ký ức"
	var familyTitle = "Family"
	var eventTitle = "Đám cưới"
	var rSVPTitle = "RSVP"
	var contactTitle = "Liên hệ"
	var groomTitle = "Chú rể"
	var brideTitle = "Cô dâu"
	var locationTitle = "Địa điểm"
	var bankTitle = "Lời nhắn"

	var groomAbout = ""
	var brideAbout = ""
	var imageLinks = []string{
		"img/home01.jpg",
		"img/home02.jpg",
		"img/home03.jpg",
		"img/home04.jpg",
	}

	var youtubeLink = "https://www.youtube.com/embed/w_17Bz-ngVM?si=vJenP1mx8mXrL5QF"

	homePageData := &HomePageData{
		Title: "Wedding page",
		NavBar: NavBar{
			GroomName:   groomName,
			BrideName:   brideName,
			NavHome:     homeTitle,
			NavAbout:    aboutTitle,
			NavGallery:  galleryTitle,
			NavStory:    storyTitle,
			NavFamily:   familyTitle,
			NavEvent:    eventTitle,
			NavRSVP:     rSVPTitle,
			NavContact:  contactTitle,
			NavLocation: locationTitle,
			NavBank:     bankTitle,
		},
		VideoModal: VideoModal{
			YoutubeLink: youtubeLink,
		},
		About: About{
			AboutTitle: aboutTitle,
			GroomTitle: groomTitle,
			BrideTitle: brideTitle,
			GroomAbout: groomAbout,
			BrideAbout: brideAbout,
			GroomName:  groomFullName,
			BrideName:  brideFullName,
		},
		Story: Story{
			StoryTitle: storyTitle,
		},
		Gallery:          Gallery{},
		Event:            Event{},
		FriendsAndFamily: FriendsAndFamily{},
		RSVP:             RSVP{},
		Footer:           Footer{},
	}

	// Update carousel data
	var carouselData = []ImageCarousel{}

	for _, link := range imageLinks {
		carouselData = append(carouselData, ImageCarousel{
			GroomName:   groomName,
			BrideName:   brideName,
			Information: "Chúng mình sẽ kết hôn",
			YoutubeLink: youtubeLink,
			ImageLink:   link,
		})
	}
	homePageData.updateCarouselData(carouselData)

	// Update story data
	updateHiepStory(homePageData)

	pageData := core.Page{
		PageFiles: []string{
			"./html/index.html",
			"./html/carousel.html",
			"./html/carousel_item.html",
			"./html/nav_bar.html",
			"./html/about.html",
			"./html/story.html",
			"./html/right_story.html",
			"./html/left_story.html",
			"./html/gallery.html",
			"./html/event.html",
			"./html/friends_and_family.html",
			"./html/rsvp.html",
			"./html/footer.html",
			"./html/maps.html",
			"./html/bank_accounts.html",
		},
		Data: *homePageData,
	}

	// Homepage
	core.RegisterPage("/hiep", pageData)
}

/*
* Update story data
* @param homePageData *HomePageData
* @return void
 */
func updateHiepStory(homePageData *HomePageData) {
	var storyLineTitle = "Hanh & Hiep's Story"
	var storyContents = []StoryContent{
		{
			Title:      "First Meet",
			Date:       "2015-01-01",
			Body:       "We met at a friend's party. We were both shy and didn't talk much. But we both felt something special. We started to talk more and more and we fell in love.",
			ImageStory: "img/story-1.jpg",
		},
		{
			Title:      "First Date",
			Date:       "2015-02-01",
			Body:       "We went to a coffee shop. We talked for hours. We found out that we have a lot in common. We both love to travel and we both love to eat.",
			ImageStory: "img/story-2.jpg",
		},
		{
			Title:      "Proposal",
			Date:       "2016-01-01",
			Body:       "I proposed to her on a beach. It was a beautiful sunset. She said yes. We were both very happy.",
			ImageStory: "img/story-3.jpg",
		},
		{
			Title:      "Enagagement",
			Date:       "2016-02-01",
			Body:       "We had a small party with our friends and family. We were both very happy.",
			ImageStory: "img/story-4.jpg",
		},
	}
	homePageData.Story.StoryLineTitle = storyLineTitle
	homePageData.Story.StoryContents = storyContents
}
