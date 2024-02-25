package main

import "fmt"

type Data struct {
	name string
	age  string
}

//var data []Data

func Gets(a, b string) []Data {

	data := []Data{

		{name: a, age: b},
	}

	return data
}

func Remove(a []Data) {

	if a != nil {

		a = nil
	}
	fmt.Println("Removed :", a)
}

func main() {

	k := Gets("Vijay", "25")
	l := Gets("Ram", "35")
	m := Gets("mani", "22")

	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(m)

	Remove(m)

	v := Gets("mani", "22")
	fmt.Println(v)

}
