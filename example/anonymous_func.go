package main

import "fmt"

func main() {
	// 这里定义了一个匿名函数，且赋值给a
	a := func(i int, j int) (int, int) {
		return j, i
	}
	fmt.Println(a(3,4))

}
