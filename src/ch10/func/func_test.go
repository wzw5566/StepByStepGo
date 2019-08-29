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

func slowFunc(op int) int  {
	time.Sleep(time.Second *1)
	return op*op
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
