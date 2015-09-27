package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// VerifyLocalhost checks if request comes from localhost
func VerifyLocalhost(r string) bool {
	addr := strings.Split(r, ":") //	[]string: wyłuskuję adres IP i port, z którego przyszło żądanie
	srcIP := addr[0]
	if srcIP != "127.0.0.1" {
		return false
	}
	return true
}

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
