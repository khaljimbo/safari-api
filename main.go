package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// SafariPark is exported to keep the linter happy
type SafariPark struct {
	ID         string   `json:"id,omitempty"`
	SafariPark string   `json:"safaripark,omitempty"`
	Big5       string   `json:"big5,omitempty"`
	Address    *address `json:"address,omitempty"`
}

type address struct {
	State string `json:"state,omitempty"`
}

var park []SafariPark

func getSafariParkEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range park {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&SafariPark{})
}

func getParksEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(park)
}

func createSafariParkEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var safariPark SafariPark
	_ = json.NewDecoder(req.Body).Decode(&safariPark)
	safariPark.ID = params["id"]
	park = append(park, safariPark)
	json.NewEncoder(w).Encode(park)
}

func deleteSafariParkEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range park {
		if item.ID == params["id"] {
			park = append(park[:index], park[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(park)
}

func main() {
	router := mux.NewRouter()
	park = append(park, SafariPark{ID: "1", SafariPark: "Kruger Park", Big5: "Yes", Address: &address{State: "Limpopo"}})
	park = append(park, SafariPark{ID: "2", SafariPark: "Addo", Big5: "No", Address: &address{State: "Eastern Cape"}})
	router.HandleFunc("/park", getSafariParkEndpoint).Methods("GET")
	router.HandleFunc("/park/{id}", getParksEndpoint).Methods("GET")
	router.HandleFunc("/park/{id}", createSafariParkEndpoint).Methods("POST")
	router.HandleFunc("/park/{id}", deleteSafariParkEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":12345", router))
}
