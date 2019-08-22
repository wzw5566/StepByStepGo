package loop_test

import (
	"testing"
)
//表示while < 5
func TestWhileLoop(t *testing.T)  {
	for n := 0; n < 5; n++ {
		t.Log(n)
	}
}