package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	prevZ, z:= 0.0, 1.0
	for i := 0.0; i < 100 ; i++{
		z -= (z*z -x) /(2*z)
		fmt.Printf("The current value of z is %v\n", z)
		aZ := math.Round(z*10000000)/10000000
		aPrevZ := math.Round(prevZ*10000000)/10000000
		if  aPrevZ == aZ {
			fmt.Printf("Z has approximated to each other and the value of z is %v.\n And the value of i is %v.\n",z,i)
			return z
		} else if i == 10{
			fmt.Println("Z never approximated in ten iterations")
		}
		prevZ = z
	}
	return z
}

func main() {
	fmt.Printf("The SQRT function returned this value %v\n",Sqrt(17))
}
