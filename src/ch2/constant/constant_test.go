package constant_test

import (
	"testing"
)

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
)

func TestConst(t *testing.T)  {
	t.Log(Wednesday)
}
