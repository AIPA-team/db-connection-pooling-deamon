package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	loadConfig()

	r := mux.NewRouter()

	// start HTTP server
	fmt.Println("starting server on port:", cfg.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), r); err != nil {
		log.Fatal("error while starting http server:", err)
	}
}
