package door

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"

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
		log.Fatal(err)
	}

	err = validateDoor(newDoor)
	if err != nil {
		log.Fatal(err)
	}

	newDoor.ID = getNextID(session)

	c := session.DB("pancarte").C("doors")
	err = c.Insert(newDoor)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
}

func validateDoor(door *door) error {
	fmt.Println("In validateDoor")
	fmt.Println(door)
	r := reflect.ValueOf(door).Elem()
	fmt.Println("---")
	fmt.Println(r)

	for i := 1; i < r.NumField(); i++ {
		if r.Field(i).Len() == 0 {
			return errors.New("Empty field: " + r.Type().Field(i).Name)
		}
	}
	return nil
}

func getNextID(s *mgo.Session) string {
	var result bson.M

	c := s.DB("pancarte").C("counters")
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"counter": 1}},
		ReturnNew: true,
	}
	_, _ = c.Find(bson.M{"_id": "doorid"}).Apply(change, &result)

	return strconv.FormatFloat(result["counter"].(float64), 'f', -1, 64)
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
