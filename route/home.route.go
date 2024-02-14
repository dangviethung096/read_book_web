package route

type HomePageData struct {
	Title        string
	Groom        string
	Bride        string
	WelcomeImage WelcomeImageData
	NavHome      string
	NavAbout     string
	NavGallery   string
	NavStory     string
	NavFamily    string
	NavEvent     string
	NavRSVP      string
	NavContact   string
}

type WelcomeImageData struct {
	Groom       string
	Bride       string
	Information string
	YoutubeLink string
}
