package main

import "fmt"

func main() {
	a()
	b()
	c()
}

func c() {
	fmt.Println("func c")
}
func b() {
	panic("panic in b")
}
func a() {
	fmt.Println("func a")
}
