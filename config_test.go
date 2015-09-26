package main

import (
	"fmt"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	cfgFile = "test.yml"
	port := 8000
	password := "test"
	host := "localhost"
	poolMax := 20
	poolMin := 2
	dbName := "test"
	user := "test"
	loadConfig()
	if cfg.Port != port {
		t.Errorf("expected port %d, got %d", port, cfg.Port)
	}
	if cfg.DB.Password != password {
		t.Errorf("expected password %v, got %v", password, cfg.Password)
	}
	if cfg.DB.Host != host {
		t.Errorf("expected host %v, got %v", host, cfg.Host)
	}
	if cfg.DB.DBname != dbName {
		t.Errorf("expected dbname %v, got %v", dbName, cfg.DBname)
	}
	if cfg.DB.User != user {
		t.Errorf("expected user %v, got %v", user, cfg.User)
	}
	if cfg.DB.PoolMax != poolMax {
		t.Errorf("expected poolMax %v, got %v", poolMax, cfg.DB.PoolMax)
	}
	if cfg.DB.PoolMin != poolMin {
		t.Errorf("expected poolMin %v, got %v", poolMin, cfg.DB.PoolMin)
	}
}

func TestDB_connString(t *testing.T) {
	db := DB{DBname: "test", Password: "test", Host: "localhost", User: "test"}
	connString := fmt.Sprintf("host=%s password=%s dbname=%s user=%s", db.Host, db.Password, db.DBname, db.User)
	db.setConnString()
	if db.connString != connString {
		t.Errorf("expected connString o be  %v, got %v", connString, cfg.connString)
	}
}
