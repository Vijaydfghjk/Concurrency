package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		fmt.Println("1")
		defer wg.Done()

	}()

	go func() {

		fmt.Println("2")
		defer wg.Done()

	}()

	go func() {
		fmt.Println("3")
		defer wg.Done()

	}()

	wg.Wait() // this wait will block each one goroutine call done method.
}

/*
1. craete the waitgroup
2. you have to call method add
3. inside each goroutine weneed call method done.
4. finally need to call method wait.


*/
