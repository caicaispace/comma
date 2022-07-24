package util

import (
	"testing"

	"comma/pkg/library/util"
)

func Test_GetCurrentAbPath(t *testing.T) {
	path := util.GetCurrentAbPath()
	t.Log(path)
}

func Test_GetRootPath(t *testing.T) {
	path := util.GetRootPath()
	t.Log(path)
}

func Test_GetRootParentPath(t *testing.T) {
	path := util.GetRootParentPath()
	t.Log(path)
}

func Test_GetAppRootPath(t *testing.T) {
	path := util.GetAppRootPath()
	t.Log(path)
}

func Test_GetAppRootPath2(t *testing.T) {
	path := util.GetAppRootPath2()
	t.Log(path)
}
