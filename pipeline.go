package main

import "fmt"

/*

What is the Pipeline Pattern?

The pipeline pattern structures a series of processing stages, with each stage executed concurrently.
Data flows through these stages sequentially, allowing efficient data transformation and processing.

*/

func main() {

	list := []int{1, 2, 3, 4, 5}

	input := make(chan int, len(list))

	for _, v := range list {

		input <- v
	}
	close(input)

	// First stage of the pipeline:

	doubleOutput := make(chan int)

	go func() {

		for val := range input {

			doubleOutput <- val * 2
		}
		defer close(doubleOutput)
	}()

	squre := make(chan int)

	go func() {

		for get := range doubleOutput {

			squre <- get * get
		}
		defer close(squre)
	}()

	for results := range squre {

		fmt.Println(results)
	}
}
