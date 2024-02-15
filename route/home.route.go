package route

type HomePageData struct {
	Title    string
	Carousel Carousel
	NavBar   NavBar
}

type Carousel struct {
	GroomName   string
	BrideName   string
	Information string
	YoutubeLink string
}

type NavBar struct {
	GroomName  string
	BrideName  string
	NavHome    string
	NavAbout   string
	NavGallery string
	NavStory   string
	NavFamily  string
	NavEvent   string
	NavRSVP    string
	NavContact string
}
