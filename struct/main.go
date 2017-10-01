package main

import (
	"fmt"
	"github.com/dirham/golearn/struct/model"
)

type human struct {
	name string // define struct field name with name and type is string
	age  int    // field age with type is int
}

func main() {
	var thisStruct = human{"Dirham", 9}
	var a = model.Robot{"I'm robot", "Android"}

	fmt.Println(a)
	// fmt.Println(a.SayHalo())
	a.SayHalo("tesing")
	fmt.Println(a)

	fmt.Println(thisStruct)
}
