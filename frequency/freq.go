package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
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
		m.freq[char] = 0
		//fmt.Println(char)
	}
	m.freq[char]++
	m.lock.Unlock()
}

func Frequency(text chan byte, dict *SafeMap) {

	//for {
	//	char, more := <-text
	//	if !more {
	//		return
	//	}
	//	dict.Insert(char)
	//}
	for item := range text {
		//fmt.Println(item)
		dict.Insert(item)
	}

}

func CountFrequency(text string, numOfWorkers int) map[byte]int {
	count := numOfWorkers
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
	wg.Wait()
	return dict.freq
}

func main() {
	data, err := os.ReadFile("frequency/text.txt")
	if err != nil {
		log.Println(err)
		return
	}
	//buffer := make([]byte, 1024)
	//n, err := f.Read(buffer)
	//text := append(text, buffer[:n]...)
	var numOfWorkers = []int{1, 2, 3, 4, 5, 6, 7, 8, 16, 32, 100}
	estimateTime := func(f func(text string, n int), text string, n int) time.Duration {
		start := time.Now()
		f(text, n)
		res := time.Since(start)
		return res
	}
	text := string(data)
	for _, num := range numOfWorkers {
		elapsed := estimateTime(func(text string, num int) {
			CountFrequency(text, num)
		}, text, num)
		fmt.Printf("it took %v for %d workers to proccess this text.\n", elapsed, num)
	}
}
