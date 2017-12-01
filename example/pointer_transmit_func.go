package main

import "fmt"

func swp(x, y *int) (*int, *int) {
	temp := *x
	*x = *y
	*y = temp
	return x, y
}

func main() {
	x, y := 1, 2
	fmt.Printf("指针传递前x为:%d,y为:%d\n", x, y)
	z, w := swp(&x, &y)
	// 指针传递实际上在交换时交换了指针所指向地址的值，所以x,y发生了变化
	fmt.Printf("指针传递后x为:%d,y为:%d\n", x, y)
	fmt.Println(*z, *w)
}
