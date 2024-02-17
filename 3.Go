package main

import (
	"fmt"
	"sync"
	"time"
)


func details(age int) int{

 if age > 18{


	  fmt.Println("He is Voter")
 }else {

	fmt.Println("Not elligible")
 }

 age = 25

 return age
}

func myday(myage int) {

	var a int =  20

	v1:=details(a)
	fmt.Println(v1)
}

type Data struct{

	name string

}


func (P *Data) call() string{

    P.name = "Ram"

  if P.name == "Ram"{

	   fmt.Println("He is Good person")
  }else {

	fmt.Println("Bad guy")
  }

  return P.name
}

type masterdata interface{

   call() string
}

func School (person chan int){

   	
	    a:=<-person

   if a < 18{

         fmt.Println("Not Voter")
   }else {

	fmt.Println("Voter")
	 
   }

}
var vj Data

func main (){

 details(50)

 var person chan int =make(chan int)
 go School(person)
    person<- 50
 time.Sleep(5)
  var MaduraiBalaji sync.WaitGroup 
      MaduraiBalaji.Add(1)
	  MaduraiBalaji.Wait()

 var vj Data
 var djsathish [3]Data
     djsathish[0].call()
	 djsathish[1].call()
	 djsathish[2].call()


     vj.call()
	 var p* Data =new(Data)
         p.call()
		var callcentreDiniesh  []Data = make([]Data, 3)
		callcentreDiniesh[0].call()
		callcentreDiniesh[1].call()
		callcentreDiniesh[2].call()

     var raj  masterdata
     
     raj =&vj
	 raj.call()
	 raj=p
	 raj.call()

}