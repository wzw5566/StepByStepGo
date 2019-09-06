package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	//显性类型转化
	var f  = math.Sqrt(float64(x*x + y*y))
	fmt.Printf("%T\n",f) // %T输出类型
	var z  = uint(f) //显性类型转化 
	fmt.Println(x, y, z)
}
