package main

import (
	"fmt"
)

func main() {
	// init map

	// opt 1 inline : map[type of key]type of value, ex:
	var myMap = map[string]string{}
	// assign key and value to new map
	myMap["key"] = "value"
	myMap["draf"] = "golang"
	// print out map
	fmt.Println(myMap)

	// loop map
	for key, val := range myMap {
		fmt.Printf("key is : '%s' and value is : '%s' \n", key, val)
	}

	// init map opt 2 : we can assign key and value when create map just like array and slice, like this :
	var myMap1 = map[string]int{"i'm key": 9}
	// or like this
	var myMap2 = map[string]int{
		"key":  9,
		"key2": 8, // note : when assign value like this the important thing is always remeber to put (,) at and of the value assigment
		"a":    0, // another thing is when you print out map the value inside the map is not ordered
	}

	fmt.Println(myMap1)
	fmt.Println(myMap2)
}
