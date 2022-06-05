package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _DictProjectMgr struct {
	*_BaseMgr
}

// DictProjectMgr open func
func DictProjectMgr(db *gorm.DB) *_DictProjectMgr {
	if db == nil {
		panic(fmt.Errorf("DictProjectMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DictProjectMgr{_BaseMgr: &_BaseMgr{DB: db.Table("dict_project"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DictProjectMgr) GetTableName() string {
	return "dict_project"
}

// Reset 重置gorm会话
func (obj *_DictProjectMgr) Reset() *_DictProjectMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_DictProjectMgr) Get() (result DictProject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictProject{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DictProjectMgr) Gets() (results []*DictProject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictProject{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_DictProjectMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(DictProject{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_DictProjectMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_DictProjectMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithIsDel is_del获取
func (obj *_DictProjectMgr) WithIsDel(isDel bool) Option {
	return optionFunc(func(o *options) { o.query["is_del"] = isDel })
}

// WithCreateTime create_time获取
func (obj *_DictProjectMgr) WithCreateTime(createTime uint32) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_DictProjectMgr) WithUpdateTime(updateTime uint32) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_DictProjectMgr) GetByOption(opts ...Option) (result DictProject, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DictProject{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DictProjectMgr) GetByOptions(opts ...Option) (results []*DictProject, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DictProject{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_DictProjectMgr) GetFromID(id uint32) (result DictProject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictProject{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_DictProjectMgr) GetBatchFromID(ids []uint32) (results []*DictProject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictProject{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_DictProjectMgr) GetFromName(name string) (results []*DictProject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictProject{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_DictProjectMgr) GetBatchFromName(names []string) (results []*DictProject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictProject{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromIsDel 通过is_del获取内容
func (obj *_DictProjectMgr) GetFromIsDel(isDel bool) (results []*DictProject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictProject{}).Where("`is_del` = ?", isDel).Find(&results).Error

	return
}

// GetBatchFromIsDel 批量查找
func (obj *_DictProjectMgr) GetBatchFromIsDel(isDels []bool) (results []*DictProject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictProject{}).Where("`is_del` IN (?)", isDels).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_DictProjectMgr) GetFromCreateTime(createTime uint32) (results []*DictProject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictProject{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_DictProjectMgr) GetBatchFromCreateTime(createTimes []uint32) (results []*DictProject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictProject{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_DictProjectMgr) GetFromUpdateTime(updateTime uint32) (results []*DictProject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictProject{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_DictProjectMgr) GetBatchFromUpdateTime(updateTimes []uint32) (results []*DictProject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictProject{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_DictProjectMgr) FetchByPrimaryKey(id uint32) (result DictProject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictProject{}).Where("`id` = ?", id).First(&result).Error

	return
}
