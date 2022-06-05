package util

import (
	"goaway/pkg/library/util"
	"testing"
)

func TestGetLastItem(t *testing.T) {
	slice := []int{1, 2, 3}
	t.Log(util.GetLastItem(slice))
}
