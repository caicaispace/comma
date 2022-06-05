package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _DictStopMgr struct {
	*_BaseMgr
}

// DictStopMgr open func
func DictStopMgr(db *gorm.DB) *_DictStopMgr {
	if db == nil {
		panic(fmt.Errorf("DictStopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DictStopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("dict_stop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DictStopMgr) GetTableName() string {
	return "dict_stop"
}

// Reset 重置gorm会话
func (obj *_DictStopMgr) Reset() *_DictStopMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_DictStopMgr) Get() (result DictStop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictStop{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DictStopMgr) Gets() (results []*DictStop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictStop{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_DictStopMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(DictStop{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_DictStopMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithWordID word_id获取
func (obj *_DictStopMgr) WithWordID(wordID uint32) Option {
	return optionFunc(func(o *options) { o.query["word_id"] = wordID })
}

// WithProjectID project_id获取
func (obj *_DictStopMgr) WithProjectID(projectID uint32) Option {
	return optionFunc(func(o *options) { o.query["project_id"] = projectID })
}

// WithIsDel is_del获取
func (obj *_DictStopMgr) WithIsDel(isDel bool) Option {
	return optionFunc(func(o *options) { o.query["is_del"] = isDel })
}

// WithCreateTime create_time获取
func (obj *_DictStopMgr) WithCreateTime(createTime uint32) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_DictStopMgr) WithUpdateTime(updateTime uint32) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_DictStopMgr) GetByOption(opts ...Option) (result DictStop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DictStop{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DictStopMgr) GetByOptions(opts ...Option) (results []*DictStop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DictStop{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_DictStopMgr) GetFromID(id uint32) (result DictStop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictStop{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_DictStopMgr) GetBatchFromID(ids []uint32) (results []*DictStop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictStop{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromWordID 通过word_id获取内容
func (obj *_DictStopMgr) GetFromWordID(wordID uint32) (results []*DictStop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictStop{}).Where("`word_id` = ?", wordID).Find(&results).Error

	return
}

// GetBatchFromWordID 批量查找
func (obj *_DictStopMgr) GetBatchFromWordID(wordIDs []uint32) (results []*DictStop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictStop{}).Where("`word_id` IN (?)", wordIDs).Find(&results).Error

	return
}

// GetFromProjectID 通过project_id获取内容
func (obj *_DictStopMgr) GetFromProjectID(projectID uint32) (results []*DictStop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictStop{}).Where("`project_id` = ?", projectID).Find(&results).Error

	return
}

// GetBatchFromProjectID 批量查找
func (obj *_DictStopMgr) GetBatchFromProjectID(projectIDs []uint32) (results []*DictStop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictStop{}).Where("`project_id` IN (?)", projectIDs).Find(&results).Error

	return
}

// GetFromIsDel 通过is_del获取内容
func (obj *_DictStopMgr) GetFromIsDel(isDel bool) (results []*DictStop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictStop{}).Where("`is_del` = ?", isDel).Find(&results).Error

	return
}

// GetBatchFromIsDel 批量查找
func (obj *_DictStopMgr) GetBatchFromIsDel(isDels []bool) (results []*DictStop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictStop{}).Where("`is_del` IN (?)", isDels).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_DictStopMgr) GetFromCreateTime(createTime uint32) (results []*DictStop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictStop{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_DictStopMgr) GetBatchFromCreateTime(createTimes []uint32) (results []*DictStop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictStop{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_DictStopMgr) GetFromUpdateTime(updateTime uint32) (results []*DictStop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictStop{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_DictStopMgr) GetBatchFromUpdateTime(updateTimes []uint32) (results []*DictStop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictStop{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_DictStopMgr) FetchByPrimaryKey(id uint32) (result DictStop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictStop{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueIndexByWordIDProjectID primary or index 获取唯一内容
func (obj *_DictStopMgr) FetchUniqueIndexByWordIDProjectID(wordID uint32, projectID uint32) (result DictStop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictStop{}).Where("`word_id` = ? AND `project_id` = ?", wordID, projectID).First(&result).Error

	return
}
