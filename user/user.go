package user

// User contains all the information concerning a user
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
