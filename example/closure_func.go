package main

import "fmt"

// 定义一个闭包函数
func a(i int) func(x int) int {
	return func(x int) int {
		return i + x
	}
}

func main() {
	// 将a(10)赋值给num，即num := func(x int) int{return 10+x}
	num := a(10)
	//这里每次都是做10+x的计算
	fmt.Println(num(1))
	fmt.Println(num(2))
	fmt.Println(num(3))
}
