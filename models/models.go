package models

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`    // URL
	ConcertDates string   `json:"concertDates"` // URL
	Relations    string   `json:"relations"`    // URL
}

type Locations struct {
	Index []LocationIndex `json:"index"`
}

type LocationIndex struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"` //URL
}

type Dates struct {
	Index []DateIndex `json:"index"`
}

type DateIndex struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	Index []RelationIndex `json:"index"`
}

type RelationIndex struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
