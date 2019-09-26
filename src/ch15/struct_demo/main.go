package main

import "fmt"

//Cat 定义一个猫的结构
type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
}

func main() {
	//实例化结构体
	var cat1 Cat
	cat1.Name = "little while"
	cat1.Age = 2
	cat1.Color = "白色"
	cat1.Hobby = "eat fishes"

	fmt.Println("cat1=", cat1)
	fmt.Println("name is", cat1.Name)

	//Person is a truct
	type Person struct {
		Name  string
		Age   int
		Hobby string
	}

}
