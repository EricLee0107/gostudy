package main

import "fmt"

func swp(x, y int) (int, int) {
	temp := x
	x = y
	y = temp
	return x, y
}

func main() {
	x, y := 1, 2
	fmt.Printf("值传递前x为:%d,y为:%d\n", x, y)
	z, w := swp(x, y)
	// 值传递只是对x,y进行了拷贝，并没有改变x,y原有的值
	fmt.Printf("值传递后x为:%d,y为:%d\n", x, y)
	// 通过x,y的拷贝运算结果返回给了z,w 并没有改变x,y
	fmt.Println(z, w)
}
