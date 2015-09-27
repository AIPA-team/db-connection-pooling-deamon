package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestQuery(t *testing.T) {

	var err error

	if err = initTestENV(); err != nil {
		t.Error(err)
		return
	}
	defer deinitTestENV()

	POSTBody := `{"query": "SELECT * from test where age > $1", "params": ["20"]}`
	expected := `[{"age":30,"firstname":"Tomasz","id":1,"lastname":"Pietrek"},{"age":32,"firstname":"Artur","id":2,"lastname":"Gawroński"},{"age":23,"firstname":"Łukasz","id":3,"lastname":"Żarnowiecki"}]`

	db, err = sqlx.Open("postgres", cfg.DB.connString+" sslmode=disable")
	if err != nil {
		fmt.Println("error while connecting to database:", err)
		panic(err)
	}

	req, err := http.NewRequest("POST", "/query", strings.NewReader(POSTBody))
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	query(w, req)

	if w.Body.String() != expected {
		t.Errorf("expected \n %v \n got \n %v", expected, w.Body.String())
	}
}
