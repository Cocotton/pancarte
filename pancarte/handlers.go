package pancarte

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/cocotton/pancarte/door"
	"github.com/cocotton/pancarte/helpers"
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
		helpers.ErrorWithText(w, err, "Some fields are missing in the new Door.", http.StatusBadRequest)
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

	helpers.SuccessJSONLogger(w, "Succesfully created door with id: "+newDoor.ID, http.StatusCreated)
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

	helpers.SuccessJSONLogger(w, string(res), http.StatusOK)
}
