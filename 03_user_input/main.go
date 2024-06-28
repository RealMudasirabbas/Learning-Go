package main

import (
	"bufio"
	"fmt"
	"os"
)


func main(){
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please write some number")
	input,_ := reader.ReadString('\n')
	fmt.Println("This is the result: ",input)
	fmt.Printf("This is the result and it's type is %T",input)
	
}

