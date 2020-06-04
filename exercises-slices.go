package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		a[i] = make([]uint8, dx)
	}
	for y, _ := range a {
		for x, _ := range a[y] {
			//a[y][x] = uint8((x+y)/2)
			a[y][x] = uint8(x*y)
			//a[y][x] = uint8(x^y)	
		}
	}
	return a
}

func main() {
	pic.Show(Pic)
}
