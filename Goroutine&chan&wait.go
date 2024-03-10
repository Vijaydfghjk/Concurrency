package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

//var mu sync.Mutex

type data struct {
	id int
}

func (p *data) call() {
	defer wg.Done()
	fmt.Printf("id is %d\n", p.id)
	//mu.Unlock()
}

type Demo struct {
	Task  []data
	value chan data
}

func (p *Demo) pro() {

	for get := range p.value {

		get.call()
	}

}

func (p *Demo) Run() {

	fmt.Println("Dummy")
	p.value = make(chan data, len(p.Task))
	for _, v := range p.Task {

		p.value <- v
	}
	close(p.value)

	p.pro()
}

func main() {

	home := make([]data, 5)
	kk := Demo{

		Task: home,
	}

	for i := 0; i < len(home); i++ {

		home[i] = data{id: i + 10}

		//kk.id = home[i].id

	}
	wg.Add(len(home))
	go kk.Run()
	//fmt.Println("struct object", kk.data)
	//fmt.Println("array of object", home[0].id)

	wg.Wait()
	fmt.Println(" Function End")
}
