package constant_test

import (
	"testing"
)
// 内存连续常量的定义，iota + 1,意思为：连续的常量 间隔为1
const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
)

func TestConst(t *testing.T)  {
	t.Log(Monday,Wednesday)
}
