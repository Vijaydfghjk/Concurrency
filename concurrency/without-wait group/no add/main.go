package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	go func() {
		defer wg.Done()
		time.Sleep(300 * time.Millisecond)
		fmt.Println("a goroutine")
	}()

	wg.Wait()
	fmt.Println("exected immdiately")

}

/*
 here we are not used  wg.add() that is the reason main does not wait for the goroutine completion.

*/
