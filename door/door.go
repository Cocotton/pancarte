package door

import (
	"errors"
	"strings"

	"github.com/cocotton/pancarte/location"
)

// Door is a struct containing all the data concerning a Door
type Door struct {
	ID          string            `json:"id"`
	Title       string            `json:"title"`
	Price       string            `json:"price"`
	Currency    string            `json:"currency"`
	Description string            `json:"description"`
	OwnerName   string            `json:"ownerName"`
	OwnerPhone  string            `json:"ownerPhone"`
	Location    location.Location `json:"location"`
}

// ValidateDoor makes sure no fields are empty in the Door object, execpt for ID.
func ValidateDoor(door *Door) error {
	emptyFields := []string{}

	if door.Title == "" {
		emptyFields = append(emptyFields, "Title")
	}
	if door.Price == "" {
		emptyFields = append(emptyFields, "Price")
	}
	if door.Currency == "" {
		emptyFields = append(emptyFields, "Currency")
	}
	if door.Description == "" {
		emptyFields = append(emptyFields, "Description")
	}
	if door.OwnerName == "" {
		emptyFields = append(emptyFields, "Owner's name")
	}
	if door.OwnerPhone == "" {
		emptyFields = append(emptyFields, "Owner's phone")
	}

	if len(emptyFields) > 0 {
		return errors.New("Empty fields: " + strings.Join(emptyFields, ","))
	}

	return location.ValidateLocation(door.Location)
}
