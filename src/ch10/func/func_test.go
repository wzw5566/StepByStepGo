package functest

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// rand.Intn()返回随机数
// 首字母小写，包内方位，私有
func returnMultiValue() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

//输入是一个函数类型，返回也是函数类型
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

// 测试返回2个随机数字
func TestFunc(t *testing.T) {
	for i := 0; i < 10; i++ {
		a, b := returnMultiValue()
		t.Log(a, b)
		tsSF := timeSpent(slowFunc)
		t.Log(tsSF(10))
	}
}

//可变参数的函数
func Sum(ops ...int) int {
	sum := 0
	for _, op := range ops {
		sum += op
	}
	return sum
}

// 测试可变参数函数
func TestVarParm(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func Clear()  {
	fmt.Println("Clear resource")	
}


//延迟执行函数
func TestDefer(t *testing.T)  {
	defer Clear()
	fmt.Println("start")
	panic("err")
}