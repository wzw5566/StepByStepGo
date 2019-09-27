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

func main() {
	 var person Person
	 person.Name = "Vincent"
	 person.Speak()
}
