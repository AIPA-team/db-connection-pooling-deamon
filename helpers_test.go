package main

import (
	"testing"
)

func TestCheckIP(t *testing.T) {
	goodIP := "127.0.0.1:7000"
	wrongIP := "195.225.36.156:80"

	if ok := checkIP(goodIP); !ok {
		t.Error("given localhost, get false verification")
	}

	if ok := checkIP(wrongIP); ok {
		t.Errorf("given %v which is not localhost, got passed!", wrongIP)
	}

}
