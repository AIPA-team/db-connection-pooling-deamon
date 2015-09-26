package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

var (
	// global variables that hold config related data
	cfg     *config
	cfgFile = "./config.yml" // default path to config file
	// global pointer to db connection
	db *sqlx.DB
)

func main() {

	var err error

	loadConfig()

	//conenct to api databases
	db, err = sqlx.Open("postgres", cfg.DB.connString+" sslmode=disable")
	if err != nil {
		log.Println("error while connecting to database:", err)
		panic(err)
	}

	db.SetMaxIdleConns(cfg.PoolMin)
	db.SetMaxOpenConns(cfg.PoolMax)

	r := mux.NewRouter()
	r.HandleFunc("/query", query)

	// start HTTP server
	fmt.Println("starting server on port:", cfg.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), r); err != nil {
		log.Fatal("error while starting http server:", err)
	}
}
