package pancarte

import (
	"log"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

// Pancarte holds all the information required by the app to run
type Pancarte struct {
	DBSession *mgo.Session
	Router    *mux.Router
}

// InitDB initializes
func (p *Pancarte) InitDB(host string) {
	var err error

	p.DBSession, err = mgo.Dial(host)
	if err != nil {
		handleFatalInitError("Unable to initialize the connection to the databse.", err)
	}
	defer p.DBSession.Close()

	p.DBSession.SetMode(mgo.Monotonic, true)
}

func handleFatalInitError(message string, err error) {
	log.Fatalf(message+"\nError: %s", err)
}
