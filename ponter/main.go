package main

import (
	"fmt"
)

// function to change value by memori address (pointer)
func changeValueByPointer(m *int) {
	*m = 0
}
func main() {
	x := 5
	changeValueByPointer(&x)
	fmt.Println(x)
}
