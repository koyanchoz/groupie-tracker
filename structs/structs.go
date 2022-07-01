package structs

type Links struct {
	ArtistsURL   string `json:"artists"`
	LocationsURL string `json:"locations"`
	DatesURL     string `json:"dates"`
	RelationURL  string `json:"relation"`
}

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Relatations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// type Locations struct {
// 	ID        int      `json:"id"`
// 	Locations []string `json:"locations"`
// 	Dates     string   `json:"dates"`
// }

// type Dates struct {
// 	ID    int      `json:"id"`
// 	Dates []string `json:"dates"`
// }

type About struct {
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Relations    map[string][]string
}
