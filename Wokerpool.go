package main

import (
	"fmt"
	"sync"
)

/*
Managing the large number of goroutine is challenging this is where worker pool coming.

A worker Pool is the pattern that control the number of goroutines that are working on a task at any given time.

it helps in resource management by limiting the number of goroutines
*/
type Task struct {
	ID int
}

func (t *Task) Process() {

	fmt.Printf("Processing task %d\n", t.ID)
	//time.Sleep(2 * time.Second)
}

type Workerpool struct {
	Tasks       []Task
	concurrency int
	taskschan   chan Task
	wg          sync.WaitGroup
}

func (wp *Workerpool) worker() {

	for task := range wp.taskschan {

		task.Process()
		wp.wg.Done()
	}
}

func (wp *Workerpool) Run() {

	// Initialize yhe task channel

	wp.taskschan = make(chan Task, len(wp.Tasks))
	wp.wg.Add(len(wp.Tasks))
	for i := 0; i < wp.concurrency; i++ {

		go wp.worker()
	}
	//fmt.Println("len is ", len(wp.Tasks))

	for _, task := range wp.Tasks {

		wp.taskschan <- task
	}
	close(wp.taskschan)
	wp.wg.Wait() // wait for all tasks to finish
}

func main() {

	jobs := make([]Task, 20)

	for i := 0; i < len(jobs); i++ {

		jobs[i] = Task{ID: i + 1}

	}

	// create a worker Pool

	wp := Workerpool{

		Tasks:       jobs,
		concurrency: 5,
	}

	// run the pool

	wp.Run()
	fmt.Println("All tasks have been processed")
}
