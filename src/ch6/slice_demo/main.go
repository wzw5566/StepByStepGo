package main

import (
	"fmt"
)

var Myslice []float64


func main()  {
	// //声明一个切片
	// var intArr [5]int = [...]int{1,33,55,34,7}
	// // slice := intArr[]
var Myslice [3]float64
Myslice[1] = 10
Myslice[2] = 20
fmt.Println(Myslice)

// 切片必须是make分配内存才可以使用
var demoSlice []float64 = make([]float64, 5, 10)
demoSlice[1] = 1
fmt.Println(demoSlice)

}