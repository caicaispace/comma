package segment

import (
	"comma/pkg/library/db"
	"comma/pkg/library/setting"
	"comma/pkg/service/segment"
	"reflect"
	"testing"

	types "comma/pkg/library/core/t"

	"github.com/stretchr/testify/assert"
)

func init() {
	db.New(&setting.DBSetting{
		Host:     "127.0.0.1",
		Port:     "3306",
		Username: "root",
		Password: "123456",
		DbName:   "comma",
	})
}

func TestLoadProjectFromDB(t *testing.T) {
	project := segment.LoadProjectFromDB()
	tAssert := assert.New(t)
	tAssert.Equal(reflect.TypeOf(project.List), reflect.TypeOf([]*segment.ProjectModel{}))
}

func TestGetLastCreateTime(t *testing.T) {
	lastCreateTime := segment.GetLastCreateTime()
	tAssert := assert.New(t)
	tAssert.Equal(reflect.TypeOf(lastCreateTime), reflect.TypeOf(1))
}

func TestLoadDictFromDB(t *testing.T) {
	dict, err := segment.LoadDictFromDB()
	if err != nil {
		t.Error(err)
	}
	tAssert := assert.New(t)
	tAssert.Equal(reflect.TypeOf(dict), reflect.TypeOf(&types.SliceMapStrAny{}))
}

func TestLoadSynonymsDictV2FromDB(t *testing.T) {
	dict, err := segment.LoadSynonymsDictFromDB(0)
	if err != nil {
		t.Error(err)
	}
	tAssert := assert.New(t)
	tAssert.Equal(reflect.TypeOf(dict), reflect.TypeOf(&types.MapStrSliceStr{}))
}

func TestLoadHighFrequencyDictFromDB(t *testing.T) {
	dict, err := segment.LoadHighFrequencyDictFromDB(0)
	if err != nil {
		t.Error(err)
	}
	tAssert := assert.New(t)
	tAssert.Equal(reflect.TypeOf(dict), reflect.TypeOf(&types.MapStrBool{}))
}

func TestLoadStopDictFromDB(t *testing.T) {
	dict, err := segment.LoadStopDictFromDB()
	if err != nil {
		t.Error(err)
	}
	tAssert := assert.New(t)
	tAssert.Equal(reflect.TypeOf(dict), reflect.TypeOf(&types.MapStrBool{}))
}

func TestLoadBannedDictFromDB(t *testing.T) {
	dict, err := segment.LoadBannedDictFromDB()
	if err != nil {
		t.Error(err)
	}
	tAssert := assert.New(t)
	tAssert.Equal(reflect.TypeOf(dict), reflect.TypeOf(&types.MapStrBool{}))
}

func TestLoadBannedDictV3FromDB(t *testing.T) {
	dict, err := segment.LoadBannedDictV3FromDB()
	if err != nil {
		t.Error(err)
	}
	tAssert := assert.New(t)
	tAssert.Equal(reflect.TypeOf(dict), reflect.TypeOf(&types.MapStrBool{}))
}

func TestLoadHyponymDictFromDB(t *testing.T) {
	dict, err := segment.LoadHyponymDictFromDB()
	if err != nil {
		t.Error(err)
	}
	tAssert := assert.New(t)
	tAssert.Equal(reflect.TypeOf(dict), reflect.TypeOf(&types.MapStrSliceStr{}))
}
