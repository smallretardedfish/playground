package main

import "fmt"

const (
	_ = iota
	CAT
	DOG
)

type Animal interface {
	MakeSound()
}

type Dog struct {
	phrase string
}

type Cat struct {
	phrase string
}

func (d *Dog) MakeSound() {
	fmt.Println(d.phrase)
}

func (c *Cat) MakeSound() {
	fmt.Println(c.phrase)
}

func NewCat() *Cat {
	return &Cat{phrase: "Meow"}
}

func NewDog() *Dog {
	return &Dog{phrase: "Bark"}
}

func MakePetSound(a Animal) {
	a.MakeSound()
}
func NewAnimal(animal int) Animal {
	var a Animal
	switch animal {
	case CAT:
		a = NewCat()
	case DOG:
		a = NewDog()
	}
	return a
}

func main() {

}
