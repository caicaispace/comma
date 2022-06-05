package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _DictHighFrequencyMgr struct {
	*_BaseMgr
}

// DictHighFrequencyMgr open func
func DictHighFrequencyMgr(db *gorm.DB) *_DictHighFrequencyMgr {
	if db == nil {
		panic(fmt.Errorf("DictHighFrequencyMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DictHighFrequencyMgr{_BaseMgr: &_BaseMgr{DB: db.Table("dict_high_frequency"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DictHighFrequencyMgr) GetTableName() string {
	return "dict_high_frequency"
}

// Reset 重置gorm会话
func (obj *_DictHighFrequencyMgr) Reset() *_DictHighFrequencyMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_DictHighFrequencyMgr) Get() (result DictHighFrequency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHighFrequency{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DictHighFrequencyMgr) Gets() (results []*DictHighFrequency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHighFrequency{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_DictHighFrequencyMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(DictHighFrequency{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_DictHighFrequencyMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithWordID word_id获取
func (obj *_DictHighFrequencyMgr) WithWordID(wordID uint32) Option {
	return optionFunc(func(o *options) { o.query["word_id"] = wordID })
}

// WithProjectID project_id获取
func (obj *_DictHighFrequencyMgr) WithProjectID(projectID uint32) Option {
	return optionFunc(func(o *options) { o.query["project_id"] = projectID })
}

// WithIsDel is_del获取
func (obj *_DictHighFrequencyMgr) WithIsDel(isDel bool) Option {
	return optionFunc(func(o *options) { o.query["is_del"] = isDel })
}

// WithCreateTime create_time获取
func (obj *_DictHighFrequencyMgr) WithCreateTime(createTime uint32) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_DictHighFrequencyMgr) WithUpdateTime(updateTime uint32) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_DictHighFrequencyMgr) GetByOption(opts ...Option) (result DictHighFrequency, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DictHighFrequency{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DictHighFrequencyMgr) GetByOptions(opts ...Option) (results []*DictHighFrequency, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DictHighFrequency{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_DictHighFrequencyMgr) GetFromID(id uint32) (result DictHighFrequency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHighFrequency{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_DictHighFrequencyMgr) GetBatchFromID(ids []uint32) (results []*DictHighFrequency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHighFrequency{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromWordID 通过word_id获取内容
func (obj *_DictHighFrequencyMgr) GetFromWordID(wordID uint32) (results []*DictHighFrequency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHighFrequency{}).Where("`word_id` = ?", wordID).Find(&results).Error

	return
}

// GetBatchFromWordID 批量查找
func (obj *_DictHighFrequencyMgr) GetBatchFromWordID(wordIDs []uint32) (results []*DictHighFrequency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHighFrequency{}).Where("`word_id` IN (?)", wordIDs).Find(&results).Error

	return
}

// GetFromProjectID 通过project_id获取内容
func (obj *_DictHighFrequencyMgr) GetFromProjectID(projectID uint32) (results []*DictHighFrequency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHighFrequency{}).Where("`project_id` = ?", projectID).Find(&results).Error

	return
}

// GetBatchFromProjectID 批量查找
func (obj *_DictHighFrequencyMgr) GetBatchFromProjectID(projectIDs []uint32) (results []*DictHighFrequency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHighFrequency{}).Where("`project_id` IN (?)", projectIDs).Find(&results).Error

	return
}

// GetFromIsDel 通过is_del获取内容
func (obj *_DictHighFrequencyMgr) GetFromIsDel(isDel bool) (results []*DictHighFrequency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHighFrequency{}).Where("`is_del` = ?", isDel).Find(&results).Error

	return
}

// GetBatchFromIsDel 批量查找
func (obj *_DictHighFrequencyMgr) GetBatchFromIsDel(isDels []bool) (results []*DictHighFrequency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHighFrequency{}).Where("`is_del` IN (?)", isDels).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_DictHighFrequencyMgr) GetFromCreateTime(createTime uint32) (results []*DictHighFrequency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHighFrequency{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_DictHighFrequencyMgr) GetBatchFromCreateTime(createTimes []uint32) (results []*DictHighFrequency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHighFrequency{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_DictHighFrequencyMgr) GetFromUpdateTime(updateTime uint32) (results []*DictHighFrequency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHighFrequency{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_DictHighFrequencyMgr) GetBatchFromUpdateTime(updateTimes []uint32) (results []*DictHighFrequency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHighFrequency{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_DictHighFrequencyMgr) FetchByPrimaryKey(id uint32) (result DictHighFrequency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictHighFrequency{}).Where("`id` = ?", id).First(&result).Error

	return
}
