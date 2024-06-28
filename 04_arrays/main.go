package main

import "fmt"

func main() {
	var evenNumbers[5]int
	evenNumbers[0] = 2
	evenNumbers[1] = 4
	evenNumbers[2] = 6
	evenNumbers[3] = 8
	evenNumbers[4] = 10

	myArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println(len(myArray))

	fmt.Printf("The type of myArray is %T",myArray)
}