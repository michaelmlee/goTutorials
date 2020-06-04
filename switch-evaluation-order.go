package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now().Weekday()
	friday(today)
	fmt.Println()
	saturday(today)
	fmt.Println()
	sunday(today)
}

func friday(today time.Weekday){
	fmt.Println("When's Friday?")
	switch time.Friday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	case today + 3:
		fmt.Println("in three days.")
	case today + 4:
		fmt.Println("in four days.")
	default:
		fmt.Println("Too far away.")
	}
}

func saturday(today time.Weekday){
	fmt.Println("When's Saturday?")
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	case today + 3:
		fmt.Println("in three days.")
	case today + 4:
		fmt.Println("in four days.")
	default:
		fmt.Println("Too far away.")
	}
}

func sunday(today time.Weekday){
	fmt.Println("When's Sunday?")
	switch time.Sunday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	case today + 3:
		fmt.Println("in three days.")
	case today + 4:
		fmt.Println("in four days.")
	default:
		fmt.Println("Too far away.")
	}
}