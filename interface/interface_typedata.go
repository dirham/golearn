package main

import (
	"fmt"
)

func main() {
	var fruits = []interface{}{
		map[string]interface{}{"name": "strawberry", "total": 10}, // map[string]interface{}
		[]string{"manggo", "pineapple", "papaya"},                 // slice of string
		"orange", // string
		true,     // boolean
	}
	for _, each := range fruits {
		fmt.Printf("type of value : %T\n", each)
	}
}
