package main

import(
	"fmt"
)


func main()  {
	var i int = 20

	// m是一个指针类型，类型为*int
	//
	var m *int = &i
	//取变量的地址 &
	fmt.Println("i memory adress is", &i )
  fmt.Printf("m=%v\n",*m)
	fmt.Println("m memory adress is", &m )

}