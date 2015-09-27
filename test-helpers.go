package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func initTestENV() error {
	if cfg.DB.User == "" {
		loadTestCfg()
	}
	testDBconn()
	deinitTestENV()
	if err := createTable(); err != nil {
		return err
	}
	return nil
}

func deinitTestENV() error {
	if _, err := db.Exec("DROP table test"); err != nil {
		return err
	}
	return nil
}

func loadTestCfg() {

	var err error

	cfgFile = "./test.yml"

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

func testDBconn() {
	var err error

	db, err = sqlx.Open("postgres", cfg.DB.connString+" sslmode=disable")
	if err != nil {
		fmt.Println("error while connecting to database:", err)
		panic(err)
	}
}

func createTable() error {

	var err error

	table := `
        CREATE table test (
            id          serial,
            firstname   varchar(40),
            lastname    varchar(40),
            age         integer
        )
    `

	mockData := `
        INSERT INTO test (
            firstname,
            lastname,
            age
        ) VALUES (
            'Tomasz',
            'Pietrek',
            30
        ), (
            'Artur',
            'Gawroński',
            32
        ), (
            'Łukasz',
            'Żarnowiecki',
            23
        )
    `

	if _, err = db.Exec(table); err != nil {
		return err
	}
	if _, err = db.Exec(mockData); err != nil {
		return err
	}
	return nil
}
