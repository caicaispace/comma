package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _DictWordMgr struct {
	*_BaseMgr
}

// DictWordMgr open func
func DictWordMgr(db *gorm.DB) *_DictWordMgr {
	if db == nil {
		panic(fmt.Errorf("DictWordMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DictWordMgr{_BaseMgr: &_BaseMgr{DB: db.Table("dict_word"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DictWordMgr) GetTableName() string {
	return "dict_word"
}

// Reset 重置gorm会话
func (obj *_DictWordMgr) Reset() *_DictWordMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_DictWordMgr) Get() (result DictWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWord{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DictWordMgr) Gets() (results []*DictWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWord{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_DictWordMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(DictWord{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_DictWordMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithWord word获取
func (obj *_DictWordMgr) WithWord(word string) Option {
	return optionFunc(func(o *options) { o.query["word"] = word })
}

// WithFrequency frequency获取 大于2W高频词，匹配度强制百分之百
func (obj *_DictWordMgr) WithFrequency(frequency uint32) Option {
	return optionFunc(func(o *options) { o.query["frequency"] = frequency })
}

// WithClassify classify获取 词性（名词、动词）
func (obj *_DictWordMgr) WithClassify(classify string) Option {
	return optionFunc(func(o *options) { o.query["classify"] = classify })
}

// WithIsDel is_del获取
func (obj *_DictWordMgr) WithIsDel(isDel uint8) Option {
	return optionFunc(func(o *options) { o.query["is_del"] = isDel })
}

// WithCreateTime create_time获取
func (obj *_DictWordMgr) WithCreateTime(createTime uint32) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_DictWordMgr) WithUpdateTime(updateTime uint32) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_DictWordMgr) GetByOption(opts ...Option) (result DictWord, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DictWord{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DictWordMgr) GetByOptions(opts ...Option) (results []*DictWord, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DictWord{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_DictWordMgr) GetFromID(id uint32) (result DictWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWord{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_DictWordMgr) GetBatchFromID(ids []uint32) (results []*DictWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWord{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromWord 通过word获取内容
func (obj *_DictWordMgr) GetFromWord(word string) (results []*DictWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWord{}).Where("`word` = ?", word).Find(&results).Error

	return
}

// GetBatchFromWord 批量查找
func (obj *_DictWordMgr) GetBatchFromWord(words []string) (results []*DictWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWord{}).Where("`word` IN (?)", words).Find(&results).Error

	return
}

// GetFromFrequency 通过frequency获取内容 大于2W高频词，匹配度强制百分之百
func (obj *_DictWordMgr) GetFromFrequency(frequency uint32) (results []*DictWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWord{}).Where("`frequency` = ?", frequency).Find(&results).Error

	return
}

// GetBatchFromFrequency 批量查找 大于2W高频词，匹配度强制百分之百
func (obj *_DictWordMgr) GetBatchFromFrequency(frequencys []uint32) (results []*DictWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWord{}).Where("`frequency` IN (?)", frequencys).Find(&results).Error

	return
}

// GetFromClassify 通过classify获取内容 词性（名词、动词）
func (obj *_DictWordMgr) GetFromClassify(classify string) (results []*DictWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWord{}).Where("`classify` = ?", classify).Find(&results).Error

	return
}

// GetBatchFromClassify 批量查找 词性（名词、动词）
func (obj *_DictWordMgr) GetBatchFromClassify(classifys []string) (results []*DictWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWord{}).Where("`classify` IN (?)", classifys).Find(&results).Error

	return
}

// GetFromIsDel 通过is_del获取内容
func (obj *_DictWordMgr) GetFromIsDel(isDel uint8) (results []*DictWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWord{}).Where("`is_del` = ?", isDel).Find(&results).Error

	return
}

// GetBatchFromIsDel 批量查找
func (obj *_DictWordMgr) GetBatchFromIsDel(isDels []uint8) (results []*DictWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWord{}).Where("`is_del` IN (?)", isDels).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_DictWordMgr) GetFromCreateTime(createTime uint32) (results []*DictWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWord{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_DictWordMgr) GetBatchFromCreateTime(createTimes []uint32) (results []*DictWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWord{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_DictWordMgr) GetFromUpdateTime(updateTime uint32) (results []*DictWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWord{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_DictWordMgr) GetBatchFromUpdateTime(updateTimes []uint32) (results []*DictWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWord{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_DictWordMgr) FetchByPrimaryKey(id uint32) (result DictWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWord{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByWord  获取多个内容
func (obj *_DictWordMgr) FetchIndexByWord(word string) (results []*DictWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWord{}).Where("`word` = ?", word).Find(&results).Error

	return
}
