package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello, world!")
	time.Sleep(1000 * time.Millisecond) // simulating some work
	fmt.Println("sayHello fucntion ended successfully")
}

func sayHi() {
	fmt.Println("hii pawan!")
}

func main() {
	fmt.Println("gorouties")
	go sayHello()
	go sayHi()

	// wait for a moment to allow the goroutine to finish
	time.Sleep(3000 * time.Millisecond)

}
