package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	var wg sync.WaitGroup
	fmt.Println("number of cores", runtime.NumCPU())
	wg.Add(10)
	for i := 0; i < 10; i++ {

		go work(&wg, i+1)
		//time.Sleep(time.Second)
	}
	wg.Wait()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("elapsed:", time.Since(now))
	fmt.Println("main is done")
}

func work(wg *sync.WaitGroup, id int) {

	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task", id, "is done")
}

/*

func main() {
	fmt.Println("number of cores", runtime.NumCPU())
	for i := 0; i < 10; i++ {

		go work(i + 1)
		//time.Sleep(time.Second)
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println("main is done")
}

func work(id int) {

	time.Sleep(100 * time.Millisecond)
	fmt.Println("task", id, "is done")
}
*/
