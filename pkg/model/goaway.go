package model

// DictBanned 违禁词
type DictBanned struct {
	ID         uint32 `gorm:"primaryKey;column:id;type:int(11) unsigned;not null" json:"id"`
	WordID     uint32 `gorm:"uniqueIndex:word_id_project_id;column:word_id;type:int(11) unsigned;not null;default:0" json:"wordId"`
	ProjectID  uint32 `gorm:"uniqueIndex:word_id_project_id;column:project_id;type:int(11) unsigned;not null;default:0" json:"projectId"`
	IsDel      uint8  `gorm:"column:is_del;type:tinyint(4) unsigned;not null;default:0" json:"isDel"`
	CreateTime uint32 `gorm:"column:create_time;type:int(11) unsigned;not null;default:0" json:"createTime"`
	UpdateTime uint32 `gorm:"column:update_time;type:int(11) unsigned;not null;default:0" json:"updateTime"`
}

// DictBannedColumns get sql column name.获取数据库列名
var DictBannedColumns = struct {
	ID         string
	WordID     string
	ProjectID  string
	IsDel      string
	CreateTime string
	UpdateTime string
}{
	ID:         "id",
	WordID:     "word_id",
	ProjectID:  "project_id",
	IsDel:      "is_del",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// DictFestival 节日
type DictFestival struct {
	ID         uint32 `gorm:"primaryKey;column:id;type:int(11) unsigned;not null" json:"id"`
	WordID     uint32 `gorm:"column:word_id;type:int(11) unsigned;not null;default:0" json:"wordId"`
	ProjectID  uint32 `gorm:"column:project_id;type:int(11) unsigned;not null;default:0" json:"projectId"`
	Name       string `gorm:"column:name;type:varchar(20);not null;default:''" json:"name"`            // 名称
	SunDate    string `gorm:"column:sun_date;type:varchar(10);not null;default:''" json:"sunDate"`     // 阳历
	LunarDate  string `gorm:"column:lunar_date;type:varchar(10);not null;default:''" json:"lunarDate"` // 阴历
	IsDel      bool   `gorm:"column:is_del;type:tinyint(1) unsigned;not null;default:0" json:"isDel"`
	CreateTime uint32 `gorm:"column:create_time;type:int(11) unsigned;not null;default:0" json:"createTime"`
	UpdateTime uint32 `gorm:"column:update_time;type:int(11) unsigned;not null;default:0" json:"updateTime"`
}

// DictFestivalColumns get sql column name.获取数据库列名
var DictFestivalColumns = struct {
	ID         string
	WordID     string
	ProjectID  string
	Name       string
	SunDate    string
	LunarDate  string
	IsDel      string
	CreateTime string
	UpdateTime string
}{
	ID:         "id",
	WordID:     "word_id",
	ProjectID:  "project_id",
	Name:       "name",
	SunDate:    "sun_date",
	LunarDate:  "lunar_date",
	IsDel:      "is_del",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// DictHighFrequency 高频词
type DictHighFrequency struct {
	ID         uint32 `gorm:"primaryKey;column:id;type:int(11) unsigned;not null" json:"id"`
	WordID     uint32 `gorm:"column:word_id;type:int(11) unsigned;not null;default:0" json:"wordId"`
	ProjectID  uint32 `gorm:"column:project_id;type:int(11) unsigned;not null;default:0" json:"projectId"`
	IsDel      bool   `gorm:"column:is_del;type:tinyint(1) unsigned;not null;default:0" json:"isDel"`
	CreateTime uint32 `gorm:"column:create_time;type:int(11) unsigned;not null;default:0" json:"createTime"`
	UpdateTime uint32 `gorm:"column:update_time;type:int(11) unsigned;not null;default:0" json:"updateTime"`
}

// DictHighFrequencyColumns get sql column name.获取数据库列名
var DictHighFrequencyColumns = struct {
	ID         string
	WordID     string
	ProjectID  string
	IsDel      string
	CreateTime string
	UpdateTime string
}{
	ID:         "id",
	WordID:     "word_id",
	ProjectID:  "project_id",
	IsDel:      "is_del",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// DictHyponym 上位词，下位词
type DictHyponym struct {
	ID             uint32  `gorm:"primaryKey;column:id;type:int(11) unsigned;not null" json:"id"`
	HypernymWordID uint32  `gorm:"column:hypernym_word_id;type:int(11) unsigned;not null;default:0" json:"hypernymWordId"` // 上位词id
	HyponymWordID  uint32  `gorm:"column:hyponym_word_id;type:int(11) unsigned;not null;default:0" json:"hyponymWordId"`   // 下位词id
	Rate           float32 `gorm:"column:rate;type:float;not null;default:0" json:"rate"`                                  // 下位词所占比例
	ProjectID      uint32  `gorm:"column:project_id;type:int(11) unsigned;not null;default:0" json:"projectId"`
	IsDel          bool    `gorm:"column:is_del;type:tinyint(1) unsigned;not null;default:0" json:"isDel"`
	CreateTime     uint32  `gorm:"column:create_time;type:int(11) unsigned;not null;default:0" json:"createTime"`
	UpdateTime     uint32  `gorm:"column:update_time;type:int(11) unsigned;not null;default:0" json:"updateTime"`
}

// DictHyponymColumns get sql column name.获取数据库列名
var DictHyponymColumns = struct {
	ID             string
	HypernymWordID string
	HyponymWordID  string
	Rate           string
	ProjectID      string
	IsDel          string
	CreateTime     string
	UpdateTime     string
}{
	ID:             "id",
	HypernymWordID: "hypernym_word_id",
	HyponymWordID:  "hyponym_word_id",
	Rate:           "rate",
	ProjectID:      "project_id",
	IsDel:          "is_del",
	CreateTime:     "create_time",
	UpdateTime:     "update_time",
}

// DictIncrID [...]
type DictIncrID struct {
	ID uint32 `gorm:"primaryKey;column:id;type:int(10) unsigned;not null" json:"id"`
}

// DictIncrIDColumns get sql column name.获取数据库列名
var DictIncrIDColumns = struct {
	ID string
}{
	ID: "id",
}

// DictPinyin 拼音
type DictPinyin struct {
	ID             uint32 `gorm:"primaryKey;column:id;type:int(11) unsigned;not null" json:"id"`
	Keyword        string `gorm:"unique;column:keyword;type:varchar(20);not null;default:''" json:"keyword"`         // 关键词文本
	Pinyin         string `gorm:"column:pinyin;type:varchar(100);not null;default:''" json:"pinyin"`                 // 拼音
	Pinyins        string `gorm:"column:pinyins;type:varchar(120);not null;default:''" json:"pinyins"`               // 拼音
	PinyinInitials string `gorm:"column:pinyin_initials;type:varchar(50);not null;default:''" json:"pinyinInitials"` // 拼音首字母
	English        string `gorm:"column:english;type:char(50);not null;default:''" json:"english"`                   // 英文
	UseNum         uint32 `gorm:"column:use_num;type:int(11) unsigned;not null;default:0" json:"useNum"`
	HomophonyNum   uint32 `gorm:"column:homophony_num;type:int(11) unsigned;not null;default:1" json:"homophonyNum"` // 同音次数
	IsDel          uint8  `gorm:"column:is_del;type:tinyint(4) unsigned;not null;default:0" json:"isDel"`
	CreateTime     uint32 `gorm:"column:create_time;type:int(11) unsigned;not null;default:0" json:"createTime"`
	UpdateTime     uint32 `gorm:"column:update_time;type:int(11) unsigned;not null;default:0" json:"updateTime"`
}

// DictPinyinColumns get sql column name.获取数据库列名
var DictPinyinColumns = struct {
	ID             string
	Keyword        string
	Pinyin         string
	Pinyins        string
	PinyinInitials string
	English        string
	UseNum         string
	HomophonyNum   string
	IsDel          string
	CreateTime     string
	UpdateTime     string
}{
	ID:             "id",
	Keyword:        "keyword",
	Pinyin:         "pinyin",
	Pinyins:        "pinyins",
	PinyinInitials: "pinyin_initials",
	English:        "english",
	UseNum:         "use_num",
	HomophonyNum:   "homophony_num",
	IsDel:          "is_del",
	CreateTime:     "create_time",
	UpdateTime:     "update_time",
}

// DictProject [...]
type DictProject struct {
	ID         uint32 `gorm:"primaryKey;column:id;type:int(11) unsigned;not null" json:"id"`
	Name       string `gorm:"column:name;type:varchar(200);not null;default:''" json:"name"`
	IsDel      bool   `gorm:"column:is_del;type:tinyint(1) unsigned;not null;default:0" json:"isDel"`
	CreateTime uint32 `gorm:"column:create_time;type:int(11) unsigned;not null;default:0" json:"createTime"`
	UpdateTime uint32 `gorm:"column:update_time;type:int(11) unsigned;not null;default:0" json:"updateTime"`
}

// DictProjectColumns get sql column name.获取数据库列名
var DictProjectColumns = struct {
	ID         string
	Name       string
	IsDel      string
	CreateTime string
	UpdateTime string
}{
	ID:         "id",
	Name:       "name",
	IsDel:      "is_del",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// DictRedirect [...]
type DictRedirect struct {
	ID         uint32 `gorm:"primaryKey;column:id;type:int(10) unsigned;not null" json:"id"`
	Word       string `gorm:"column:word;type:varchar(10);not null;default:''" json:"word"`
	Redirect   string `gorm:"column:redirect;type:varchar(10);not null;default:''" json:"redirect"`
	CreateTime uint32 `gorm:"column:create_time;type:int(11) unsigned;not null;default:0" json:"createTime"`
	UpdateTime uint32 `gorm:"column:update_time;type:int(11) unsigned;not null;default:0" json:"updateTime"`
}

// DictRedirectColumns get sql column name.获取数据库列名
var DictRedirectColumns = struct {
	ID         string
	Word       string
	Redirect   string
	CreateTime string
	UpdateTime string
}{
	ID:         "id",
	Word:       "word",
	Redirect:   "redirect",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// DictStop 停词
type DictStop struct {
	ID         uint32 `gorm:"primaryKey;column:id;type:int(11) unsigned;not null" json:"id"`
	WordID     uint32 `gorm:"uniqueIndex:word_id_project_id;column:word_id;type:int(11) unsigned;not null;default:0" json:"wordId"`
	ProjectID  uint32 `gorm:"uniqueIndex:word_id_project_id;column:project_id;type:int(11) unsigned;not null;default:0" json:"projectId"`
	IsDel      bool   `gorm:"column:is_del;type:tinyint(1) unsigned;not null;default:0" json:"isDel"`
	CreateTime uint32 `gorm:"column:create_time;type:int(11) unsigned;not null;default:0" json:"createTime"`
	UpdateTime uint32 `gorm:"column:update_time;type:int(11) unsigned;not null;default:0" json:"updateTime"`
}

// DictStopColumns get sql column name.获取数据库列名
var DictStopColumns = struct {
	ID         string
	WordID     string
	ProjectID  string
	IsDel      string
	CreateTime string
	UpdateTime string
}{
	ID:         "id",
	WordID:     "word_id",
	ProjectID:  "project_id",
	IsDel:      "is_del",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// DictSynonyms 同义词
type DictSynonyms struct {
	ID         uint32  `gorm:"primaryKey;column:id;type:int(11) unsigned;not null" json:"id"`
	WordIDs    string  `gorm:"column:word_ids;type:varchar(255);not null;default:''" json:"wordIds"`
	Rate       float32 `gorm:"column:rate;type:float;not null;default:1" json:"rate"`                       // 比例
	ProjectID  uint32  `gorm:"column:project_id;type:int(11) unsigned;not null;default:0" json:"projectId"` // 项目组id默认0
	IsDel      bool    `gorm:"index:WORD;column:is_del;type:tinyint(1) unsigned;not null;default:0" json:"isDel"`
	CreateTime uint32  `gorm:"column:create_time;type:int(11) unsigned;not null;default:0" json:"createTime"`
	UpdateTime uint32  `gorm:"column:update_time;type:int(11) unsigned;not null;default:0" json:"updateTime"`
}

// DictSynonymsColumns get sql column name.获取数据库列名
var DictSynonymsColumns = struct {
	ID         string
	WordIDs    string
	Rate       string
	ProjectID  string
	IsDel      string
	CreateTime string
	UpdateTime string
}{
	ID:         "id",
	WordIDs:    "word_ids",
	Rate:       "rate",
	ProjectID:  "project_id",
	IsDel:      "is_del",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// DictVersion ES 版本更新表，每次修改词设置需增加版本号
type DictVersion struct {
	ID         uint32 `gorm:"primaryKey;column:id;type:int(11) unsigned;not null" json:"id"`
	Version    string `gorm:"column:version;type:varchar(50);not null;default:''" json:"version"`
	CreateTime uint32 `gorm:"column:create_time;type:int(11) unsigned;not null;default:0" json:"createTime"`
	UpdateTime uint32 `gorm:"column:update_time;type:int(11) unsigned;not null;default:0" json:"updateTime"`
}

// DictVersionColumns get sql column name.获取数据库列名
var DictVersionColumns = struct {
	ID         string
	Version    string
	CreateTime string
	UpdateTime string
}{
	ID:         "id",
	Version:    "version",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// DictWeight es搜索词权重表
type DictWeight struct {
	ID         uint32  `gorm:"primaryKey;column:id;type:int(10) unsigned;not null" json:"id"`
	WordID     uint32  `gorm:"unique;column:word_id;type:int(10) unsigned;not null;default:0" json:"wordId"` // 降权词id
	Weight     float64 `gorm:"column:weight;type:decimal(4,2) unsigned;not null;default:0.00" json:"weight"` // 权重
	ProjectID  uint32  `gorm:"column:project_id;type:int(11) unsigned;not null;default:0" json:"projectId"`  // 项目组id默认0
	IsDel      uint8   `gorm:"column:is_del;type:tinyint(2) unsigned;not null;default:0" json:"isDel"`
	CreateTime uint32  `gorm:"column:create_time;type:int(11) unsigned;not null;default:0" json:"createTime"`
	UpdateTime uint32  `gorm:"column:update_time;type:int(11) unsigned;not null;default:0" json:"updateTime"`
}

// DictWeightColumns get sql column name.获取数据库列名
var DictWeightColumns = struct {
	ID         string
	WordID     string
	Weight     string
	ProjectID  string
	IsDel      string
	CreateTime string
	UpdateTime string
}{
	ID:         "id",
	WordID:     "word_id",
	Weight:     "weight",
	ProjectID:  "project_id",
	IsDel:      "is_del",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// DictWord ES 扩展词表
type DictWord struct {
	ID         uint32 `gorm:"primaryKey;column:id;type:int(11) unsigned;not null" json:"id"`
	Word       string `gorm:"index:word;column:word;type:varchar(50);not null;default:''" json:"word"`
	Frequency  uint32 `gorm:"column:frequency;type:int(11) unsigned;not null;default:3" json:"frequency"` // 大于2W高频词，匹配度强制百分之百
	Classify   string `gorm:"column:classify;type:varchar(10);not null;default:n" json:"classify"`        // 词性（名词、动词）
	IsDel      uint8  `gorm:"column:is_del;type:tinyint(4) unsigned;not null;default:0" json:"isDel"`
	CreateTime uint32 `gorm:"column:create_time;type:int(11) unsigned;not null;default:0" json:"createTime"`
	UpdateTime uint32 `gorm:"column:update_time;type:int(11) unsigned;not null;default:0" json:"updateTime"`
}

// DictWordColumns get sql column name.获取数据库列名
var DictWordColumns = struct {
	ID         string
	Word       string
	Frequency  string
	Classify   string
	IsDel      string
	CreateTime string
	UpdateTime string
}{
	ID:         "id",
	Word:       "word",
	Frequency:  "frequency",
	Classify:   "classify",
	IsDel:      "is_del",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// DictWordCopy ES 扩展词表
type DictWordCopy struct {
	ID         uint32 `gorm:"primaryKey;column:id;type:int(11) unsigned;not null" json:"id"`
	Word       string `gorm:"index:word;column:word;type:varchar(50);not null;default:''" json:"word"`
	Frequency  uint32 `gorm:"column:frequency;type:int(11) unsigned;not null;default:3" json:"frequency"` // 大于2W高频词，匹配度强制百分之百
	Classify   string `gorm:"column:classify;type:varchar(10);not null;default:n" json:"classify"`        // 词性（名词、动词）
	IsDel      bool   `gorm:"column:is_del;type:tinyint(1) unsigned;not null;default:0" json:"isDel"`
	CreateTime uint32 `gorm:"column:create_time;type:int(11) unsigned;not null;default:0" json:"createTime"`
	UpdateTime uint32 `gorm:"column:update_time;type:int(11) unsigned;not null;default:0" json:"updateTime"`
}

// DictWordCopyColumns get sql column name.获取数据库列名
var DictWordCopyColumns = struct {
	ID         string
	Word       string
	Frequency  string
	Classify   string
	IsDel      string
	CreateTime string
	UpdateTime string
}{
	ID:         "id",
	Word:       "word",
	Frequency:  "frequency",
	Classify:   "classify",
	IsDel:      "is_del",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}
