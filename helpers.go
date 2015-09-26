package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Recover is generic function that recovers from panic
func Recover(functionName string) {
	if r := recover(); r != nil {
		log.Println("\trecovered in function", functionName, "from error:", r)
	}
}

// JSONerr writes json error in format {err: msg}
func JSONerr(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	errMsg := fmt.Sprintf("{\"err\":\"%v\"}", err.Error())
	http.Error(w, errMsg, http.StatusInternalServerError)
}

// WriteJSON hson gets interface and creates a json response from it
func WriteJSON(w http.ResponseWriter, i interface{}) error {

	defer Recover("WriteJson")

	var j []byte

	j, err := json.Marshal(i)
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return err
	}

	// set Headers
	w.Header().Set("Content-Type", "application/json")
	// write to writer
	w.Write(j)
	return nil
}
