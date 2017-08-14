package main

import "fmt"

func main() {
	// single array element
	var array = [...]int{1,3,4,5}

	fmt.Println("jumlah array \t", array)

	var arrays = [2][3]int{[3]int{2,4,55}, [3]int{333,3,4}}
	// in multiple array we can not include lenght of array like this:
	var arrays2 = [2][3]int{{2,4,55}, {333,3,4}}

	fmt.Println(arrays)
	fmt.Println(arrays2)

	fmt.Println()
	// let's loop the multiple dimention array using basic for
	for index := range arrays {
		for i2, v2 := range arrays[index] {
			fmt.Printf("element array %d index %d value is : %d \n", index, i2, v2)
		}
	}
}