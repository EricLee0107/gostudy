package main

import "fmt"

func swp(x, y int) (int, int) {
	return y, x
}

func main() {
	a,b := 1,2
	c, d := swp(a, b)
	fmt.Println(c, d)
}



