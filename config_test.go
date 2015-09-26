package main

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	cfgFile = "test.yml"
	port := 8000
	loadConfig()
	if cfg.Port != 8000 {
		t.Errorf("expected port %d, got %d", port, cfg.Port)
	}
	t.Log("listening on:", cfg.Port)
}
