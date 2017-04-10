package door

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/cocotton/pancarte/pancarte/response"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type door struct {
	ID          string `json:"id"`
	Address     string `json:"address"`
	Description string `json:"description"`
	Price       string `json:"price"`
	OwnerName   string `json:"ownerName"`
	OwnerPhone  string `json:"ownerPhone"`
}

//AddDoor creates a new door in the mongo database
func AddDoor(w http.ResponseWriter, r *http.Request, s *mgo.Session) {
	session := s.Copy()
	defer session.Close()

	newDoor := new(door)
	err := json.NewDecoder(r.Body).Decode(&newDoor)
	if err != nil {
		response.ErrorWithJSON(w, "Incorrect body", http.StatusInternalServerError)
	}

	err = validateDoor(newDoor)
	if err != nil {
		response.ErrorWithJSON(w, err.Error(), http.StatusBadRequest)
	}

	newDoor.ID, err = getNextID(session)
	if err != nil {
		response.ErrorWithJSON(w, err.Error(), http.StatusInternalServerError)
	}

	c := session.DB("pancarte").C("doors")
	err = c.Insert(newDoor)
	if err != nil {
		response.ErrorWithJSON(w, "Can't create the new door object in database", http.StatusInternalServerError)
	} else {
		response.ResponseWithJSON(w, []byte("Successfully created the new door"), http.StatusCreated)
	}
}

func validateDoor(door *door) error {
	r := reflect.ValueOf(door).Elem()

	for i := 1; i < r.NumField(); i++ {
		if r.Field(i).Len() == 0 {
			return errors.New("Empty field: " + r.Type().Field(i).Name)
		}
	}
	return nil
}

func getNextID(s *mgo.Session) (string, error) {
	var result bson.M

	c := s.DB("pancarte").C("counters")
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"counter": 1}},
		ReturnNew: true,
	}
	_, err := c.Find(bson.M{"_id": "doorid"}).Apply(change, &result)
	if err != nil {
		return "", errors.New("Can't create new ID")
	}

	return strconv.FormatFloat(result["counter"].(float64), 'f', -1, 64), nil
}

// GetDoor gets a door from the mongo database using the provided ID
func GetDoor(w http.ResponseWriter, r *http.Request, s *mgo.Session) {
	vars := mux.Vars(r)
	doorID := vars["doorID"]
	var result door

	session := s.Copy()
	defer session.Close()

	c := session.DB("pancarte").C("doors")

	err := c.Find(bson.M{"id": doorID}).One(result)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

}
