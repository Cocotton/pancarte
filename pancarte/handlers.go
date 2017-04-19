package pancarte

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cocotton/pancarte/door"
	"github.com/cocotton/pancarte/helpers"
)

func (p *Pancarte) addDoorHandler(w http.ResponseWriter, r *http.Request) {
	session := p.DBSession.Copy()
	defer session.Close()

	newDoor := door.Door{}
	err := json.NewDecoder(r.Body).Decode(&newDoor)
	if err != nil {
		log.Fatal("Can't decode Door json body.")
	}

	err = door.ValidateDoor(&newDoor)
	if err != nil {
		//response.ErrorWithText(w, err.Error(), http.StatusBadRequest)
		log.Fatal(err)
		return
	}

	newDoor.ID, err = helpers.MongoGetNextID(session, p.DBName, p.DBCountersCollection, p.DBDoorCounterID)
	if err != nil {
		log.Fatal("Can't get the new door ID.")
		//response.ErrorWithText(w, err.Error(), http.StatusInternalServerError)
		return
	}

	collection := session.DB(p.DBName).C(p.DBDoorCollection)
	err = collection.Insert(newDoor)
	if err != nil {
		//response.ErrorWithText(w, "Can't create the new door object in database", http.StatusInternalServerError)
		log.Fatal("Can't create the new door object in datase")
		return
	}

	//response.SuccessWithJSON(w, newDoorJSON, http.StatusCreated)
	log.Println("Succesfully created door with id: " + newDoor.ID)
}
