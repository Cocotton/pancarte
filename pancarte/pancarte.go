package pancarte

import (
	"log"
	"net/http"

	"github.com/cocotton/pancarte/pancarte/door"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

// Pancarte structure is used to manage Pancarte
type Pancarte struct {
	DB     *mgo.Session
	Router *mux.Router
}

// InitDB setups the initial DB connection and http routes
func (p *Pancarte) InitDB(dbName string) {
	var err error

	p.DB, err = mgo.Dial(dbName)
	if err != nil {
		log.Fatal(err)
	}
}

func (p *Pancarte) initRouter() {
	p.Router = mux.NewRouter()

	p.Router.HandleFunc("/addDoor", func(w http.ResponseWriter, r *http.Request) {
		door.AddDoor(w, r, p.DB)
	}).Methods("POST")
	p.Router.HandleFunc("/getDoor/{doorID}", func(w http.ResponseWriter, r *http.Request) {
		door.GetDoor(w, r, p.DB)
	}).Methods("GET")
}
