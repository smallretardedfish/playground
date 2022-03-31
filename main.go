package main

import (
	"fmt"
	"io"
	"reflect"
)

type Foo struct {
	a int
}

var _ MyReader = (*Foo)(nil)

func (f *Foo) Read([]byte) (int, error) {
	return 0, nil
}
func (f *Foo) Foo() error {
	return nil
}

type MyReader interface {
	Read([]byte) (int, error)
}
type AnotherReader interface {
	Read([]byte) (int, error)
	Foo() error
}

func main() {
	fmt.Println((map[int]*int64)(nil)[123])
	var f = &Foo{a: 1}
	var mr MyReader
	var ar AnotherReader
	ar = f
	mr = ar
	var empty interface{}
	var r io.Reader

	r = mr
	fmt.Println(reflect.TypeOf(r), reflect.ValueOf(r))
	fmt.Printf("%T - %v\n", empty, empty)
	fmt.Printf("%T - %v\n", r, r)
}
