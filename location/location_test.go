package location

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateLocation(t *testing.T) {
	assert := assert.New(t)

	var validateLocationTests = []struct {
		location    Location
		expectError bool
	}{
		{Location{GeoLocation{"Point", []float64{-73.583008, 45.494660}}, "Canada", "Qc", "Montréal", "Rue Saint-Marc", "2295", "H3H2G9"}, false},
		{Location{GeoLocation{"", []float64{-73.583008, 45.494660}}, "Canada", "Qc", "Montréal", "Rue Saint-Marc", "2295", "H3H2G9"}, true},
		{Location{GeoLocation{"Point", []float64{45.494660}}, "Canada", "Qc", "Montréal", "Rue Saint-Marc", "2295", "H3H2G9"}, true},
		{Location{GeoLocation{"Point", []float64{-73.583008, 45.494660}}, "", "Qc", "Montréal", "Rue Saint-Marc", "2295", "H3H2G9"}, true},
		{Location{GeoLocation{"Point", []float64{-73.583008, 45.494660}}, "Canada", "", "Montréal", "Rue Saint-Marc", "2295", "H3H2G9"}, true},
		{Location{GeoLocation{"Point", []float64{-73.583008, 45.494660}}, "Canada", "Qc", "", "Rue Saint-Marc", "2295", "H3H2G9"}, true},
		{Location{GeoLocation{"Point", []float64{-73.583008, 45.494660}}, "Canada", "Qc", "Montréal", "", "2295", "H3H2G9"}, true},
		{Location{GeoLocation{"Point", []float64{-73.583008, 45.494660}}, "Canada", "Qc", "Montréal", "Rue Saint-Marc", "", "H3H2G9"}, true},
		{Location{GeoLocation{"Point", []float64{-73.583008, 45.494660}}, "Canada", "Qc", "Montréal", "Rue Saint-Marc", "2295", ""}, true},
	}

	for _, test := range validateLocationTests {
		actual := ValidateLocation(test.location)
		if test.expectError {
			assert.Error(actual, "Expected an error. None was returned.")
		} else {
			assert.NoError(actual, "Did not expected error. Received one.")
		}
	}
}

func TestValidateGeoLocation(t *testing.T) {
	assert := assert.New(t)

	var validateGeoLocationTests = []struct {
		geoLocation GeoLocation
		expectError bool
	}{
		{GeoLocation{"Point", []float64{-73.583008, 45.494660}}, false},
		{GeoLocation{"Point", []float64{45.494660}}, true},
		{GeoLocation{"", []float64{-73.583008, 45.494660}}, true},
	}

	for _, test := range validateGeoLocationTests {
		actual := ValidateGeoLocation(test.geoLocation)
		if test.expectError {
			assert.Error(actual, "Expected an error. None was returned.")
		} else {
			assert.NoError(actual, "Did not expected error. Received one.")
		}
	}
}
