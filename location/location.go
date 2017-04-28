package location

// Location contains all the informations used to locate a Door
type Location struct {
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Country     string  `json:"country"`
	Province    string  `json:"province"`
	City        string  `json:"city"`
	StreetName  string  `json:"streetname"`
	CivicNumber string  `json:"civicnumber"`
	PostalCode  string  `json:"postalcode"`
}
