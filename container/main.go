package main

import (
	"fmt"
)

type Speaker interface {
	SayHello()
}

type Human struct {
	Greeting string
}

func (s Human) SayHello() {
	fmt.Println(s.Greeting)
}

func Sum(args ...int) int {
	if len(args) == 0 {
		return 0
	}
	return args[0] + Sum(args[1:]...)
}

func main() {
	ints := []int{1, 2, 3, 4, 5}
	fmt.Println(Sum(ints...))

	h := Human{Greeting: "Hello"}
	var s Speaker = h
	h.Greeting = "Meow"
	s.SayHello()
}
