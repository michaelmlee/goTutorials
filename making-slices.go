package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a) //len=5 cap=5 [0 0 0 0 0]

	b := make([]int, 0, 5)
	printSlice("b", b) //len=0 cap=5 []

	c := b[:2]
	printSlice("c", c) //len=2 cap=5 [0 0]

	d := c[2:5]
	printSlice("d", d) //len=3 cap=3 [0 0 0]
	
	e := b[:5]
	e[1] = 1
	printSlice("e",e) //len=5 cap=5 [0 1 0 0 0]
	printSlice("c",c) //len=2 cap=5 [0 1]

	e[4] = 5
	printSlice("e",e) //len=5 cap=5 [0 1 0 0 5]
	
	c = c[:5]
	printSlice("c",c) //len=5 cap=5 [0 1 0 0 5]
	
	c[2] = 2
	printSlice("e",e) //len=5 cap=5 [0 1 2 0 5]
	
	printSlice("b",b) //len=0 cap=5 []
	
	//b[0] = 1 -- this throws an error
	//panic: runtime error: index out of range

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
