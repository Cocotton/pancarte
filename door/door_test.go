package door

import (
	"testing"

	"github.com/cocotton/pancarte/location"
	"github.com/stretchr/testify/assert"
)

func TestValidateDoor(t *testing.T) {
	assert := assert.New(t)

	loc := location.Location{
		GeoLocation: location.GeoLocation{Type: "Point", Coordinates: []float64{-73.583008, 45.494660}},
		Country:     "Canada",
		Province:    "Qc",
		City:        "Montréal",
		StreetName:  "Rue Saint-Marc",
		CivicNumber: "2295",
		PostalCode:  "H3H2G9",
	}

	var validateDoorTests = []struct {
		door        *Door
		expectError bool
	}{
		{&Door{"1", "House", "500", "CAD", "House description", "John Smith", "123-456-7890", loc}, false},
		{&Door{"", "House", "500", "CAD", "House description", "John Smith", "123-456-7890", loc}, false},
		{&Door{"", "", "500", "CAD", "House description", "John Smith", "123-456-7890", loc}, true},
		{&Door{"", "House", "", "CAD", "House description", "John Smith", "123-456-7890", loc}, true},
		{&Door{"", "House", "500", "", "House description", "John Smith", "123-456-7890", loc}, true},
		{&Door{"", "House", "500", "CAD", "", "John Smith", "123-456-7890", loc}, true},
		{&Door{"", "House", "500", "CAD", "House description", "", "123-456-7890", loc}, true},
		{&Door{"", "House", "500", "CAD", "House description", "John Smith", "", loc}, true},
	}

	for _, test := range validateDoorTests {
		actual := ValidateDoor(test.door)
		if test.expectError {
			assert.Error(actual, "Expected an error. None was returned.")
		} else {
			assert.NoError(actual, "Did not expected error. Received one.")
		}
	}
}
