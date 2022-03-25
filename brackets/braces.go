package main

import (
	"Playground/stack"
	"bufio"
	"fmt"
	"log"
	"os"
)

func IsBalanced(str string) bool {
	stack := stack.NewStack[string]()
	for _, item := range []byte(str) {
		val := string(item)
		var expected string
		if val == "(" || val == "{" || val == "[" {
			stack.Push(val)
			continue
		}
		switch val {
		case ")":
			expected = "("
		case "]":
			expected = "["
		case "}":
			expected = "{"
		default:
			continue
		}
		if stack.Pop() != expected {
			return false
		}
	}
	return stack.IsEmpty()
}

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		buffer := make([]byte, 1024)
		_, err := reader.Read(buffer)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(IsBalanced(string(buffer)))
	}
}
