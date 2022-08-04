package segment

import (
	"strings"

	"github.com/caicaispace/gohelper/orm/gorm"
	orm "gorm.io/gorm"
)

const (
	dicWordTable = "dict_word"
	// dicWordTable = "dict_word_copy"
)

// GetLastCreateTime
func GetLastCreateTime() int {
	sqlStr := "select max(create_time) as last_create_time from dict_version"
	var lastCreateTime int
	gorm.GetInstance().GetDB("").Raw(sqlStr).Scan(&lastCreateTime)
	return lastCreateTime
}

type wordModel struct {
	Id        string `gorm:"id"`
	Word      string `gorm:"word"`
	Frequency int    `gorm:"frequency"`
	Classify  string `gorm:"classify"`
}

// LoadDictFromDB
func LoadDictFromDB() ([]*wordModel, error) {
	outData := make([]*wordModel, 0)
	// models := make([]*wordModel, 0)
	table := gorm.GetInstance().GetDB("").Table(dicWordTable)
	table.Where("is_del", 0).Find(&outData)
	// table.FindInBatches(&models, 5000, func(tx *gorm.DB, batch int) error {
	// 	outData = append(outData, models...)
	// 	return nil
	// })
	return outData, nil
}

type synonymsModel struct {
	WordIds string `gorm:"word_ids"`
	Rate    string `gorm:"rate"`
}

// LoadSynonymsDictFromDB
func LoadSynonymsDictFromDB(projectId int) (*map[string][]string, error) {
	outData := make(map[string][]string)
	models := make([]*synonymsModel, 0)
	gorm.GetInstance().GetDB("").Table("dict_synonyms").Where("is_del", 0).Where("project_id", projectId).Find(&models)
	var wordSynonymsIds []string
	var rates []string
	for _, model := range models {
		wordSynonymsIds = append(wordSynonymsIds, model.WordIds)
		rates = append(rates, model.Rate)
	}
	for i, v := range wordSynonymsIds {
		r := synonymsAgg(v, rates[i])
		if r == nil {
			continue
		}
		for k2, v2 := range *r {
			outData[k2] = v2
			// fmt.Println(k2, v2)
		}
	}
	return &outData, nil
}

type highFrequencyModel struct {
	Word string `gorm:"word"`
}

// LoadHighFrequencyDictFromDB
func LoadHighFrequencyDictFromDB(projectId int) (*map[string]bool, error) {
	models := make([]*highFrequencyModel, 0)
	table := gorm.GetInstance().GetDB("").Table("dict_high_frequency")
	table.Select(`
dict_high_frequency.id AS id,
dict_high_frequency.word_id AS word_id,
dict_word.word AS word
`)
	table.Joins("left join dict_word ON dict_high_frequency.word_id = dict_word.id")
	table.Where("project_id", projectId)
	table.Where("dict_word.is_del", 0)
	table.Where("dict_high_frequency.is_del", 0)
	table.Order("dict_high_frequency.id DESC")
	table.Find(&models)
	outData := make(map[string]bool)
	for _, model := range models {
		outData[model.Word] = true
	}
	return &outData, nil
}

type stopModel struct {
	Word string `gorm:"word"`
}

// LoadStopDictFromDB
func LoadStopDictFromDB() (*map[string]bool, error) {
	modes := make([]*stopModel, 0)
	table := gorm.GetInstance().GetDB("").Table("dict_stop")
	table.Select(`
dict_stop.id AS id,
dict_stop.word_id AS word_id,
dict_word.word AS word
`)
	table.Joins("left join dict_word ON dict_stop.word_id = dict_word.id")
	table.Where("dict_word.is_del", 0)
	table.Where("dict_stop.is_del", 0)
	table.Order("dict_stop.id DESC")
	table.Find(&modes)
	outData := make(map[string]bool)
	for _, model := range modes {
		outData[model.Word] = true
	}
	return &outData, nil
}

type bannedModel struct {
	Word string `gorm:"word"`
}

// LoadBannedDictFromDB
func LoadBannedDictFromDB() (*map[string]bool, error) {
	models := make([]*bannedModel, 0)
	table := gorm.GetInstance().GetDB("").Table("dict_banned")
	table.Select(`
dict_banned.id AS id,
dict_banned.word_id AS word_id,
dict_word.word AS word
`)
	table.Joins("left join dict_word ON dict_banned.word_id = dict_word.id")
	table.Order("dict_banned.id DESC")
	table.Find(&models)
	outData := make(map[string]bool)
	for _, model := range models {
		outData[model.Word] = true
	}
	return &outData, nil
}

// LoadBannedDictV3FromDB
func LoadBannedDictV3FromDB() (*map[string]bool, error) {
	outData := make(map[string]bool)
	models := make([]*bannedModel, 0)
	table := gorm.GetInstance().GetDB("").Table("cd_word_blacklist")
	table.Order("id DESC")
	// table.Where("id < 10000")
	table.FindInBatches(&models, 5000, func(tx *orm.DB, batch int) error {
		for _, row := range models {
			outData[row.Word] = true
		}
		return nil
	})
	return &outData, nil
}

type hyponymModel struct {
	Word    string `gorm:"word"`
	WordHyp string `gorm:"word_hyp"`
	Rate    string `gorm:"rate"`
}

// LoadHyponymDictFromDB
// 返回参数第一参数是 下位词列表
// 返回参数第二参数是 上位词列表 这个有没有必要进行上位词处理 暂不处理上位词 主要在无搜索结果情况下可以使用上位 暂时舍弃
func LoadHyponymDictFromDB() (*map[string][]string, error) {
	models := make([]*hyponymModel, 0)
	table := gorm.GetInstance().GetDB("").Table("dict_hyponym")
	table.Select(`
dict_hyponym.id AS id,
dict_word.word AS word,
hyponym_word.word AS word_hyp,
dict_hyponym.rate AS rate
`)
	table.Joins("left join dict_word ON dict_hyponym.hypernym_word_id = dict_word.id")
	table.Joins("left join dict_word as hyponym_word ON dict_hyponym.hyponym_word_id = hyponym_word.id")
	table.Order("dict_hyponym.id DESC")
	table.Find(&models)
	outData := make(map[string][]string)
	for _, model := range models {
		if v, ok := outData[model.WordHyp]; ok {
			outData[model.WordHyp] = append(v, model.Word+"|"+model.Rate)
		} else {
			outData[model.WordHyp] = []string{model.Word + "|" + model.Rate}
		}
	}
	return &outData, nil
}

type ProjectModel struct {
	ID int `gorm:"primaryKey"`
}

type ProjectOutData struct {
	List   []*ProjectModel
	Total  int64
	LastId int
}

// LoadProjectFromDB
func LoadProjectFromDB() *ProjectOutData {
	var total int64
	gorm.GetInstance().GetDB("").Table("dict_project").Select("count(*)").Count(&total)
	list := make([]*ProjectModel, 0)
	gorm.GetInstance().GetDB("").Table("dict_project").Select("id").Order("id DESC").Find(&list)
	return &ProjectOutData{
		List:   list,
		Total:  total,
		LastId: (*(list[:len(list)-1][0])).ID,
	}
}

type synonymsAggField struct {
	Word string `gorm:"word"`
}

// synonymsAgg
func synonymsAgg(wordSynonymsIds string, rate string) *map[string][]string {
	models := make([]*synonymsAggField, 0)
	table := gorm.GetInstance().GetDB("").Table("dict_word")
	table.Where("is_del", 0)
	table.Where("id", strings.Split(wordSynonymsIds, ",")).Find(&models)
	words := make([]string, 0)
	for _, model := range models {
		words = append(words, strings.Trim(model.Word, " "))
	}
	outData := make(map[string][]string)
	for i, word := range words {
		var v1 []string
		for j, v2 := range words {
			if i == j {
				continue
			}
			v1 = append(v1, v2+"|"+rate)
		}
		outData[word] = v1
	}
	return &outData
}
