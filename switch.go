package main

import (
	"fmt"
)

func findOs(os string){
	fmt.Print("Go runs on ")
	switch os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func main() {
	findOs("darwin")
	findOs("linux")
}