package main

import (
	"testing"
)

var numOfWorkers = []int{1, 2, 3, 4, 5, 6, 7, 8}
var texts = []string{
	"uwosdlksdkljfklj",
	"уУкраїнааї",
	"ыыыыыъъъбдвлоа",
}

func BenchmarkCountFrequency(b *testing.B) {
	b.StopTimer()
	for _, item := range numOfWorkers {
		for _, text := range texts {
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				CountFrequency(text, item)
			}
			b.StopTimer()
		}
	}
}
