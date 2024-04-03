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
	Maps             Maps
	FriendsAndFamily FriendsAndFamily
	RSVP             RSVP
	Footer           Footer
	BankAccounts     BankAccounts
}

/*
* Update carousel data
* @param data []ImageCarousel
* @return void
 */
func (h *HomePageData) updateCarouselData(data []ImageCarousel) {
	for i, carousel := range data {
		if i == 0 {
			h.Carousel.ActiveImageCarousel = carousel
			continue
		}
		h.Carousel.ImageCarousel = append(h.Carousel.ImageCarousel, carousel)
	}
}

type Carousel struct {
	ActiveImageCarousel ImageCarousel
	ImageCarousel       []ImageCarousel
}

type NavBar struct {
	GroomName   string
	BrideName   string
	NavHome     string
	NavAbout    string
	NavGallery  string
	NavStory    string
	NavFamily   string
	NavEvent    string
	NavRSVP     string
	NavContact  string
	NavLocation string
	NavBank     string
}

type VideoModal struct {
	YoutubeLink string
}

type About struct {
	AboutTitle string
	GroomTitle string
	GroomAbout string
	BrideTitle string
	BrideAbout string
	GroomName  string
	BrideName  string
	GroomImage string
	BrideImage string
}

type Story struct {
	StoryTitle     string
	StoryLineTitle string
	StoryContents  []StoryContent
}

type StoryContent struct {
	Title      string
	Date       string
	Body       string
	ImageStory string
}

type Gallery struct {
	Images []GalleryImage
}

type GalleryImage struct {
	ImageLink string
}

type Event struct {
}

type FriendsAndFamily struct {
}

type RSVP struct {
}

type Footer struct {
}

type ActiveImageCarousel struct {
	GroomName   string
	BrideName   string
	Information string
	YoutubeLink string
	ImageLink   string
}

type ImageCarousel struct {
	GroomName   string
	BrideName   string
	Information string
	YoutubeLink string
	ImageLink   string
}

type Maps struct {
}

type BankAccounts struct {
}
