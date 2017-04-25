package user

import "gopkg.in/mgo.v2"

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

// AddUser adds a user to the mongo collection
func AddUser(collection *mgo.Collection, user User) error {
	return nil
}
