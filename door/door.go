package door

import (
	"errors"
	"strings"

	"github.com/cocotton/pancarte/location"
)

// Door is a struct containing all the data concerning a Door
type Door struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Price       string `json:"price"`
	OwnerName   string `json:"ownerName"`
	OwnerPhone  string `json:"ownerPhone"`
	Location    location.Location
}

// ValidateDoor makes sure no fields are empty in the Door object, execpt for ID.
func ValidateDoor(door *Door) error {
	emptyFields := []string{}

	if door.Description == "" {
		emptyFields = append(emptyFields, "description")
	}
	if door.Price == "" {
		emptyFields = append(emptyFields, "price")
	}
	if door.OwnerName == "" {
		emptyFields = append(emptyFields, "owner's name")
	}
	if door.OwnerPhone == "" {
		emptyFields = append(emptyFields, "owner's phone")
	}
	if len(emptyFields) > 0 {
		return errors.New("Empty fields: " + strings.Join(emptyFields, ","))
	}

	return nil
}
