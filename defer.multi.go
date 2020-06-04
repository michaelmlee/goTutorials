package main

import "fmt"

func main() {
	fmt.Println("starting main")
	//this will be last
	//into the stack first
	defer fmt.Println("defer in main 0")
	
	//this will run first
	deferStatement()
	
	//this will be second to last
	//into the stack second
	defer fmt.Println("defer in main 1")

	fmt.Println("done in main")
}

func deferStatement(){
	
	fmt.Println("counting")
	for i := 2; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("for loop done")
	
}