package door

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type door struct {
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
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newDoor)
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB("pancarte").C("doors")

	err = c.Insert(newDoor)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
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
