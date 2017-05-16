package location

import (
	"errors"
	"strings"
)

// Location contains all the informations used to locate a Door
type Location struct {
	GeoLocation GeoLocation `json:"geolocation"`
	Country     string      `json:"country"`
	Province    string      `json:"province"`
	City        string      `json:"city"`
	StreetName  string      `json:"streetname"`
	CivicNumber string      `json:"civicnumber"`
	PostalCode  string      `json:"postalcode"`
}

// GeoLocation contains a Geographic position
type GeoLocation struct {
	Type        string    `json:"geotype"`
	Coordinates []float64 `json:"geocoordinates"`
}

// ValidateLocation makes sure no fields are empty in the Location object.
func ValidateLocation(location Location) error {
	emptyFields := []string{}

	if location.GeoLocation.Type == "" {
		emptyFields = append(emptyFields, "GeoLocation Type")
	}
	if len(location.GeoLocation.Coordinates) < 2 {
		emptyFields = append(emptyFields, "GeoLocation Coordinates")
	}
	if location.Country == "" {
		emptyFields = append(emptyFields, "Country")
	}
	if location.Province == "" {
		emptyFields = append(emptyFields, "Province")
	}
	if location.City == "" {
		emptyFields = append(emptyFields, "City")
	}
	if location.StreetName == "" {
		emptyFields = append(emptyFields, "StreetName")
	}
	if location.CivicNumber == "" {
		emptyFields = append(emptyFields, "CivicNumber")
	}
	if location.PostalCode == "" {
		emptyFields = append(emptyFields, "PostalCode")
	}

	if len(emptyFields) > 0 {
		return errors.New("Empty fields: " + strings.Join(emptyFields, ","))
	}
	return nil
}
