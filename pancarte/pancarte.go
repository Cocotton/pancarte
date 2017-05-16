package pancarte

import (
	"errors"
	"log"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

// Pancarte holds all the information required by the app to run
type Pancarte struct {
	DBDoorCollection     string
	DBDoorCounterID      string
	DBUserCollection     string
	DBCountersCollection string
	DBName               string
	DBSession            *mgo.Session
	JWTSecret            string
	Router               *mux.Router
}

// InitDB initializes the connection to the database and its indexes
func (p *Pancarte) InitDB(host string, dbName string) {
	var err error

	p.DBName = dbName
	p.DBDoorCollection = "doors"
	p.DBDoorCounterID = "doorid"
	p.DBUserCollection = "users"
	p.DBCountersCollection = "counters"
	p.DBSession, err = mgo.Dial(host)
	if err != nil {
		handleFatalInitError("Unable to initialize the connection to the databse.", err)
	}

	p.DBSession.SetMode(mgo.Monotonic, true)

	p.initDoorIndex()
	p.initGeoLocationIndex()
	p.initUserIndex()
}

func (p *Pancarte) initDoorIndex() {
	index := mgo.Index{
		Key:        []string{"id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	c := p.DBSession.DB(p.DBName).C(p.DBDoorCollection)
	err := c.EnsureIndex(index)
	if err != nil {
		handleFatalInitError("Can't ensure the doors collection index.", err)
	}
}

func (p *Pancarte) initGeoLocationIndex() {
	index := mgo.Index{
		Key:        []string{"$2dsphere:location.geolocation"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	c := p.DBSession.DB(p.DBName).C(p.DBDoorCollection)
	err := c.EnsureIndex(index)
	if err != nil {
		handleFatalInitError("Can't ensure the geolocation index.", err)
	}
}

func (p *Pancarte) initUserIndex() {
	index := mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	c := p.DBSession.DB(p.DBName).C(p.DBUserCollection)
	err := c.EnsureIndex(index)
	if err != nil {
		handleFatalInitError("Can't ensure the users collection index", err)
	}
}

// InitRouter initializes the mux Router and its routes
func (p *Pancarte) InitRouter() {
	p.Router = mux.NewRouter()

	p.Router.HandleFunc("/addDoor", p.validateJWTHandler(p.addDoorHandler)).Methods("POST")
	p.Router.HandleFunc("/addUser", p.validateJWTHandler(p.addUserHandler)).Methods("POST")
	p.Router.HandleFunc("/getDoor/{doorID}", p.getDoorHandler).Methods("GET")
	p.Router.HandleFunc("/health", p.healthHandler).Methods("GET")
	p.Router.HandleFunc("/login", p.loginHandler).Methods("POST")
	p.Router.HandleFunc("/logout", p.logoutHandler).Methods("GET")
}

// SetJWTSecret sets the JWT secrets in the object
func (p *Pancarte) SetJWTSecret(secret string) {
	if len(secret) == 0 {
		err := errors.New("JWT Secret is empty")
		handleFatalInitError("", err)
	}
	p.JWTSecret = secret
}

func handleFatalInitError(message string, err error) {
	log.Fatalf(message+"\nError: %s", err)
}
