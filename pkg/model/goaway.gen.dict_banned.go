package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _DictBannedMgr struct {
	*_BaseMgr
}

// DictBannedMgr open func
func DictBannedMgr(db *gorm.DB) *_DictBannedMgr {
	if db == nil {
		panic(fmt.Errorf("DictBannedMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DictBannedMgr{_BaseMgr: &_BaseMgr{DB: db.Table("dict_banned"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DictBannedMgr) GetTableName() string {
	return "dict_banned"
}

// Reset 重置gorm会话
func (obj *_DictBannedMgr) Reset() *_DictBannedMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_DictBannedMgr) Get() (result DictBanned, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictBanned{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DictBannedMgr) Gets() (results []*DictBanned, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictBanned{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_DictBannedMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(DictBanned{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_DictBannedMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithWordID word_id获取
func (obj *_DictBannedMgr) WithWordID(wordID uint32) Option {
	return optionFunc(func(o *options) { o.query["word_id"] = wordID })
}

// WithProjectID project_id获取
func (obj *_DictBannedMgr) WithProjectID(projectID uint32) Option {
	return optionFunc(func(o *options) { o.query["project_id"] = projectID })
}

// WithIsDel is_del获取
func (obj *_DictBannedMgr) WithIsDel(isDel uint8) Option {
	return optionFunc(func(o *options) { o.query["is_del"] = isDel })
}

// WithCreateTime create_time获取
func (obj *_DictBannedMgr) WithCreateTime(createTime uint32) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_DictBannedMgr) WithUpdateTime(updateTime uint32) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_DictBannedMgr) GetByOption(opts ...Option) (result DictBanned, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DictBanned{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DictBannedMgr) GetByOptions(opts ...Option) (results []*DictBanned, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DictBanned{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_DictBannedMgr) GetFromID(id uint32) (result DictBanned, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictBanned{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_DictBannedMgr) GetBatchFromID(ids []uint32) (results []*DictBanned, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictBanned{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromWordID 通过word_id获取内容
func (obj *_DictBannedMgr) GetFromWordID(wordID uint32) (results []*DictBanned, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictBanned{}).Where("`word_id` = ?", wordID).Find(&results).Error

	return
}

// GetBatchFromWordID 批量查找
func (obj *_DictBannedMgr) GetBatchFromWordID(wordIDs []uint32) (results []*DictBanned, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictBanned{}).Where("`word_id` IN (?)", wordIDs).Find(&results).Error

	return
}

// GetFromProjectID 通过project_id获取内容
func (obj *_DictBannedMgr) GetFromProjectID(projectID uint32) (results []*DictBanned, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictBanned{}).Where("`project_id` = ?", projectID).Find(&results).Error

	return
}

// GetBatchFromProjectID 批量查找
func (obj *_DictBannedMgr) GetBatchFromProjectID(projectIDs []uint32) (results []*DictBanned, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictBanned{}).Where("`project_id` IN (?)", projectIDs).Find(&results).Error

	return
}

// GetFromIsDel 通过is_del获取内容
func (obj *_DictBannedMgr) GetFromIsDel(isDel uint8) (results []*DictBanned, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictBanned{}).Where("`is_del` = ?", isDel).Find(&results).Error

	return
}

// GetBatchFromIsDel 批量查找
func (obj *_DictBannedMgr) GetBatchFromIsDel(isDels []uint8) (results []*DictBanned, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictBanned{}).Where("`is_del` IN (?)", isDels).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_DictBannedMgr) GetFromCreateTime(createTime uint32) (results []*DictBanned, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictBanned{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_DictBannedMgr) GetBatchFromCreateTime(createTimes []uint32) (results []*DictBanned, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictBanned{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_DictBannedMgr) GetFromUpdateTime(updateTime uint32) (results []*DictBanned, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictBanned{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_DictBannedMgr) GetBatchFromUpdateTime(updateTimes []uint32) (results []*DictBanned, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictBanned{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_DictBannedMgr) FetchByPrimaryKey(id uint32) (result DictBanned, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictBanned{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueIndexByWordIDProjectID primary or index 获取唯一内容
func (obj *_DictBannedMgr) FetchUniqueIndexByWordIDProjectID(wordID uint32, projectID uint32) (result DictBanned, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictBanned{}).Where("`word_id` = ? AND `project_id` = ?", wordID, projectID).First(&result).Error

	return
}
