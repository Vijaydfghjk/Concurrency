package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {

		time.Sleep(time.Second)
		wg.Done()
		fmt.Println("simple")
		wg.Add(1)
	}()

	wg.Wait()
}
