package door

import (
	"errors"
	"reflect"
)

// Door is a struct containing all the data concerning a Door
type Door struct {
	ID          string `json:"id"`
	Address     string `json:"address"`
	Description string `json:"description"`
	Price       string `json:"price"`
	OwnerName   string `json:"ownerName"`
	OwnerPhone  string `json:"ownerPhone"`
}

// ValidateDoor makes sure no fields are empty in the Door object, execpt for ID.
func ValidateDoor(door *Door) error {
	r := reflect.ValueOf(door).Elem()

	for i := 1; i < r.NumField(); i++ {
		if r.Field(i).Len() == 0 {
			return errors.New("Empty field: " + r.Type().Field(i).Name)
		}
	}
	return nil
}
