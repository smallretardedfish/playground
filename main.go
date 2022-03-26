package main

import (
	"fmt"
)

const (
	_ = 1 << iota
	UA
	ENG
	US
)

func main() {
	const NIGER int = 100
	countryCodes := []int{UA, ENG, US, NIGER}
	for _, code := range countryCodes {
		switch code {
		case UA:
			fmt.Println("UKRAINE")
		case ENG:
			fmt.Println("ENGLAND")
		case US:
			fmt.Println("UNITED STATES")
		default:
			fmt.Println("WFT")
		}
	}
}
