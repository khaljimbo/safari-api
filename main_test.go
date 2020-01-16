package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/park", getSafariParkEndpoint).Methods("GET")
	return router
}

func TestGetSafariEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/park", nil)
	response := httptest.NewRecorder()
	GetRouter().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func GetParksRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/park", getParksEndpoint).Methods("GET")
	return router
}

func TestGetParksEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/park", nil)
	response := httptest.NewRecorder()
	GetParksRouter().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func DeleteRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/park", deleteSafariParkEndpoint).Methods("DELETE")
	return router
}

func TestDeleteSafariEndpoint(t *testing.T) {
	request, _ := http.NewRequest("DELETE", "/park", nil)
	response := httptest.NewRecorder()
	DeleteRouter().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

//Failing and I can't work out why right now. Will revisit
func CreateParkRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/park", createSafariParkEndpoint).Methods("POST")
	return router
}

func TestCreateSafariParkEndpoint(t *testing.T) {
	request, _ := http.NewRequest("POST", "/park", nil)
	response := httptest.NewRecorder()
	CreateParkRouter().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}
