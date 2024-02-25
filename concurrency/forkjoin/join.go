package main

import (
	"fmt"
	"time"
)

func main() {

	go work() // fork point
	time.Sleep(100 * time.Microsecond)

	fmt.Println("waiting done main exits")
	//join poit.
}

func work() {

	time.Sleep(500 * time.Millisecond)
	fmt.Println("printing the stuff")
}
