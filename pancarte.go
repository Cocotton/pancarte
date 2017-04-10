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

func main() {
	s, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()
	s.SetMode(mgo.Monotonic, true)

	router := mux.NewRouter()
	router.HandleFunc("/addDoor", func(w http.ResponseWriter, r *http.Request) {
		addDoor(w, r, s)
	}).Methods("POST")
	router.HandleFunc("/getDoor/{doorID}", func(w http.ResponseWriter, r *http.Request) {
		getDoor(w, r, s)
	}).Methods("GET")
	http.ListenAndServe(":8080", router)
}

func addDoor(w http.ResponseWriter, r *http.Request, s *mgo.Session) {
	session := s.Copy()
	defer s.Close()

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
func getDoor(w http.ResponseWriter, r *http.Request, s *mgo.Session) {
	vars := mux.Vars(r)
	doorID := vars["doorID"]
	var doorr door

	session := s.Copy()
	defer session.Close()

	c := session.DB("pancarte").C("doors")

	err := c.Find(bson.M{"id": doorID}).One(doorr)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(doorr)

}
