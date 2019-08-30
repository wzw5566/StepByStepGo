package extensiontest

import (
	"fmt"
	"testing"
)

// 定义一个别名
type Code string

// 定义一个接口
type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World!\")"
}

type JavaProgrammer struct {
}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World Java\")"
}

// 定义一个方法，传入的是一个接口
func WriteFirstProgrammer(p Programmer) {
	fmt.Printf("%T %v", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	goProg := new(GoProgrammer)
	javaProg := new(JavaProgrammer)

	WriteFirstProgrammer(goProg)
	WriteFirstProgrammer(javaProg)


}
