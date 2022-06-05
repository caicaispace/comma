package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _DictPinyinMgr struct {
	*_BaseMgr
}

// DictPinyinMgr open func
func DictPinyinMgr(db *gorm.DB) *_DictPinyinMgr {
	if db == nil {
		panic(fmt.Errorf("DictPinyinMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DictPinyinMgr{_BaseMgr: &_BaseMgr{DB: db.Table("dict_pinyin"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DictPinyinMgr) GetTableName() string {
	return "dict_pinyin"
}

// Reset 重置gorm会话
func (obj *_DictPinyinMgr) Reset() *_DictPinyinMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_DictPinyinMgr) Get() (result DictPinyin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictPinyin{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DictPinyinMgr) Gets() (results []*DictPinyin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictPinyin{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_DictPinyinMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(DictPinyin{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_DictPinyinMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithKeyword keyword获取 关键词文本
func (obj *_DictPinyinMgr) WithKeyword(keyword string) Option {
	return optionFunc(func(o *options) { o.query["keyword"] = keyword })
}

// WithPinyin pinyin获取 拼音
func (obj *_DictPinyinMgr) WithPinyin(pinyin string) Option {
	return optionFunc(func(o *options) { o.query["pinyin"] = pinyin })
}

// WithPinyins pinyins获取 拼音
func (obj *_DictPinyinMgr) WithPinyins(pinyins string) Option {
	return optionFunc(func(o *options) { o.query["pinyins"] = pinyins })
}

// WithPinyinInitials pinyin_initials获取 拼音首字母
func (obj *_DictPinyinMgr) WithPinyinInitials(pinyinInitials string) Option {
	return optionFunc(func(o *options) { o.query["pinyin_initials"] = pinyinInitials })
}

// WithEnglish english获取 英文
func (obj *_DictPinyinMgr) WithEnglish(english string) Option {
	return optionFunc(func(o *options) { o.query["english"] = english })
}

// WithUseNum use_num获取
func (obj *_DictPinyinMgr) WithUseNum(useNum uint32) Option {
	return optionFunc(func(o *options) { o.query["use_num"] = useNum })
}

// WithHomophonyNum homophony_num获取 同音次数
func (obj *_DictPinyinMgr) WithHomophonyNum(homophonyNum uint32) Option {
	return optionFunc(func(o *options) { o.query["homophony_num"] = homophonyNum })
}

// WithIsDel is_del获取
func (obj *_DictPinyinMgr) WithIsDel(isDel uint8) Option {
	return optionFunc(func(o *options) { o.query["is_del"] = isDel })
}

// WithCreateTime create_time获取
func (obj *_DictPinyinMgr) WithCreateTime(createTime uint32) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_DictPinyinMgr) WithUpdateTime(updateTime uint32) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_DictPinyinMgr) GetByOption(opts ...Option) (result DictPinyin, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DictPinyin{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DictPinyinMgr) GetByOptions(opts ...Option) (results []*DictPinyin, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DictPinyin{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_DictPinyinMgr) GetFromID(id uint32) (result DictPinyin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictPinyin{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_DictPinyinMgr) GetBatchFromID(ids []uint32) (results []*DictPinyin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictPinyin{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromKeyword 通过keyword获取内容 关键词文本
func (obj *_DictPinyinMgr) GetFromKeyword(keyword string) (result DictPinyin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictPinyin{}).Where("`keyword` = ?", keyword).First(&result).Error

	return
}

// GetBatchFromKeyword 批量查找 关键词文本
func (obj *_DictPinyinMgr) GetBatchFromKeyword(keywords []string) (results []*DictPinyin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictPinyin{}).Where("`keyword` IN (?)", keywords).Find(&results).Error

	return
}

// GetFromPinyin 通过pinyin获取内容 拼音
func (obj *_DictPinyinMgr) GetFromPinyin(pinyin string) (results []*DictPinyin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictPinyin{}).Where("`pinyin` = ?", pinyin).Find(&results).Error

	return
}

// GetBatchFromPinyin 批量查找 拼音
func (obj *_DictPinyinMgr) GetBatchFromPinyin(pinyins []string) (results []*DictPinyin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictPinyin{}).Where("`pinyin` IN (?)", pinyins).Find(&results).Error

	return
}

// GetFromPinyins 通过pinyins获取内容 拼音
func (obj *_DictPinyinMgr) GetFromPinyins(pinyins string) (results []*DictPinyin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictPinyin{}).Where("`pinyins` = ?", pinyins).Find(&results).Error

	return
}

// GetBatchFromPinyins 批量查找 拼音
func (obj *_DictPinyinMgr) GetBatchFromPinyins(pinyinss []string) (results []*DictPinyin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictPinyin{}).Where("`pinyins` IN (?)", pinyinss).Find(&results).Error

	return
}

// GetFromPinyinInitials 通过pinyin_initials获取内容 拼音首字母
func (obj *_DictPinyinMgr) GetFromPinyinInitials(pinyinInitials string) (results []*DictPinyin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictPinyin{}).Where("`pinyin_initials` = ?", pinyinInitials).Find(&results).Error

	return
}

// GetBatchFromPinyinInitials 批量查找 拼音首字母
func (obj *_DictPinyinMgr) GetBatchFromPinyinInitials(pinyinInitialss []string) (results []*DictPinyin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictPinyin{}).Where("`pinyin_initials` IN (?)", pinyinInitialss).Find(&results).Error

	return
}

// GetFromEnglish 通过english获取内容 英文
func (obj *_DictPinyinMgr) GetFromEnglish(english string) (results []*DictPinyin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictPinyin{}).Where("`english` = ?", english).Find(&results).Error

	return
}

// GetBatchFromEnglish 批量查找 英文
func (obj *_DictPinyinMgr) GetBatchFromEnglish(englishs []string) (results []*DictPinyin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictPinyin{}).Where("`english` IN (?)", englishs).Find(&results).Error

	return
}

// GetFromUseNum 通过use_num获取内容
func (obj *_DictPinyinMgr) GetFromUseNum(useNum uint32) (results []*DictPinyin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictPinyin{}).Where("`use_num` = ?", useNum).Find(&results).Error

	return
}

// GetBatchFromUseNum 批量查找
func (obj *_DictPinyinMgr) GetBatchFromUseNum(useNums []uint32) (results []*DictPinyin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictPinyin{}).Where("`use_num` IN (?)", useNums).Find(&results).Error

	return
}

// GetFromHomophonyNum 通过homophony_num获取内容 同音次数
func (obj *_DictPinyinMgr) GetFromHomophonyNum(homophonyNum uint32) (results []*DictPinyin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictPinyin{}).Where("`homophony_num` = ?", homophonyNum).Find(&results).Error

	return
}

// GetBatchFromHomophonyNum 批量查找 同音次数
func (obj *_DictPinyinMgr) GetBatchFromHomophonyNum(homophonyNums []uint32) (results []*DictPinyin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictPinyin{}).Where("`homophony_num` IN (?)", homophonyNums).Find(&results).Error

	return
}

// GetFromIsDel 通过is_del获取内容
func (obj *_DictPinyinMgr) GetFromIsDel(isDel uint8) (results []*DictPinyin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictPinyin{}).Where("`is_del` = ?", isDel).Find(&results).Error

	return
}

// GetBatchFromIsDel 批量查找
func (obj *_DictPinyinMgr) GetBatchFromIsDel(isDels []uint8) (results []*DictPinyin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictPinyin{}).Where("`is_del` IN (?)", isDels).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_DictPinyinMgr) GetFromCreateTime(createTime uint32) (results []*DictPinyin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictPinyin{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_DictPinyinMgr) GetBatchFromCreateTime(createTimes []uint32) (results []*DictPinyin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictPinyin{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_DictPinyinMgr) GetFromUpdateTime(updateTime uint32) (results []*DictPinyin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictPinyin{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_DictPinyinMgr) GetBatchFromUpdateTime(updateTimes []uint32) (results []*DictPinyin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictPinyin{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_DictPinyinMgr) FetchByPrimaryKey(id uint32) (result DictPinyin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictPinyin{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByKeyword primary or index 获取唯一内容
func (obj *_DictPinyinMgr) FetchUniqueByKeyword(keyword string) (result DictPinyin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictPinyin{}).Where("`keyword` = ?", keyword).First(&result).Error

	return
}
