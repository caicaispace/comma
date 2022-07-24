package segment_test

import (
	"fmt"
	"regexp"
	"testing"

	"comma/pkg/library/db"
	"comma/pkg/library/setting"
	"comma/pkg/service/segment"
)

func Test_Segmenter(t *testing.T) {
	db.New(&setting.DBSetting{
		Host:     "127.0.0.1",
		Port:     "3306",
		Username: "root",
		Password: "123456",
		DbName:   "comma",
	})
	// sego.LoadDb("s913_yanfa:Xuxinxin_s913@tcp(127.0.0.1:3312)/db_699pic_words?charset=utf8")

	var seg segment.SegmenterService
	seg.LoadDataToT2cDict()
	seg.LoadDict()

	word := "春天"
	word = `
当时 一系列接连的失败证明，以城市为中心的革命道路，在中国根本走不通。
革命之路，何去何从？放弃攻打城市，转战井冈山！毛泽东踏出的这一步，成了中国革命崭新的起点。
在井冈山，中国共产党人立足于中国革命现实，把马克思主义普遍真理同中国革命具体实践紧密结合，探索出“农村包围城市、武装夺取政权”的崭新道路。
实事求是、敢闯新路，是井冈山精神的核心。革命如此，建设和改革也如此，都必须从实际出发，敢于开辟前人没有走过的路。
2017年2月，井冈山市正式宣布在全国率先脱贫摘帽，成为我国贫困退出机制建立后首个脱贫摘帽的贫困县，井冈山的历史开启新篇章。
初心如磐，砥砺奋进，100年来，中国共产党带领亿万人民，在险滩激流中开辟出崭新航程，走出了一条中国特色社会主义的康庄大道。
习近平总书记指出：“中国特色社会主义道路、理论、制度、文化不断发展，拓展了发展中国家走向现代化的途径，给世界上那些既希望加快发展又希望保持自身独立性的国家和民族提供了全新选择，为解决人类问题贡献了中国智慧和中国方案。”
站在2021年瞻望未来，迈向中华民族伟大复兴的道路更加清晰、步伐更加坚定——
“中华民族迎来了从站起来、富起来到强起来的伟大飞跃，实现中华民族伟大复兴进入了不可逆转的历史进程！”
`
	re := regexp.MustCompile(`[\p{P}\p{S}\n\d]`) // 去除换行符以及所有符号
	// re := regexp.MustCompile(`([#@])|[\p{P}\p{S}\d]`) // 替换除＃和@以外的所有符号
	// re := regexp.MustCompile(`[\p{P}\p{S}\s\n\d]`) // 去除空格、去除换行符以及所有符号
	word = re.ReplaceAllString(word, "$1")
	fmt.Println(word)
	// fmt.Println(strings.ToLower(word))

	arr := seg.SegmentSearchMode(word)
	fmt.Println("--------------")
	for _, val := range arr {
		fmt.Print(val + "|")
	}
	fmt.Println("")

	fmt.Println("--------------")
	m1, m2, m3, s4 := seg.SegmentIndexMode(word, false, 0, false)

	fmt.Println(*m1)
	fmt.Println(*m2)
	fmt.Println(*m3)
	fmt.Println(*s4)

	// fmt.Println(seg.GetFre("深海"))
}
