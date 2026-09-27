package main

import "testing"

func TestVersionNotEmpty(t *testing.T) {
	if Version == "" {
		t.Fatal("Version should not be empty")
	}
}
