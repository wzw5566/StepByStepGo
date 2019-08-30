package interfacetest

import "testing"

// 定义一个接口
type Programmer interface {
	WriteHelloWorld() string
}

// 定义一个结构体
type GoProgrammer struct {
}

// 实现了接口的方法，并绑定到了结构体上，方法签名必须跟接口一致
func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())

}
