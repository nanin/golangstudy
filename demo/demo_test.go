package demo

import "testing"

func TestSub(t *testing.T) {
	res := Sub(2, 1)
	if res != 1 {
		t.Fatalf("执行错误，期望值是%v,实际值是%v", 1, res)
	}
	t.Logf("执行正确，期望值是%v,实际值是%v", 1, res)
}
