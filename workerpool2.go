package main

import (
	"fmt"
	"sync"
)

// This pattern is valuable when you have a large number of tasks to execute.

/*

  This pattern is valuable when you have a large number of tasks to execute concurrently.

  Pattern Explanation:

  i)   numWorkers := 3
       - This creates 3 worker goroutines.
       - Each goroutine calls the Worker() function.

  ii)  Each goroutine tries to read values from the job channel concurrently.

  iii) The job channel has 10 values to process.
       - These values are sent into the channel before closing it.

  iv)  The 3 goroutines (G1, G2, G3) take values from the single job channel.
       - The jobs are distributed among the goroutines by the Go runtime.

  v)   This approach is faster because:
       - Multiple goroutines process the jobs in parallel.
       - The overall processing time is reduced as the workload is shared.

  vi)  Example:
       - Let's say the jobs have IDs 1 to 10.
       - G1 might process jobs 1, 4, 7, 10.
       - G2 might process jobs 2, 5, 8.
       - G3 might process jobs 3, 6, 9.
       - The exact distribution depends on the runtime scheduling.

  vii) This pattern improves throughput and reduces latency compared to a single worker approach.
       - In the single worker approach, all jobs would be processed sequentially by one goroutine, resulting in higher total processing time.
*/

func Worker(id int, jobs <-chan int, results chan<- int) {

	for job := range jobs {

		fmt.Printf("Worker %d processing job %d\n", id, job)
		results <- job * 2
	}

}

func main() {

	var wg sync.WaitGroup

	numJobs := 10
	numWorkers := 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		//wg.Add(1)
		go func(workerid int) {
			defer wg.Done()
			Worker(workerid, jobs, results)
		}(i)
	}

	for i := 1; i <= 10; i++ {

		jobs <- i
	}
	close(jobs)
	go func() {

		wg.Wait()
		close(results)

	}()

	for v := range results {

		fmt.Printf("Result: %d\n", v)
	}

}
