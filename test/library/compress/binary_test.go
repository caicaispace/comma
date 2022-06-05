package compress_test

import (
	"strconv"
	"testing"
)

func TestBinary(t *testing.T) {
	var v int64 = 989898956898958 //默认10进制
	s2 := strconv.FormatInt(v, 2) //10 转2进制
	t.Log(s2)
}
