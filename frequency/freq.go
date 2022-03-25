package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	lock sync.Mutex
	freq map[byte]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		lock: sync.Mutex{},
		freq: make(map[byte]int),
	}
}
func (m *SafeMap) Insert(char byte) {
	m.lock.Lock()
	if _, ok := m.freq[char]; !ok {
		m.freq[char] = 1
		//fmt.Println(char)
	}
	m.freq[char]++
	m.lock.Unlock()
}

func Frequency(text <-chan byte, dict *SafeMap) {
	fmt.Println(len(text))
	for item := range text {
		//fmt.Println(item)
		dict.Insert(item)
	}
}

func main() {
	count := 1
	text := "AAAABBCCC!!!"
	textCh := make(chan byte, len(text))
	dict := NewSafeMap()
	var wg sync.WaitGroup
	wg.Add(count)
	for w := 0; w < count; w++ {
		go func() {
			defer wg.Done()
			Frequency(textCh, dict)
		}()
	}
	for i := range []byte(text) {
		textCh <- text[i]
		//fmt.Println(rune(text[i]))
	}
	close(textCh)
	fmt.Println(dict.freq)
}
