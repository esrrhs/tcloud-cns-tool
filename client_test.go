package main

import "testing"

func TestNewClient(t *testing.T) {
	cli := New("id", "key")
	if cli == nil {
		t.Fatal("New returned nil")
	}
	if cli.SecretId != "id" || cli.SecretKey != "key" {
		t.Fatalf("unexpected client fields: %+v", cli)
	}
}
