package main

import "fmt"

func main() {
	var oddNumbers[]int = []int {1,3,5,7,11}
	fmt.Println(oddNumbers)
	oddNumbers[2] = 13
	fmt.Println(oddNumbers)

	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println(len(mySlice))

	fmt.Printf("The type of mySlice is %T",mySlice)
}