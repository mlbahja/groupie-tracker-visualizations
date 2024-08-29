package global

type ApiOfArtist struct {
	ArtistData ArtistData `json:"artists"`
	Locations  Locations  `json:"locations"`
	Dates      Dates      `json:"dates"`
	Relation   Relations  `json:"relation"`
}

type ArtistData struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

const Api = "https://groupietrackers.herokuapp.com/api"
