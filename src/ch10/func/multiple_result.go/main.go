package main

import "fmt"

//swap 交换函数 返回2个字符串 
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
