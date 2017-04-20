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
		helpers.ErrorWithText(w, err, "Door object is malformed.", http.StatusBadRequest)
		return
	}

	err = door.ValidateDoor(&newDoor)
	if err != nil {
		helpers.ErrorWithText(w, err, "Some fields are missing in the new Door.", http.StatusBadRequest)
		log.Fatal(err)
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

	//response.SuccessWithJSON(w, newDoorJSON, http.StatusCreated)
	helpers.SuccessJSONLogger(w, "Succesfully created door with id: "+newDoor.ID, http.StatusCreated)
}
