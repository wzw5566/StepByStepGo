package array_test


import (
	"testing"
)

func TestArrayInit(t *testing.T)  {

	var arr [3]int 
	arr1 := [4]int{1,2,3,4}
	arr3 := [...]int {1,2,3,4,5}

	t.Log(arr[1], arr1[2])
	t.Log(arr1,arr3)
	
}

// idx 数组的下标，e 数组的值
func TestArrayTravel(t *testing.T)  {
	arr := [...]int{1,3,4,5}
	for idx, e := range arr{
		t.Log(idx,e)
	}
}

// e 为数组的值,不需要返回的时候使用 _ 代替
func TestArrayTravel2(t *testing.T)  {
	arr := [...]int{1,3,4,5}
	for _, e := range arr{
		t.Log(e)
	}
}

// 数组截取
func TestArraySection(t *testing.T)  {
	arr3 := [...]int{1,2,3,4,5}
	arr3_sec := arr3[3:]
	t.Log(arr3_sec)
}