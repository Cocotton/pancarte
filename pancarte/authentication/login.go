package authentication

import (
	"encoding/json"
	"net/http"

	"github.com/cocotton/pancarte/pancarte/response"
	"github.com/cocotton/pancarte/pancarte/user"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login makes sure the user exists in the database and that its password match
func Login(w http.ResponseWriter, r *http.Request, session *mgo.Session) {
	l := login{}

	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		response.ErrorWithText(w, "Malformated username/password", http.StatusBadRequest)
		return
	}

	fetchedUser := user.User{}
	c := session.DB("pancarte").C("users")
	err = c.Find(bson.M{"username": l.Username}).One(&fetchedUser)
	if err != nil {
		response.ErrorWithText(w, "Unable to find user: "+l.Username, http.StatusNotFound)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(fetchedUser.Password), []byte(l.Password)); err != nil {
		response.ErrorWithText(w, "Password does not match", http.StatusForbidden)
		return
	}
	response.SuccessWithJSON(w, []byte("Password matched"), http.StatusOK)
}
