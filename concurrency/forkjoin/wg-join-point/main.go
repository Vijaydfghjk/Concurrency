package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	var wg sync.WaitGroup

	wg.Add(1) // basically execute the one goroutine.

	go func() {
		defer wg.Done()
		work()
	}()

	go work() // fork point

	wg.Wait() // join to the main.
	fmt.Println("elapsed :", time.Since(now))
	fmt.Println("waiting done main exits")

}

func work() {

	time.Sleep(500 * time.Millisecond)
	fmt.Println("printing the stuff")
}
