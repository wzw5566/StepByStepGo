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
	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	//make 初始化切片，两个参数 len，cap 长度和容量
	// make初始化s2 ，长度为3，容量为5的int 变量,访问超过长度len的，会报数组下标越界
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))

	t.Log(s2[0], s2[1], s2[2])

}

//可变空间的切片
func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

// 切片共享内存,当其中一个值改变，与其共享内存的其他变量也会改变
func TestSliceShareMomery(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Jun", "Jul", "Aug", "Seq", "Oct", "Nov", "Dec"}

	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))

	sunmmer := year[5:8]
	t.Log(sunmmer, len(sunmmer), cap(sunmmer))

	//修改切片的值
	Q2[0] = "unknow"
	t.Log(Q2)
	t.Log(sunmmer)
	t.Log(year)

}


