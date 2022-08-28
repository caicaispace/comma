package pinyin_test

import (
	"fmt"
	"testing"

	"comma/pkg/service/admin"
	"comma/pkg/service/correction/pinyin"
)

func Test_index(t *testing.T) {
	list := loadTestData()
	py := pinyin.GetInstance()
	for _, item := range list {
		py.Add(int(item.ID), item.Keyword, item.Pinyin, item.PinyinInitials)
	}
	fmt.Println(py.FindByPinyinPrefix("kj"))
	fmt.Println(py.FindByPinyin("keji"))
}

func loadTestData() []*admin.Pinyin {
	list := make([]*admin.Pinyin, 0)
	list = append(list, &admin.Pinyin{
		ID:             0,
		Keyword:        "科技",
		Pinyin:         "keji_0",
		PinyinInitials: "kj_0",
	})
	list = append(list, &admin.Pinyin{
		ID:             1,
		Keyword:        "客机",
		Pinyin:         "keji_1",
		PinyinInitials: "kj_1",
	})
	list = append(list, &admin.Pinyin{
		ID:             2,
		Keyword:        "柯基",
		Pinyin:         "keji_2",
		PinyinInitials: "kj_2",
	})
	list = append(list, &admin.Pinyin{
		ID:             3,
		Keyword:        "可及",
		Pinyin:         "keji_3",
		PinyinInitials: "kj_3",
	})
	list = append(list, &admin.Pinyin{
		ID:             4,
		Keyword:        "科级",
		Pinyin:         "keji_4",
		PinyinInitials: "kj_4",
	})
	return list
}
