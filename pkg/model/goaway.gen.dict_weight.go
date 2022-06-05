package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _DictWeightMgr struct {
	*_BaseMgr
}

// DictWeightMgr open func
func DictWeightMgr(db *gorm.DB) *_DictWeightMgr {
	if db == nil {
		panic(fmt.Errorf("DictWeightMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DictWeightMgr{_BaseMgr: &_BaseMgr{DB: db.Table("dict_weight"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DictWeightMgr) GetTableName() string {
	return "dict_weight"
}

// Reset 重置gorm会话
func (obj *_DictWeightMgr) Reset() *_DictWeightMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_DictWeightMgr) Get() (result DictWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWeight{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DictWeightMgr) Gets() (results []*DictWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWeight{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_DictWeightMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(DictWeight{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_DictWeightMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithWordID word_id获取 降权词id
func (obj *_DictWeightMgr) WithWordID(wordID uint32) Option {
	return optionFunc(func(o *options) { o.query["word_id"] = wordID })
}

// WithWeight weight获取 权重
func (obj *_DictWeightMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithProjectID project_id获取 项目组id默认0
func (obj *_DictWeightMgr) WithProjectID(projectID uint32) Option {
	return optionFunc(func(o *options) { o.query["project_id"] = projectID })
}

// WithIsDel is_del获取
func (obj *_DictWeightMgr) WithIsDel(isDel uint8) Option {
	return optionFunc(func(o *options) { o.query["is_del"] = isDel })
}

// WithCreateTime create_time获取
func (obj *_DictWeightMgr) WithCreateTime(createTime uint32) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_DictWeightMgr) WithUpdateTime(updateTime uint32) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_DictWeightMgr) GetByOption(opts ...Option) (result DictWeight, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DictWeight{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DictWeightMgr) GetByOptions(opts ...Option) (results []*DictWeight, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DictWeight{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_DictWeightMgr) GetFromID(id uint32) (result DictWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWeight{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_DictWeightMgr) GetBatchFromID(ids []uint32) (results []*DictWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWeight{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromWordID 通过word_id获取内容 降权词id
func (obj *_DictWeightMgr) GetFromWordID(wordID uint32) (result DictWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWeight{}).Where("`word_id` = ?", wordID).First(&result).Error

	return
}

// GetBatchFromWordID 批量查找 降权词id
func (obj *_DictWeightMgr) GetBatchFromWordID(wordIDs []uint32) (results []*DictWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWeight{}).Where("`word_id` IN (?)", wordIDs).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容 权重
func (obj *_DictWeightMgr) GetFromWeight(weight float64) (results []*DictWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWeight{}).Where("`weight` = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量查找 权重
func (obj *_DictWeightMgr) GetBatchFromWeight(weights []float64) (results []*DictWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWeight{}).Where("`weight` IN (?)", weights).Find(&results).Error

	return
}

// GetFromProjectID 通过project_id获取内容 项目组id默认0
func (obj *_DictWeightMgr) GetFromProjectID(projectID uint32) (results []*DictWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWeight{}).Where("`project_id` = ?", projectID).Find(&results).Error

	return
}

// GetBatchFromProjectID 批量查找 项目组id默认0
func (obj *_DictWeightMgr) GetBatchFromProjectID(projectIDs []uint32) (results []*DictWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWeight{}).Where("`project_id` IN (?)", projectIDs).Find(&results).Error

	return
}

// GetFromIsDel 通过is_del获取内容
func (obj *_DictWeightMgr) GetFromIsDel(isDel uint8) (results []*DictWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWeight{}).Where("`is_del` = ?", isDel).Find(&results).Error

	return
}

// GetBatchFromIsDel 批量查找
func (obj *_DictWeightMgr) GetBatchFromIsDel(isDels []uint8) (results []*DictWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWeight{}).Where("`is_del` IN (?)", isDels).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_DictWeightMgr) GetFromCreateTime(createTime uint32) (results []*DictWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWeight{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_DictWeightMgr) GetBatchFromCreateTime(createTimes []uint32) (results []*DictWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWeight{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_DictWeightMgr) GetFromUpdateTime(updateTime uint32) (results []*DictWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWeight{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_DictWeightMgr) GetBatchFromUpdateTime(updateTimes []uint32) (results []*DictWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWeight{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_DictWeightMgr) FetchByPrimaryKey(id uint32) (result DictWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWeight{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByIDxWordid primary or index 获取唯一内容
func (obj *_DictWeightMgr) FetchUniqueByIDxWordid(wordID uint32) (result DictWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictWeight{}).Where("`word_id` = ?", wordID).First(&result).Error

	return
}
