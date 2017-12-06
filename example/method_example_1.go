package main

import "fmt"

type TZ int

func main() {
	var a TZ
	a.Print()
}
func (a *TZ) Print() {
	fmt.Println("TZ")
}
