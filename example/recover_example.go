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
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in B")
		}
	}()
	panic("panic in b")
}
func a() {
	fmt.Println("func a")
}
