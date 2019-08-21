package array_test


import (
	"testing"
)

func TestArrayInit(t *testing.T)  {

	var arr [3]int 
	arr1 := [4]int{1,2,3,4}
	arr4 := [...]int {1,2,3,4,5}

	t.Log(arr[1], arr[2])
	t.Log(arr1,arr4)
	
}