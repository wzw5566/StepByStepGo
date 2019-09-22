package main

import (
	"fmt"
)

//Calculation function
//input two parameters float64 m1,float64 m2, byte operator
//return float64 type result
func Calculation(m1 float64, m2 float64, operator byte) float64 {
	var result float64
	switch operator {
	case '+':
		result = m1 + m2
	case '-':
		result = m1 - m2
	case '*':
		result = m1 * m2
	case '/':
		result = m1 / m2
	default:
		fmt.Println("operator undefined")
	}
	return result
}

func main() {
	m1, m2 := 2.5, 3.6
	var operator byte = '+'
	result := Calculation(m1, m2, operator) //call the function
	fmt.Println("result", result)
}
