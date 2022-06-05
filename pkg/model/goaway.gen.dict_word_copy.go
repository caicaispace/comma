package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _DictWordCopyMgr struct {
	*_BaseMgr
}

// DictWordCopyMgr open func
func DictWordCopyMgr(db *gorm.DB) *_DictWordCopyMgr {
	if db == nil {
		panic(fmt.Errorf("DictWordCopyMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DictWordCopyMgr{_BaseMgr: &_BaseMgr{DB: db.Table("dict_word_copy"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DictWordCopyMgr) GetTableName() string {
	return "dict_word_copy"
}

// Reset 重置gorm会话
func (obj *_DictWordCopyMgr) Reset() *_DictWordCopyMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_DictWordCopyMgr) Get() (result DictWordCopy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWordCopy{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DictWordCopyMgr) Gets() (results []*DictWordCopy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWordCopy{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_DictWordCopyMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(DictWordCopy{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_DictWordCopyMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithWord word获取
func (obj *_DictWordCopyMgr) WithWord(word string) Option {
	return optionFunc(func(o *options) { o.query["word"] = word })
}

// WithFrequency frequency获取 大于2W高频词，匹配度强制百分之百
func (obj *_DictWordCopyMgr) WithFrequency(frequency uint32) Option {
	return optionFunc(func(o *options) { o.query["frequency"] = frequency })
}

// WithClassify classify获取 词性（名词、动词）
func (obj *_DictWordCopyMgr) WithClassify(classify string) Option {
	return optionFunc(func(o *options) { o.query["classify"] = classify })
}

// WithIsDel is_del获取
func (obj *_DictWordCopyMgr) WithIsDel(isDel bool) Option {
	return optionFunc(func(o *options) { o.query["is_del"] = isDel })
}

// WithCreateTime create_time获取
func (obj *_DictWordCopyMgr) WithCreateTime(createTime uint32) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_DictWordCopyMgr) WithUpdateTime(updateTime uint32) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_DictWordCopyMgr) GetByOption(opts ...Option) (result DictWordCopy, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DictWordCopy{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DictWordCopyMgr) GetByOptions(opts ...Option) (results []*DictWordCopy, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DictWordCopy{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_DictWordCopyMgr) GetFromID(id uint32) (result DictWordCopy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWordCopy{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_DictWordCopyMgr) GetBatchFromID(ids []uint32) (results []*DictWordCopy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWordCopy{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromWord 通过word获取内容
func (obj *_DictWordCopyMgr) GetFromWord(word string) (results []*DictWordCopy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWordCopy{}).Where("`word` = ?", word).Find(&results).Error

	return
}

// GetBatchFromWord 批量查找
func (obj *_DictWordCopyMgr) GetBatchFromWord(words []string) (results []*DictWordCopy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWordCopy{}).Where("`word` IN (?)", words).Find(&results).Error

	return
}

// GetFromFrequency 通过frequency获取内容 大于2W高频词，匹配度强制百分之百
func (obj *_DictWordCopyMgr) GetFromFrequency(frequency uint32) (results []*DictWordCopy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWordCopy{}).Where("`frequency` = ?", frequency).Find(&results).Error

	return
}

// GetBatchFromFrequency 批量查找 大于2W高频词，匹配度强制百分之百
func (obj *_DictWordCopyMgr) GetBatchFromFrequency(frequencys []uint32) (results []*DictWordCopy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWordCopy{}).Where("`frequency` IN (?)", frequencys).Find(&results).Error

	return
}

// GetFromClassify 通过classify获取内容 词性（名词、动词）
func (obj *_DictWordCopyMgr) GetFromClassify(classify string) (results []*DictWordCopy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWordCopy{}).Where("`classify` = ?", classify).Find(&results).Error

	return
}

// GetBatchFromClassify 批量查找 词性（名词、动词）
func (obj *_DictWordCopyMgr) GetBatchFromClassify(classifys []string) (results []*DictWordCopy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWordCopy{}).Where("`classify` IN (?)", classifys).Find(&results).Error

	return
}

// GetFromIsDel 通过is_del获取内容
func (obj *_DictWordCopyMgr) GetFromIsDel(isDel bool) (results []*DictWordCopy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWordCopy{}).Where("`is_del` = ?", isDel).Find(&results).Error

	return
}

// GetBatchFromIsDel 批量查找
func (obj *_DictWordCopyMgr) GetBatchFromIsDel(isDels []bool) (results []*DictWordCopy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWordCopy{}).Where("`is_del` IN (?)", isDels).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_DictWordCopyMgr) GetFromCreateTime(createTime uint32) (results []*DictWordCopy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWordCopy{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_DictWordCopyMgr) GetBatchFromCreateTime(createTimes []uint32) (results []*DictWordCopy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWordCopy{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_DictWordCopyMgr) GetFromUpdateTime(updateTime uint32) (results []*DictWordCopy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWordCopy{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_DictWordCopyMgr) GetBatchFromUpdateTime(updateTimes []uint32) (results []*DictWordCopy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWordCopy{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_DictWordCopyMgr) FetchByPrimaryKey(id uint32) (result DictWordCopy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWordCopy{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByWord  获取多个内容
func (obj *_DictWordCopyMgr) FetchIndexByWord(word string) (results []*DictWordCopy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWordCopy{}).Where("`word` = ?", word).Find(&results).Error

	return
}
