package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _DictVersionMgr struct {
	*_BaseMgr
}

// DictVersionMgr open func
func DictVersionMgr(db *gorm.DB) *_DictVersionMgr {
	if db == nil {
		panic(fmt.Errorf("DictVersionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DictVersionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("dict_version"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DictVersionMgr) GetTableName() string {
	return "dict_version"
}

// Reset 重置gorm会话
func (obj *_DictVersionMgr) Reset() *_DictVersionMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_DictVersionMgr) Get() (result DictVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictVersion{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DictVersionMgr) Gets() (results []*DictVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictVersion{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_DictVersionMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(DictVersion{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_DictVersionMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithVersion version获取
func (obj *_DictVersionMgr) WithVersion(version string) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithCreateTime create_time获取
func (obj *_DictVersionMgr) WithCreateTime(createTime uint32) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_DictVersionMgr) WithUpdateTime(updateTime uint32) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_DictVersionMgr) GetByOption(opts ...Option) (result DictVersion, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DictVersion{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DictVersionMgr) GetByOptions(opts ...Option) (results []*DictVersion, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DictVersion{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_DictVersionMgr) GetFromID(id uint32) (result DictVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictVersion{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_DictVersionMgr) GetBatchFromID(ids []uint32) (results []*DictVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictVersion{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容
func (obj *_DictVersionMgr) GetFromVersion(version string) (results []*DictVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictVersion{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找
func (obj *_DictVersionMgr) GetBatchFromVersion(versions []string) (results []*DictVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictVersion{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_DictVersionMgr) GetFromCreateTime(createTime uint32) (results []*DictVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictVersion{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_DictVersionMgr) GetBatchFromCreateTime(createTimes []uint32) (results []*DictVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictVersion{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_DictVersionMgr) GetFromUpdateTime(updateTime uint32) (results []*DictVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictVersion{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_DictVersionMgr) GetBatchFromUpdateTime(updateTimes []uint32) (results []*DictVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictVersion{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_DictVersionMgr) FetchByPrimaryKey(id uint32) (result DictVersion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictVersion{}).Where("`id` = ?", id).First(&result).Error

	return
}
