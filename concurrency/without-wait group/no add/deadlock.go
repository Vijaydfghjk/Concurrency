package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {

		fmt.Println("Hello")
	}()
	wg.Wait()
}

/*

if we don't use wg.done() inside goroutine we get deadlock error.

*/
