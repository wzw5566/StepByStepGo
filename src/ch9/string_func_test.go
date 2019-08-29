package stringtest

import (
	"testing"
	
	"strings"

)

func TestStringFun(t *testing.T) {
	s := "A, B ,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}

}
