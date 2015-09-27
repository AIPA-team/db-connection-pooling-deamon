package main

import (
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

	req, err := http.NewRequest("POST", "", strings.NewReader(POSTBody))
	if err != nil {
		t.Error(err)
	}
	req.RemoteAddr = "127.0.0.1:8000"

	w := httptest.NewRecorder()
	query(w, req)

	if w.Code != 200 {
		t.Errorf("expected 200, got %v", w.Code)
	}

	if w.Body.String() != expected {
		t.Errorf("expected \n %v \n got \n %v", expected, w.Body.String())
	}
}
