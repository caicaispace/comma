package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _DictRedirectMgr struct {
	*_BaseMgr
}

// DictRedirectMgr open func
func DictRedirectMgr(db *gorm.DB) *_DictRedirectMgr {
	if db == nil {
		panic(fmt.Errorf("DictRedirectMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DictRedirectMgr{_BaseMgr: &_BaseMgr{DB: db.Table("dict_redirect"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DictRedirectMgr) GetTableName() string {
	return "dict_redirect"
}

// Reset 重置gorm会话
func (obj *_DictRedirectMgr) Reset() *_DictRedirectMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_DictRedirectMgr) Get() (result DictRedirect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictRedirect{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DictRedirectMgr) Gets() (results []*DictRedirect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictRedirect{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_DictRedirectMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(DictRedirect{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_DictRedirectMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithWord word获取
func (obj *_DictRedirectMgr) WithWord(word string) Option {
	return optionFunc(func(o *options) { o.query["word"] = word })
}

// WithRedirect redirect获取
func (obj *_DictRedirectMgr) WithRedirect(redirect string) Option {
	return optionFunc(func(o *options) { o.query["redirect"] = redirect })
}

// WithCreateTime create_time获取
func (obj *_DictRedirectMgr) WithCreateTime(createTime uint32) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_DictRedirectMgr) WithUpdateTime(updateTime uint32) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_DictRedirectMgr) GetByOption(opts ...Option) (result DictRedirect, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DictRedirect{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DictRedirectMgr) GetByOptions(opts ...Option) (results []*DictRedirect, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DictRedirect{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_DictRedirectMgr) GetFromID(id uint32) (result DictRedirect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictRedirect{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_DictRedirectMgr) GetBatchFromID(ids []uint32) (results []*DictRedirect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictRedirect{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromWord 通过word获取内容
func (obj *_DictRedirectMgr) GetFromWord(word string) (results []*DictRedirect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictRedirect{}).Where("`word` = ?", word).Find(&results).Error

	return
}

// GetBatchFromWord 批量查找
func (obj *_DictRedirectMgr) GetBatchFromWord(words []string) (results []*DictRedirect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictRedirect{}).Where("`word` IN (?)", words).Find(&results).Error

	return
}

// GetFromRedirect 通过redirect获取内容
func (obj *_DictRedirectMgr) GetFromRedirect(redirect string) (results []*DictRedirect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictRedirect{}).Where("`redirect` = ?", redirect).Find(&results).Error

	return
}

// GetBatchFromRedirect 批量查找
func (obj *_DictRedirectMgr) GetBatchFromRedirect(redirects []string) (results []*DictRedirect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictRedirect{}).Where("`redirect` IN (?)", redirects).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_DictRedirectMgr) GetFromCreateTime(createTime uint32) (results []*DictRedirect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictRedirect{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_DictRedirectMgr) GetBatchFromCreateTime(createTimes []uint32) (results []*DictRedirect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictRedirect{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_DictRedirectMgr) GetFromUpdateTime(updateTime uint32) (results []*DictRedirect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictRedirect{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_DictRedirectMgr) GetBatchFromUpdateTime(updateTimes []uint32) (results []*DictRedirect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictRedirect{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_DictRedirectMgr) FetchByPrimaryKey(id uint32) (result DictRedirect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictRedirect{}).Where("`id` = ?", id).First(&result).Error

	return
}
