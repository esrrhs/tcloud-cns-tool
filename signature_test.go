package main

import (
	"net/url"
	"testing"
)

func TestSignatureDeterministic(t *testing.T) {
	param := url.Values{}
	param.Set("Action", "DomainList")
	param.Set("SecretId", "AKIDtest")
	param.Set("Timestamp", "1700000000")
	param.Set("Nonce", "123456")

	sig1 := Signature(param, "GET", Uri, "secret")
	sig2 := Signature(param, "GET", Uri, "secret")
	if sig1 == "" || sig1 != sig2 {
		t.Fatalf("Signature should be stable, got %q and %q", sig1, sig2)
	}
}

func TestSignatureSha256DiffersFromSha1(t *testing.T) {
	param := url.Values{}
	param.Set("Action", "RecordList")
	param.Set("domain", "example.com")

	sha1Sig := Signature(param, "GET", Uri, "secret")
	sha256Sig := SignatureSha256(param, "GET", Uri, "secret")
	if sha1Sig == "" || sha256Sig == "" {
		t.Fatal("signatures should not be empty")
	}
	if sha1Sig == sha256Sig {
		t.Fatal("sha1 and sha256 signatures should differ")
	}
}

func TestSignatureReplacesUnderscoreInKeys(t *testing.T) {
	param := url.Values{}
	param.Set("record_type", "A")
	param.Set("Action", "RecordCreate")

	// Should not panic and should produce a signature even with underscore keys.
	sig := Signature(param, "GET", Uri, "secret")
	if sig == "" {
		t.Fatal("expected non-empty signature")
	}
}
