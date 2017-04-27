package user

import (
	"errors"
	"reflect"

	"gopkg.in/mgo.v2"
)

// User contains all the information concerning a user
type User struct {
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

// Login contains the username/password provided during login
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AddUser adds a giving user to the database
func AddUser(collection *mgo.Collection, user User) error {
	err := ValidateUser(user)
	if err != nil {
		return err
	}

	err = collection.Insert(user)
	if err != nil {
		return err
	}

	return nil
}

// ValidateUser makes sure no fields are empty in the User object
func ValidateUser(user User) error {
	r := reflect.ValueOf(&user).Elem()

	for i := 0; i < r.NumField(); i++ {
		if r.Field(i).Len() == 0 {
			return errors.New("Empty fields: " + r.Type().Field(i).Name)
		}
	}
	return nil
}
