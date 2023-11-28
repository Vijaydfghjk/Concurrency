package main

import "fmt"

func slicetochannel(nums []int) <-chan int {

	out := make(chan int) // unbuffered channel we can't send more than one value to this unbuffered channel.
	go func() {
		for _, n := range nums {

			out <- n
		}

		close(out)

	}()
	return out
}

func sq(in <-chan int) <-chan int {

	out := make(chan int)

	go func() {

		for n := range in {

			out <- n * n
		}
		close(out)

	}()

	return out
}

func main() {

	var a []int = []int{1, 4, 5, 6, 10}
	b := slicetochannel(a)

	d := sq(b)

	for n := range d {

		fmt.Println(n)
	}

	//fmt.Println(d)
	//fmt.Println(b)
}

/*
slicetochannel() and sq () running same time.

1. when slicetochannel() try to read first element 1 from n it iwll block untill  receive from in sq function
2. then second irate 4 , out channel wait for receive 4 once received it will block untill in sq function out chennel getting the data.
3. thi is sending one element only because of unbuffered channel.
4. close () statement is used to indicate that no more elements will be sent to the out channel.

*/
