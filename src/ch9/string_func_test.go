package stringtest

import (
	"strconv"
	"testing"

	"strings"
)

func TestStringFun(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))

}

//类型转换，int 类型转字符串类型
//字符串类型转int类型
func TestConv(t *testing.T) {

	s := strconv.Itoa(10)
	t.Log("str" + s)
	if i, err:=	strconv.Atoi("10"); err== nil{
		t.Log(10 + i)
	}

}
