package main

import "fmt"
import "os"

func main() {
	f := createFile("tmp/defer.txt")
	// 在打开文件的位置调用defer函数，保证文件一定会被关闭
	defer closeFile(f)
	writeFile(f)
}
func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	// 这里用到了panic函数，如果在打开文件时发生错误就会中断函数
	if err != nil {
		panic(err)
	}
	return f
}
func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}
func closeFile(f *os.File) {
	fmt.Println("closeing")
	f.Close()
}
