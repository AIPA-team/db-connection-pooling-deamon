package main

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// struct holding whole config data
type config struct {
	Port int `yaml:"port"` // HTTP port to listen for requests
}

// global variables that hold config related data
var (
	cfg     *config
	cfgFile = "./config.yml" // default path to config file
)

// loadConfig loads data from yml file to config struct
func loadConfig() {

	// get config file location from flag, or use default one
	flag.StringVar(&cfgFile, "configFile", cfgFile, "http server listening port")
	flag.Parse()

	data, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		log.Fatalf("error while reading cfg file: %v", err)
	}
	yaml.Unmarshal(data, &cfg)
}
