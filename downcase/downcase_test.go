package main

import (
	"testing"
)

var testCases = []struct {
	input    string
	expected string
}{
	{
		input:    "",
		expected: "",
	},
	{
		input:    "Hello World",
		expected: "hello world",
	},
	{
		input:    "HeliOS",
		expected: "helios",
	},
	{
		input:    "barong is oauth server for peatio.tech stack.",
		expected: "barong is oauth server for peatio.tech stack.",
	},
	{
		input:    "tag 2.0.34",
		expected: "tag 2.0.34",
	},
	{
		input:    "CryptoCurrency Exange PEATIO",
		expected: "cryptocurrency exange peatio",
	},
}

func TestToLower(t *testing.T) {
	for _, tt := range testCases {
		actual := ToLower(tt.input)
		// We don't expect errors for any of the test cases
		if actual != tt.expected {
			t.Fatalf("Downcase(%q) was expected to return %v but returned %v.",
				tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkToLower(b *testing.B) {
	b.StopTimer()
	for _, tt := range testCases {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			ToLower(tt.input)
		}
		b.StopTimer()
	}
}
