package main

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

var session *mgo.Session

func main() {
	var err error
	session, err = mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	router := mux.NewRouter()
	router.HandleFunc("/addDoor", addDoor).Methods("POST")
	router.HandleFunc("/getDoor/{doorID}", getDoor).Methods("GET")
	http.ListenAndServe(":8080", router)
}

func addDoor(w http.ResponseWriter, r *http.Request) {
	s := session.Copy()
	defer s.Close()

	newDoor := new(door)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newDoor)
	if err != nil {
		log.Fatal(err)
	}

	c := s.DB("pancarte").C("doors")

	err = c.Insert(newDoor)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
}
func getDoor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	doorID := vars["doorID"]
	var doorr door

	s := session.Copy()
	defer s.Close()

	c := s.DB("pancarte").C("doors")

	err := c.Find(bson.M{"id": doorID}).One(doorr)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(doorr)

}
