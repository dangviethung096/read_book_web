package route

type HomePageData struct {
	Title            string
	Carousel         Carousel
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

type VideoModel struct {
}

type About struct {
}

type Story struct {
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
