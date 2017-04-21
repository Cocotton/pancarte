package door

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateDoor(t *testing.T) {
	assert := assert.New(t)
	var validateDoorTests = []struct {
		door        *Door
		expected    error
		expectedErr bool
	}{
		{&Door{"1", "1 Road", "House", "500$", "John Smith", "123-456-7890"}, nil, false},
		{&Door{"", "1 Road", "House", "500$", "John Smith", "123-456-7890"}, nil, false},
		{&Door{"1", "1 Road", "", "500$", "John Smith", "123-456-7890"}, errors.New("Missing field."), true},
	}

	for _, test := range validateDoorTests {
		actual := ValidateDoor(test.door)
		if test.expectedErr {
			assert.Error(actual, "Expected an error. None was returned.")
		} else {
			assert.Equal(nil, actual, "Did not expected error. Received one.")
		}
	}
}
