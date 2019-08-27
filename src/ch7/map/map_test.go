package map_test

import (
	"testing"
)

//初始化Map
func TestMapInit(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))

	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))

}

//map判断key是否存在
func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1)
	m1[3] = 3
	if v, ok := m1[3]; ok {
		t.Logf("key 3 value is %d", v)
	} else {
		t.Log("key 3 is not existing")
	}
}

// 遍历map
func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}

// value 是一个方法
func TestMapWIthFuncValue(t *testing.T) {
	//定义一个value是一个方法func(op int)，返回值是int的map
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Logf("m[1] value is %d,m[2] value is%d,m[3] value is %d", m[1](2), m[2](2), m[3](2))
}
