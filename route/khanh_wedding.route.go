package route

import "github.com/dangviethung096/core"

func khanhWedding() {
	var groomName = "Khánh"
	var brideName = "Cúc"
	var groomFullName = "Phạm Duy Khánh"
	var brideFullName = "Nguyễn Thu Cúc"
	var homeTitle = "Trang chủ"
	var aboutTitle = "Về chúng tớ"
	var galleryTitle = "Thư viện ảnh"
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
		"img/khanh/home/home_01.jpg",
		"img/khanh/home/home_02.jpg",
		"img/khanh/home/home_03.jpg",
		"img/khanh/home/home_04.jpg",
		"img/khanh/home/home_05.jpg",
	}

	var youtubeLink = "https://www.youtube.com/embed/w_17Bz-ngVM?si=vJenP1mx8mXrL5QF"

	homePageData := &HomePageData{
		Title: "Đám cưới chúng mình",
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
			GroomImage: "img/khanh/about/chu_re.jpg",
			BrideImage: "img/khanh/about/co_dau.jpg",
		},
		Story: Story{
			StoryTitle: storyTitle,
		},
		Gallery: Gallery{},
		Event: Event{
			EventTitle:        eventTitle,
			EventContent:      "Trân trọng kính mời bạn và người thương tới dự buổi tiệc chung vui cùng chúng mình",
			BrideEventTitle:   "Lễ vu quy",
			BrideEventContent: "Tại nhà gái - Nhà văn hoá thôn Quảng Minh, xã Mỹ Hưng, huyện Thanh Oai, Hà Nội",
			BrideEventDate:    "10:00 AM 13/04/2024",
			GroomEventTitle:   "Lễ thành hôn",
			GroomEventContent: "Tại nhà trai - 155 Phố Lưu Trọng Lư, Phúc Lộc, Hưng Đạo, Dương Kinh, Hải Phòng",
			GroomEventDate:    "10:00 AM 14/04/2024",
			EventImage1:       "img/khanh/event/1.jpg",
			EventImage2:       "img/khanh/event/2.jpg",
		},
		FriendsAndFamily: FriendsAndFamily{},
		RSVP:             RSVP{},
		Footer:           Footer{},
		Maps: Maps{
			MapsTitle:         "Địa chỉ",
			MapsTitleMessage:  "Địa chỉ nhà chúng mình",
			GroomAddressTitle: "Nhà trai",
			BrideAddressTitle: "Nhà gái",
			GroomAddress:      "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3729.7016885128496!2d106.67318267574507!3d20.803353080790245!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x314a70f62e95b357%3A0xbc7494c4d829e822!2zMTU1IFBo4buRIEzGsHUgVHLhu41uZyBMxrAsIEjGsG5nIMSQ4bqhbywgRMawxqFuZyBLaW5oLCBI4bqjaSBQaMOybmcsIFZpZXRuYW0!5e0!3m2!1sen!2s!4v1712238300217!5m2!1sen!2s",
			BrideAddress:      "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3727.0102549693706!2d105.78920497574764!3d20.91190708070844!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x3135b357485e67db%3A0x7c0adc6e0ce91843!2zTmjDoCB2xINuIGjDs2E!5e0!3m2!1sen!2s!4v1712238371353!5m2!1sen!2s",
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
	updateKhanhStory(homePageData)

	homePageData.Gallery.GalleryTitle = galleryTitle
	homePageData.Gallery.GalleryContent = "Đây là những khoảnh khắc đẹp nhất của chúng mình"

	// Update image in gallery
	fileNames := getImageFileNamesInFolder("html/img/khanh/gallery")
	for _, fileName := range fileNames {
		homePageData.Gallery.GalleryImages = append(homePageData.Gallery.GalleryImages, GalleryImage{
			ImageLink: "img/khanh/gallery/" + fileName,
		})
	}

	// Update event data
	homePageData.Event.EventTitle = eventTitle
	homePageData.Event.EventContent = "Trân trọng kính mời bạn và người thương tới dự buổi tiệc chung vui cùng chúng mình"

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
			"./html/gallery_item.html",
		},
		Data: *homePageData,
	}
	// Homepage
	core.RegisterPage("/", pageData)
}

func updateKhanhStory(homePageData *HomePageData) {
	var storyLineTitle = homePageData.NavBar.GroomName + " & " + homePageData.NavBar.BrideName + " 's Story"
	var storyContents = []StoryContent{
		{
			Title:      "First Meet",
			Date:       "2015-01-01",
			Body:       "We met at a friend's party. We were both shy and didn't talk much. But we both felt something special. We started to talk more and more and we fell in love.",
			ImageStory: "img/khanh/story/story_1.jpg",
		},
		{
			Title:      "First Date",
			Date:       "2015-02-01",
			Body:       "We went to a coffee shop. We talked for hours. We found out that we have a lot in common. We both love to travel and we both love to eat.",
			ImageStory: "img/khanh/story/story_2.jpg",
		},
		{
			Title:      "Proposal",
			Date:       "2016-01-01",
			Body:       "I proposed to her on a beach. It was a beautiful sunset. She said yes. We were both very happy.",
			ImageStory: "img/khanh/story/story_3.jpg",
		},
		{
			Title:      "Enagagement",
			Date:       "2016-02-01",
			Body:       "We had a small party with our friends and family. We were both very happy.",
			ImageStory: "img/khanh/story/story_4.jpg",
		},
	}
	homePageData.Story.StoryLineTitle = storyLineTitle
	homePageData.Story.StoryContents = storyContents
}
