package main

import (
	"log"
	"net/http"

	"github.com/cocotton/pancarte/pancarte/door"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

func main() {
	s, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()
	s.SetMode(mgo.Monotonic, true)

	router := mux.NewRouter()
	router.HandleFunc("/addDoor", func(w http.ResponseWriter, r *http.Request) {
		door.addDoor(w, r, s)
	}).Methods("POST")
	router.HandleFunc("/getDoor/{doorID}", func(w http.ResponseWriter, r *http.Request) {
		door.getDoor(w, r, s)
	}).Methods("GET")
	http.ListenAndServe(":8080", router)
}
