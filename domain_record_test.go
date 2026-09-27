package main

import (
	"encoding/json"
	"testing"
)

func TestRecordJSONUnmarshal(t *testing.T) {
	const raw = `{
		"Id": 12,
		"Name": "www",
		"Type": "A",
		"Value": "1.2.3.4",
		"Line": "默认",
		"updated_on": "2024-01-01 00:00:00",
		"line_id": "0",
		"q_project_id": 0
	}`
	var record Record
	if err := json.Unmarshal([]byte(raw), &record); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if record.Id != 12 || record.Name != "www" || record.Value != "1.2.3.4" || record.Type != "A" {
		t.Fatalf("unexpected record: %+v", record)
	}
}

func TestDomainListResponseJSONUnmarshal(t *testing.T) {
	const raw = `{
		"code": 0,
		"message": "",
		"codeDesc": "Success",
		"data": {
			"info": {"domain_total": 1},
			"domains": [{"id": 1, "name": "example.com", "status": "enable"}]
		}
	}`
	var resp DomainListResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if resp.Code != 0 || len(resp.Data.Domains) != 1 || resp.Data.Domains[0].Name != "example.com" {
		t.Fatalf("unexpected response: %+v", resp)
	}
}
