package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d started \n", i)
	// some task is happening
	fmt.Printf("worker %d end \n", i)
	//wg.Done()
}

func main() {

	//fmt.Println("Explore goroutine started")
	var wg sync.WaitGroup
	//start 3 worker goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the waitgroup counter
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("workers task complete")
}
