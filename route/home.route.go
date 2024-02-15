package route

type HomePageData struct {
	Title            string
	Carousel         Carousel
	VideoModal       VideoModal
	NavBar           NavBar
	About            About
	Story            Story
	Gallery          Gallery
	Event            Event
	FriendsAndFamily FriendsAndFamily
	RSVP             RSVP
	Footer           Footer
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

type VideoModal struct {
}

type About struct {
	AboutTitle string
	GroomTitle string
	GroomAbout string
	BrideTitle string
	BrideAbout string
	GroomName  string
	BrideName  string
}

type Story struct {
	StoryTitle     string
	StoryLineTitle string
	StoryContents  []StoryContent
}

type StoryContent struct {
	Title string
	Date  string
	Body  string
}

type Gallery struct {
}

type Event struct {
}

type FriendsAndFamily struct {
}

type RSVP struct {
}

type Footer struct {
}
