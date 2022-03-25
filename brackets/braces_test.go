package main

import (
	"testing"
)

type TestIndex struct {
	input    string
	expected bool
}

var testTable = []TestIndex{
	{
		"()()()()",
		true,
	},
	{
		"{{{}",
		false,
	},
	{
		"",
		true,
	},
	{
		")",
		false,
	},
	{
		"}{}{",
		false,
	},
}

func TestIsBalanced(t *testing.T) {
	rangeTestTable(t, IsBalanced, testTable)
}
func rangeTestTable(t *testing.T, f func(string) bool, testTable []TestIndex) {
	for _, testCase := range testTable {
		want := testCase.expected
		got := f(testCase.input)
		if want != got {
			t.Errorf("input: %v, output: %v, expected: %v", testCase.input, got, want)
		}
	}
}

func BenchmarkIsBalanced(b *testing.B) {
	b.StopTimer()
	for _, tt := range testTable {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			IsBalanced(tt.input)
		}
		b.StopTimer()
	}
}
