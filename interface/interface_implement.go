package main

import (
	"fmt"
)

type LivingThing interface {
	Walk(lenght int)
	Run(lenght, speed int)
}

// create struct for implement the interface
type Human struct {
	Age  int
	Name string
}

// create method for Human struct that implement interface
func (h *Human) Walk(lenght int) {
	fmt.Println(h.Name+"Walk for ", lenght, " meters")
}

func (h *Human) Run(lenght, speed int) {
	fmt.Println(h.Name+" Run for ", lenght, " with speed ", speed, " m/s. when his is", h.Age, "years old")
}

// create function for inject the interface

func TheLivingThing(l LivingThing) {
	l.Run(700, 30)
	l.Walk(900)
}

func main() {
	human := Human{
		Age:  20,
		Name: "Angga",
	}
	//  access method through implemented struct
	human.Run(200, 50)
	human.Walk(300)

	// call method that inject by interface and pass the object that implement interface as argument
	TheLivingThing(&human)

}
