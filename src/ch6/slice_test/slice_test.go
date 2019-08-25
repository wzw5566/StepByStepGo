package slice_test

import (
	"testing"
)

// len()长度， cap 容量
func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	//增加值到切片 append
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	//初始化切片
	s1 := []int{1,2,3,4}
	t.Log(len(s1), cap(s1))

	//make 初始化切片，两个参数 len，cap 长度和容量
	// make初始化s2 ，长度为3，容量为5的int 变量,访问超过长度len的，会报数组下标越界
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))

	t.Log(s2[0], s2[1], s2[2])



}
