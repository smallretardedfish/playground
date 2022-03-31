package main

import (
	"fmt"
	"time"
)

func BubbleSort(arr []int) {
	for i := range arr {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func SelectionSort(arr []int) {
	minIdx := 0
	for i := range arr {
		for j := i; j < len(arr); j++ {
			if arr[minIdx] > arr[j] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func DoJob() {
	time.Sleep(2 * time.Second)
}

func main() {
	// implemented concurrency pattern introduced on this vid
	//https://www.youtube.com/watch?v=M-ofC23gSZ4
	count := 5
	jobIDs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	queue := make(chan struct{}, count)

	for _, id := range jobIDs {
		queue <- struct{}{}
		go func(ID int) {
			DoJob()
			fmt.Println(ID)
			<-queue
		}(id)
	}
	for i := 0; i < count; i++ {
		queue <- struct{}{}
	}
}
