package main

import (
	"bytes"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var JsonStr = []byte(`[
		  {
			"id": 1,
			"type": "dog",
			"price": 249.99
		  },
		  {
			"id": 2,
			"type": "cat",
			"price": 124.99
		  },
		  {
			"id": 3,
			"type": "fish",
			"price": 0.99
		  }
		]`)

func TestPostPets(t *testing.T) {
	req, err := http.NewRequest("POST", "/pets", bytes.NewBuffer(JsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostPets)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetPets(t *testing.T) {
	req, err := http.NewRequest("POST", "/pets", bytes.NewBuffer(JsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostPets)
	handler.ServeHTTP(rr, req)

	req, err = http.NewRequest("GET", "/pets", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(GetPets)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"Id":1,"Type":"dog","Price":249.99},{"Id":2,"Type":"cat","Price":124.99},{"Id":3,"Type":"fish","Price":0.99}]`
	if strings.TrimRight(rr.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetPetsByID(t *testing.T) {
	req, err := http.NewRequest("POST", "/pets", bytes.NewBuffer(JsonStr))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostPets)
	handler.ServeHTTP(rr, req)

	req, err = http.NewRequest("GET", "/pets/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	urlVars := map[string]string{
		"petId": "1",
	}
	req = mux.SetURLVars(req, urlVars)

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(GetPetsByID)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"Id":1,"Type":"dog","Price":249.99}`
	if strings.TrimRight(rr.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
