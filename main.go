package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Pet struct {
	Id    int
	Type  string
	Price float64
}

var Pets []Pet

func GetPets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Pets)
}

func PostPets(w http.ResponseWriter, r *http.Request) {
	var pets []Pet

	err := json.NewDecoder(r.Body).Decode(&pets)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Pets = append(Pets, pets...)
}

func GetPetsByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	petID := params["petId"]

	intPetID, err := strconv.Atoi(petID)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	if len(Pets) > 0 {
		for _, pet := range Pets {
			if pet.Id == intPetID {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(pet)
				return
			}
		}
	}

	http.Error(w, "petId not found", http.StatusNotFound)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/pets", GetPets).Methods("GET")
	r.HandleFunc("/pets", PostPets).Methods("POST")
	r.HandleFunc("/pets/{petId}", GetPetsByID).Methods("GET")
	http.ListenAndServe(":8081", r)
}
