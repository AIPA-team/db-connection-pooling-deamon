package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// struct holding whole config data
type config struct {
	Port int `yaml:"port"` // HTTP port to listen for requests
	DB   `yaml:"db"`
}

// DB holds data for connection to db
type DB struct {
	DBname     string `yaml:"dbName"`
	Password   string `yaml:"password"`
	Host       string `yaml:"host"`
	User       string `yaml:"user"`
	PoolMin    int    `yaml:"poolMin"`
	PoolMax    int    `yaml:"poolMax"`
	connString string
}

// methos that creates database string based on struct data
func (db *DB) setConnString() {
	db.connString = fmt.Sprintf("host=%s password=%s dbname=%s user=%s", db.Host, db.Password, db.DBname, db.User)
}

// loadConfig loads data from yml file to config struct
func loadConfig() {

	var err error

	// get config file location from flag, or use default one
	flag.StringVar(&cfgFile, "configFile", cfgFile, "http server listening port")
	flag.Parse()

	data, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		log.Fatalf("error while reading cfg file: %v", err)
	}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatal("error while unmarshalling config file:", err)
	}

	cfg.DB.setConnString()

}
