package main

import "fmt"

//Person defined struct Person
type Person struct {
	Name string
	Age  int
}

// Speak is a Person struct function
func (p Person) Speak() {
	fmt.Println("hello", p.Name)
}

func (p Person) AddNum(n1 int, n2 int) int {

	sum := n1 + n2
	return sum
}

func main() {
	var person Person
	person.Name = "Vincent"
	person.Speak() // use struct function
	n1 := 100
	n2 := 200
	fmt.Println(person.AddNum(n1, n2))
}
