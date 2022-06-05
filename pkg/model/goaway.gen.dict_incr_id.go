package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _DictIncrIDMgr struct {
	*_BaseMgr
}

// DictIncrIDMgr open func
func DictIncrIDMgr(db *gorm.DB) *_DictIncrIDMgr {
	if db == nil {
		panic(fmt.Errorf("DictIncrIDMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DictIncrIDMgr{_BaseMgr: &_BaseMgr{DB: db.Table("dict_incr_id"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DictIncrIDMgr) GetTableName() string {
	return "dict_incr_id"
}

// Reset 重置gorm会话
func (obj *_DictIncrIDMgr) Reset() *_DictIncrIDMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_DictIncrIDMgr) Get() (result DictIncrID, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictIncrID{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DictIncrIDMgr) Gets() (results []*DictIncrID, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictIncrID{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_DictIncrIDMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(DictIncrID{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_DictIncrIDMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// GetByOption 功能选项模式获取
func (obj *_DictIncrIDMgr) GetByOption(opts ...Option) (result DictIncrID, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DictIncrID{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DictIncrIDMgr) GetByOptions(opts ...Option) (results []*DictIncrID, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DictIncrID{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_DictIncrIDMgr) GetFromID(id uint32) (result DictIncrID, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictIncrID{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_DictIncrIDMgr) GetBatchFromID(ids []uint32) (results []*DictIncrID, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictIncrID{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_DictIncrIDMgr) FetchByPrimaryKey(id uint32) (result DictIncrID, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictIncrID{}).Where("`id` = ?", id).First(&result).Error

	return
}
