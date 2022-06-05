package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _DictHyponymMgr struct {
	*_BaseMgr
}

// DictHyponymMgr open func
func DictHyponymMgr(db *gorm.DB) *_DictHyponymMgr {
	if db == nil {
		panic(fmt.Errorf("DictHyponymMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DictHyponymMgr{_BaseMgr: &_BaseMgr{DB: db.Table("dict_hyponym"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DictHyponymMgr) GetTableName() string {
	return "dict_hyponym"
}

// Reset 重置gorm会话
func (obj *_DictHyponymMgr) Reset() *_DictHyponymMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_DictHyponymMgr) Get() (result DictHyponym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHyponym{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DictHyponymMgr) Gets() (results []*DictHyponym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHyponym{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_DictHyponymMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(DictHyponym{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_DictHyponymMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithHypernymWordID hypernym_word_id获取 上位词id
func (obj *_DictHyponymMgr) WithHypernymWordID(hypernymWordID uint32) Option {
	return optionFunc(func(o *options) { o.query["hypernym_word_id"] = hypernymWordID })
}

// WithHyponymWordID hyponym_word_id获取 下位词id
func (obj *_DictHyponymMgr) WithHyponymWordID(hyponymWordID uint32) Option {
	return optionFunc(func(o *options) { o.query["hyponym_word_id"] = hyponymWordID })
}

// WithRate rate获取 下位词所占比例
func (obj *_DictHyponymMgr) WithRate(rate float32) Option {
	return optionFunc(func(o *options) { o.query["rate"] = rate })
}

// WithProjectID project_id获取
func (obj *_DictHyponymMgr) WithProjectID(projectID uint32) Option {
	return optionFunc(func(o *options) { o.query["project_id"] = projectID })
}

// WithIsDel is_del获取
func (obj *_DictHyponymMgr) WithIsDel(isDel bool) Option {
	return optionFunc(func(o *options) { o.query["is_del"] = isDel })
}

// WithCreateTime create_time获取
func (obj *_DictHyponymMgr) WithCreateTime(createTime uint32) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_DictHyponymMgr) WithUpdateTime(updateTime uint32) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_DictHyponymMgr) GetByOption(opts ...Option) (result DictHyponym, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DictHyponym{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DictHyponymMgr) GetByOptions(opts ...Option) (results []*DictHyponym, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DictHyponym{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_DictHyponymMgr) GetFromID(id uint32) (result DictHyponym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHyponym{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_DictHyponymMgr) GetBatchFromID(ids []uint32) (results []*DictHyponym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHyponym{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromHypernymWordID 通过hypernym_word_id获取内容 上位词id
func (obj *_DictHyponymMgr) GetFromHypernymWordID(hypernymWordID uint32) (results []*DictHyponym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHyponym{}).Where("`hypernym_word_id` = ?", hypernymWordID).Find(&results).Error

	return
}

// GetBatchFromHypernymWordID 批量查找 上位词id
func (obj *_DictHyponymMgr) GetBatchFromHypernymWordID(hypernymWordIDs []uint32) (results []*DictHyponym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHyponym{}).Where("`hypernym_word_id` IN (?)", hypernymWordIDs).Find(&results).Error

	return
}

// GetFromHyponymWordID 通过hyponym_word_id获取内容 下位词id
func (obj *_DictHyponymMgr) GetFromHyponymWordID(hyponymWordID uint32) (results []*DictHyponym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHyponym{}).Where("`hyponym_word_id` = ?", hyponymWordID).Find(&results).Error

	return
}

// GetBatchFromHyponymWordID 批量查找 下位词id
func (obj *_DictHyponymMgr) GetBatchFromHyponymWordID(hyponymWordIDs []uint32) (results []*DictHyponym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHyponym{}).Where("`hyponym_word_id` IN (?)", hyponymWordIDs).Find(&results).Error

	return
}

// GetFromRate 通过rate获取内容 下位词所占比例
func (obj *_DictHyponymMgr) GetFromRate(rate float32) (results []*DictHyponym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHyponym{}).Where("`rate` = ?", rate).Find(&results).Error

	return
}

// GetBatchFromRate 批量查找 下位词所占比例
func (obj *_DictHyponymMgr) GetBatchFromRate(rates []float32) (results []*DictHyponym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHyponym{}).Where("`rate` IN (?)", rates).Find(&results).Error

	return
}

// GetFromProjectID 通过project_id获取内容
func (obj *_DictHyponymMgr) GetFromProjectID(projectID uint32) (results []*DictHyponym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHyponym{}).Where("`project_id` = ?", projectID).Find(&results).Error

	return
}

// GetBatchFromProjectID 批量查找
func (obj *_DictHyponymMgr) GetBatchFromProjectID(projectIDs []uint32) (results []*DictHyponym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHyponym{}).Where("`project_id` IN (?)", projectIDs).Find(&results).Error

	return
}

// GetFromIsDel 通过is_del获取内容
func (obj *_DictHyponymMgr) GetFromIsDel(isDel bool) (results []*DictHyponym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHyponym{}).Where("`is_del` = ?", isDel).Find(&results).Error

	return
}

// GetBatchFromIsDel 批量查找
func (obj *_DictHyponymMgr) GetBatchFromIsDel(isDels []bool) (results []*DictHyponym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHyponym{}).Where("`is_del` IN (?)", isDels).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_DictHyponymMgr) GetFromCreateTime(createTime uint32) (results []*DictHyponym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHyponym{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_DictHyponymMgr) GetBatchFromCreateTime(createTimes []uint32) (results []*DictHyponym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHyponym{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_DictHyponymMgr) GetFromUpdateTime(updateTime uint32) (results []*DictHyponym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHyponym{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_DictHyponymMgr) GetBatchFromUpdateTime(updateTimes []uint32) (results []*DictHyponym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHyponym{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_DictHyponymMgr) FetchByPrimaryKey(id uint32) (result DictHyponym, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHyponym{}).Where("`id` = ?", id).First(&result).Error

	return
}
