package emtyinterface

import (
	"fmt"
	"testing"
)

//使用if判断语句实现
func DoSomethingA(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("Interger", i)
		return
	}
	if s, ok := p.(string); ok {
		fmt.Println("string", s)
		return
	}
	fmt.Println("Unknow Type")
}

// 使用switch优化if写法
func DoSomethingB(p interface{})  {
	switch v:= p.(type) {
	case int:
		fmt.Println("Integer",v)
	case string:
		fmt.Println("String",v)
	default:
		fmt.Println("Unknow Type")

	}
	
	
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomethingA(10)
	DoSomethingA("10")
	DoSomethingB(11)
	DoSomethingB("11")
}
