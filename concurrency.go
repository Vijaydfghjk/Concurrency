package main

import (
	"context"
	"fmt"
	"sync"
)

func faIn(ctx context.Context, fetchers ...<-chan interface{}) <-chan interface{} { //<-chan going to read
	// ... mean going mention 3 to 4 channel
	fmt.Println("fain starts")

	//fetchers ...<-chan interface{} variable number of argument menas I can take number of channel type

	combinedfercher := make(chan interface{})

	//1

	var wg sync.WaitGroup
	wg.Add(len(fetchers)) // menas 3 chennal

	fmt.Println("3 go routinues waits")

	for _, f := range fetchers {

		fmt.Println("Go routinues rages:", f)

		f := f

		//3

		go func() {

			fmt.Println("unnamed goroutine")

			defer wg.Done()

			for {

				select {

				case res := <-f:
					fmt.Println("data copied from one channel to another ")

					combinedfercher <- res

				case m := <-ctx.Done():

					fmt.Println("context exited goroutine", m)

					return
				}

			}

		}() // registering the 3 goroutine

	}

	// 4

	// channel cleanup

	fmt.Println("2n unnamed goroutines may register")

	go func() {

		fmt.Println("waiting for the all goroutine")

		wg.Wait()

		close(combinedfercher)

		fmt.Println("waiting is over")
	}()

	fmt.Println("value is :", combinedfercher)

	return combinedfercher
}

func main() {

	fmt.Println("main starts")

	ctx := context.TODO()

	ch1 := make(chan interface{}) // interface menas type is any type
	ch2 := make(chan interface{})
	ch3 := make(chan interface{})
	ch4 := faIn(ctx, ch1, ch2, ch3)

	fmt.Println("ch1 written")

	ch1 <- "Vijay"

	fmt.Println("Ch2 written")

	ch2 <- "Vijay is Golang Developer"

	ch3 <- "He is Genious in coding"

	fmt.Println("main ends", ch4)

}

/*

1.     ... mean going mention 3 to 4 channel

2.     fetchers ...<-chan interface{} variable number of argument menas I can take number of channel type( strig or int)

3.	   combinedfercher := make(chan interface{})

4.

  wg.Add(len(fetchers)): This line is part of the use of sync.WaitGroup.
  wg.Add is used to increment the WaitGroup counter by the number of goroutines you're planning to wait for. In your code.
  which menas 3.

5.close(combinedfercher)

  This line is used to signal that no more values will be sent to the combinedfetcher channel.


*/
