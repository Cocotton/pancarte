package pancarte

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"gopkg.in/mgo.v2/bson"

	"github.com/cocotton/pancarte/authentication"
	"github.com/cocotton/pancarte/door"
	"github.com/cocotton/pancarte/helpers"
	"github.com/cocotton/pancarte/user"
	"github.com/gorilla/mux"
)

func (p *Pancarte) addDoorHandler(w http.ResponseWriter, r *http.Request) {
	session := p.DBSession.Copy()
	defer session.Close()

	newDoor := door.Door{}
	err := json.NewDecoder(r.Body).Decode(&newDoor)
	if err != nil {
		helpers.ErrorWithText(w, err, "Door object is malformed.", http.StatusBadRequest)
		return
	}

	err = door.ValidateDoor(&newDoor)
	if err != nil {
		helpers.ErrorWithText(w, err, "Some fields are missing. "+err.Error(), http.StatusBadRequest)
		return
	}

	newDoor.ID, err = helpers.MongoGetNextID(session, p.DBName, p.DBCountersCollection, p.DBDoorCounterID)
	if err != nil {
		helpers.ErrorWithText(w, err, "Database error.", http.StatusInternalServerError)
		return
	}

	collection := session.DB(p.DBName).C(p.DBDoorCollection)
	err = collection.Insert(newDoor)
	if err != nil {
		helpers.ErrorWithText(w, err, "Database error.", http.StatusInternalServerError)
		return
	}

	helpers.SuccessWithJSON(w, "Succesfully created door with id: "+newDoor.ID, http.StatusCreated)
}

func (p *Pancarte) addUserHandler(w http.ResponseWriter, r *http.Request) {
	session := p.DBSession.Copy()
	defer session.Close()

	newUser := user.User{}
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		helpers.ErrorWithText(w, err, "User object is malformed.", http.StatusBadRequest)
		return
	}

	err = user.ValidateUser(newUser)
	if err != nil {
		helpers.ErrorWithText(w, errors.New("Error adding new user."), err.Error(), http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		helpers.ErrorWithText(w, err, "Unable to create new user.", http.StatusInternalServerError)
		return
	}
	newUser.Password = string(hash)

	collection := session.DB(p.DBName).C(p.DBUserCollection)
	err = user.AddUser(collection, newUser)
	if err != nil {
		helpers.ErrorWithText(w, err, "Unable to create new user.", http.StatusInternalServerError)
		return
	}

	helpers.SuccessWithJSON(w, "Succesfully created user with username: "+newUser.Username, http.StatusCreated)
}

func (p *Pancarte) getDoorHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	doorID := vars["doorID"]
	fetchedDoor := door.Door{}

	session := p.DBSession.Copy()
	defer session.Close()

	collection := session.DB(p.DBName).C(p.DBDoorCollection)

	err := collection.Find(bson.M{"id": doorID}).One(&fetchedDoor)
	if err != nil {
		helpers.ErrorWithText(w, err, "Unable to find door with ID: "+doorID, http.StatusNotFound)
		return
	}

	res, err := json.Marshal(fetchedDoor)
	if err != nil {
		helpers.ErrorWithText(w, err, "Problem with fetched door.", http.StatusInternalServerError)
		return
	}

	helpers.SuccessWithJSON(w, string(res), http.StatusOK)
}

func (p *Pancarte) loginHandler(w http.ResponseWriter, r *http.Request) {
	loginInfo := user.Login{}

	err := json.NewDecoder(r.Body).Decode(&loginInfo)
	if err != nil {
		helpers.ErrorWithText(w, err, "Malformed username/password", http.StatusBadRequest)
		return
	}

	session := p.DBSession.Copy()
	defer session.Close()

	fetchedUser := user.User{}
	collection := session.DB(p.DBName).C(p.DBUserCollection)
	err = collection.Find(bson.M{"username": loginInfo.Username}).One(&fetchedUser)
	if err != nil {
		helpers.ErrorWithText(w, err, "Username/Password mismatch", http.StatusNotFound)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(fetchedUser.Password), []byte(loginInfo.Password)); err != nil {
		helpers.ErrorWithText(w, err, "Username/Password mismatch", http.StatusForbidden)
		return
	}

	cookie, err := authentication.CreateJWTCookie(fetchedUser.Username, p.JWTSecret)
	if err != nil {
		helpers.ErrorWithText(w, err, "Something went wrong", http.StatusInternalServerError)
	}

	http.SetCookie(w, cookie)
	helpers.SuccessWithJSON(w, "User logged in", http.StatusOK)
}

func (p *Pancarte) logoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "Auth", Value: "none", Expires: time.Now()})
	helpers.SuccessWithJSON(w, "User logged out", http.StatusOK)
}

func (p *Pancarte) validateJWTHandler(protectedPage http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Auth")
		if err != nil {
			helpers.ErrorWithText(w, err, "User not logged in", http.StatusForbidden)
			return
		}

		token, err := authentication.GetJWT(*cookie, p.JWTSecret)
		if err != nil {
			helpers.ErrorWithText(w, err, "JWT is not valid", http.StatusForbidden)
		}

		claims, err := authentication.GetJWTClaims(token)
		if err != nil {
			helpers.ErrorWithText(w, err, "Something went wrong", http.StatusInternalServerError)
		}

		context := authentication.GetContextWithClaims(r, claims)

		protectedPage(w, r.WithContext(context))
	})
}
