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

// Init creates a Pancarte object and then call the init functions
func Init(dbName string) *Pancarte {
	var p Pancarte

	p.initDB(dbName)
	p.initRouter()

	return &p
}

func (p *Pancarte) initDB(dbName string) {
	var err error

	p.DB, err = mgo.Dial(dbName)
	if err != nil {
		log.Fatal(err)
	}
	p.DB.SetMode(mgo.Monotonic, true)
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
