package main

import (
	"testing"
)

func TestVerifyLocalhost(t *testing.T) {
	goodIP := "127.0.0.1:7000"
	wrongIP := "195.225.36.156:80"

	if ok := VerifyLocalhost(goodIP); !ok {
		t.Error("given localhost, get false verification")
	}

	if ok := VerifyLocalhost(wrongIP); ok {
		t.Errorf("given %v which is not localhost, got passed!", wrongIP)
	}

}
