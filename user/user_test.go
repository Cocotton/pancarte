package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateUser(t *testing.T) {
	assert := assert.New(t)
	var validateuserTests = []struct {
		user        *User
		expected    error
		expectError bool
	}{
		{&User{"username1", "john", "smith", "1qaz2wsx", "1234567890", "john@smith.com"}, nil, false},
		{&User{"username1", "john", "", "1qaz2wsx", "1234567890", "john@smith.com"}, nil, true},
	}

	for _, test := range validateuserTests {
		actual := ValidateUser(*test.user)
		if test.expectError {
			assert.Error(actual, "Expected an error. None was returned.")
		} else {
			assert.Equal(nil, actual, "Did not expected error. Received one.")
		}
	}
}
