package pancarte

import (
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

// Pancarte holds all the information required by the app to run
type Pancarte struct {
	DBSession *mgo.Session
	Router    *mux.Router
}
