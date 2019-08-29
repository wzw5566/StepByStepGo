package stringtest

import "testing"

func TestString(t *testing.T) {
	//string默认初始化
	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s))
	t.Log(s)

	//中文汉字“严”的二进制数
	s = "\xE4\xB8\xA5"
	t.Log(s)
	t.Log(len(s))

	s = "中"

	c := []rune(s)
	t.Logf("中的Unicode %x", c[0])
	t.Logf("中 UTF8 %x",s)
}


func TestStringRune(t *testing.T)  {
	s := "我爱中华"
	for _, c := range s {
		t.Logf("%[1]c %[1]d",c)
	}
}