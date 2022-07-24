package util

import (
	"testing"

	"comma/pkg/library/util"
)

func TestGetLastItem(t *testing.T) {
	slice := []int{1, 2, 3}
	t.Log(util.GetLastItem(slice))
}
