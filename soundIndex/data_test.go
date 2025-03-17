package main

import (
	"testing"

	"github.com/xrash/smetrics"
)

func TestSoundex(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"example", "E251"},
		{"sound", "S530"},
		{"index", "I532"},
		{"", ""},
		{"a", "A000"},
	}

	for _, test := range tests {
		result := soundex(test.input)
		if result != test.expected {
			t.Errorf("soundex(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestProcessData(t *testing.T) {
	input := "example sound index"
	expected := []FormData{
		{"example", "E251"},
		{"sound", "S530"},
		{"index", "I532"},
	}

	result := processData(input)
	if len(result) != len(expected) {
		t.Fatalf("processData(%q) returned %d items; want %d", input, len(result), len(expected))
	}

	for i, data := range result {
		if data.UserInput != expected[i].UserInput || data.Result != expected[i].Result {
			t.Errorf("processData(%q)[%d] = %v; want %v", input, i, data, expected[i])
		}
	}
}

func TestConvertToTable(t *testing.T) {
	input := []FormData{
		{"example", "E251"},
		{"sound", "S530"},
		{"index", "I532"},
	}
	expected := []Table{
		{"example", "E251", smetrics.Soundex("example"), "PASS"},
		{"sound", "S530", smetrics.Soundex("sound"), "PASS"},
		{"index", "I532", smetrics.Soundex("index"), "PASS"},
	}

	result := convertToTable(input)
	if len(result) != len(expected) {
		t.Fatalf("convertToTable(%v) returned %d items; want %d", input, len(result), len(expected))
	}

	for i, table := range result {
		if table.Name != expected[i].Name || table.Soundex != expected[i].Soundex || table.Expected != expected[i].Expected || table.Test != expected[i].Test {
			t.Errorf("convertToTable(%v)[%d] = %v; want %v", input, i, table, expected[i])
		}
	}
}
