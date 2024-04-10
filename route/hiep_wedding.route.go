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
		"img/hiep/home/home01.jpg",
		"img/hiep/home/home02.jpg",
		"img/hiep/home/home03.jpg",
		"img/hiep/home/home04.jpg",
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
			GroomImage: "img/hiep/about/about_1.jpg",
			BrideImage: "img/hiep/about/about_2.jpg",
		},
		Story: Story{
			StoryTitle: storyTitle,
		},
		Gallery: Gallery{},
		Event: Event{
			EventTitle:        eventTitle,
			EventContent:      "Trân trọng kính mời bạn và người thương tới dự buổi tiệc chung vui cùng chúng mình",
			BrideEventTitle:   "Lễ vu quy",
			BrideEventContent: "Tại nhà gái - Đội 7, Xã Giao Châu, Huyện Giao Thủy, Nam Định",
			BrideEventDate:    "14/04/2024",
			GroomEventTitle:   "Lễ thành hôn",
			GroomEventContent: "Tại nhà trai - Xóm 6, Xã Hồng Thuận, Huyện Giao Thủy, Nam Định",
			GroomEventDate:    "14/04/2024",
			EventImage1:       "img/hiep/event/event-1.jpg",
			EventImage2:       "img/hiep/event/event-2.jpg",
		},
		FriendsAndFamily: FriendsAndFamily{},
		RSVP:             RSVP{},
		Footer:           Footer{},
		Maps: Maps{
			MapsTitle:         "Địa chỉ",
			MapsTitleMessage:  "Địa chỉ nhà chúng mình",
			GroomAddressTitle: "Nhà trai",
			BrideAddressTitle: "Nhà gái",
			GroomAddress:      "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3742.527571936929!2d106.47424307573255!3d20.278418681189773!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x313601dd15935db3%3A0x6e6ccd5d303643f0!2zVUJORCB4w6MgSOG7k25nIFRodeG6rW4!5e0!3m2!1sen!2s!4v1712761547259!5m2!1sen!2s",
			BrideAddress:      "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d4532.0766862429755!2d106.41182098874792!3d20.25519793072554!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x313606e126b89a01%3A0x9d6e037d0fa60eca!2zVUJORCB4w6MgR2lhbyBDaMOidQ!5e0!3m2!1sen!2s!4v1712761508761!5m2!1sen!2s",
		},
		BankAccounts: BankAccounts{
			GroomName:        groomFullName,
			GroomBankAccount: "PGBank",
			GroomBankQR:      "img/hiep/bank_account/groom.jpeg",
			BrideName:        brideFullName,
			BrideBankAccount: "VietinBank",
			BrideBankQR:      "img/hiep/bank_account/bride.jpeg",
		},
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

	homePageData.Gallery.GalleryTitle = galleryTitle
	homePageData.Gallery.GalleryContent = "Đây là những khoảnh khắc đẹp nhất của chúng mình"

	// Update image in gallery
	fileNames := getImageFileNamesInFolder("html/img/hiep/gallery")
	for _, fileName := range fileNames {
		homePageData.Gallery.GalleryImages = append(homePageData.Gallery.GalleryImages, GalleryImage{
			ImageLink: "img/hiep/gallery/" + fileName,
		})
	}

	// Update event data
	homePageData.Event.EventTitle = eventTitle
	homePageData.Event.EventContent = "Trân trọng kính mời bạn và người thương tới dự buổi tiệc chung vui cùng chúng mình"

	pageData := core.Page{
		PageFiles: []string{
			"./html/index_hiep.html",
			"./html/carousel.html",
			"./html/carousel_item.html",
			"./html/nav_bar_hiep.html",
			"./html/about.html",
			"./html/story.html",
			"./html/right_story.html",
			"./html/left_story.html",
			"./html/gallery.html",
			"./html/gallery_item.html",
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
