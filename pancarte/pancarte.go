package pancarte

import (
	"log"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

// Pancarte holds all the information required by the app to run
type Pancarte struct {
	DBSession *mgo.Session
	Database  string
	Router    *mux.Router
}

// InitDB initializes the connection to the database and its indexes
func (p *Pancarte) InitDB(host string, database string) {
	var err error

	p.Database = database
	p.DBSession, err = mgo.Dial(host)
	if err != nil {
		handleFatalInitError("Unable to initialize the connection to the databse.", err)
	}
	defer p.DBSession.Close()

	p.DBSession.SetMode(mgo.Monotonic, true)

	p.initDoorIndex()
}

func (p *Pancarte) initDoorIndex() {
	index := mgo.Index{
		Key:        []string{"id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	c := p.DBSession.DB(p.Database).C("doors")

	err := c.EnsureIndex(index)
	if err != nil {
		handleFatalInitError("Can ensure the doors collection indexes.", err)
	}
}

// InitRouter initializes the mux Router and its routes
func (p *Pancarte) InitRouter() {
	p.Router = mux.NewRouter()
}

func handleFatalInitError(message string, err error) {
	log.Fatalf(message+"\nError: %s", err)
}
