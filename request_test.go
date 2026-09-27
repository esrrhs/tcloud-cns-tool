package main

import (
	"strings"
	"testing"
)

func TestBaseResponseErrorNilWhenOK(t *testing.T) {
	resp := BaseResponse{Code: 0, Message: "ok", CodeDesc: "Success"}
	if err := resp.Error(); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}

func TestBaseResponseErrorWhenFailed(t *testing.T) {
	resp := BaseResponse{Code: 4000, Message: "bad request", CodeDesc: "InvalidParameter"}
	err := resp.Error()
	if err == nil {
		t.Fatal("expected error")
	}
	msg := err.Error()
	if !strings.Contains(msg, "4000") || !strings.Contains(msg, "InvalidParameter") || !strings.Contains(msg, "bad request") {
		t.Fatalf("unexpected error message: %s", msg)
	}
}
