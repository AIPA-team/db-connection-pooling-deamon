package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

var (
	// global variables that hold config related data
	cfg     *config
	cfgFile = "./config.yml" // default path to config file
	// global pointer to db connection
	db *sql.DB
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
