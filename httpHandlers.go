package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Query struct holds struct data from http query
type Query struct {
	Query  string        `json:"query"`
	Params []interface{} `json:"params"`
}

// NewQuery is a constructor function for Query Struct
func NewQuery() *Query {
	q := &Query{}
	q.Params = make([]interface{}, 0)
	return q
}

type rowData map[string]interface{}

func query(w http.ResponseWriter, r *http.Request) {
	defer Recover("query request")

	if ok := checkIP(r.RemoteAddr); !ok {
		JSONerr(w, errors.New("failed remoteAddr check!"))
		return
	}

	var results []rowData

	// parsing form
	query := NewQuery()
	if err := json.NewDecoder(r.Body).Decode(query); err != nil {
		fmt.Println("newDecoder", err)
		JSONerr(w, err)
		panic(err)
	}
	defer r.Body.Close()

	rows, err := db.Queryx(query.Query, query.Params...)
	if err != nil {
		fmt.Println("queryX", err)
		JSONerr(w, err)
		panic(err)
	}
	for rows.Next() {
		m := make(rowData, 0)
		err = rows.MapScan(m)
		for k, v := range m {
			if b, ok := v.([]byte); ok {
				m[k] = string(b)
			}
		}
		if err != nil {
			JSONerr(w, err)
			rows.Close()
		}
		results = append(results, m)
	}
	if rows.Err() != nil {
		fmt.Println("ERRR na koniec rows")
		JSONerr(w, rows.Err())
		panic(err)
	}

	err = WriteJSON(w, results)
	if err != nil {
		fmt.Println("WriteTOJSON")
		JSONerr(w, err)
		panic(err)
	}
	return
}
