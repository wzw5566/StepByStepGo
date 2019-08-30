package encap

import (
	"fmt"
	"testing"
	"unsafe"
)

//定义一个结构体
type Employee struct {
	ID   string
	Name string
	Age  int
}

func (e *Employee) string() string {

	fmt.Printf("Andress is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID is %s/Name is %s,Age is %d",e.ID,e.Name,e.Age)
}

//初始化实例
func TestCreateEmployee(t *testing.T) {
	e := Employee{"0", "Vincent", 18}
	e1 := Employee{Name: "Mike", Age: 18}
	e2 := new(Employee) //定义一个指针
	e2.ID = "2"
	e2.Age = 22
	e2.Name = "Vincent2"

	t.Log(e)
	t.Log(e1)
	t.Log(e1.ID)
	t.Log(e2)
	// %T类型 输出
	t.Logf("e is %T", e)
	t.Logf("e2 is %T", e2)
}


func TestStructOperations(t *testing.T)  {
	e := Employee{"0","Vincent",18}
	fmt.Printf("Andress is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.string())
}