package main

import "fmt"

func main() {
	fmt.Println("starting main")
	
	//this will print last
	defer fmt.Println("defer in main 12")
	
	deferStatement()
	
	//this will be second to last
	defer fmt.Println("defer in main 11")

	fmt.Println("done in main")
}

func deferStatement(){
	fmt.Println("counting")
	for i := 10; i > 2; i-- {
		defer fmt.Println(i)
	}
	fmt.Println("for loop done")
	defer fmt.Println("defer in 2")
	defer fmt.Println("defer in 1")
	
}