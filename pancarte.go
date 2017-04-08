package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type door struct {
	Address     string `json:"address"`
	Description string `json:"descrition"`
	Price       string `json:"price"`
	OwnerName   string `json:"ownerName"`
	OwnerPhone  string `json:"ownerPhone"`
}

var doors []door

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/addDoor", addDoor).Methods("POST")
	http.ListenAndServe(":8080", router)
}

func addDoor(w http.ResponseWriter, r *http.Request) {
	newDoor := new(door)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newDoor)
	if err != nil {
		log.Fatal(err)
	}
}
