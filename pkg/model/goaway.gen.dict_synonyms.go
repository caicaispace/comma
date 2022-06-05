package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _DictSynonymsMgr struct {
	*_BaseMgr
}

// DictSynonymsMgr open func
func DictSynonymsMgr(db *gorm.DB) *_DictSynonymsMgr {
	if db == nil {
		panic(fmt.Errorf("DictSynonymsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DictSynonymsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("dict_synonyms"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DictSynonymsMgr) GetTableName() string {
	return "dict_synonyms"
}

// Reset 重置gorm会话
func (obj *_DictSynonymsMgr) Reset() *_DictSynonymsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_DictSynonymsMgr) Get() (result DictSynonyms, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictSynonyms{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DictSynonymsMgr) Gets() (results []*DictSynonyms, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictSynonyms{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_DictSynonymsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(DictSynonyms{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_DictSynonymsMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithWordIDs word_ids获取
func (obj *_DictSynonymsMgr) WithWordIDs(wordIDs string) Option {
	return optionFunc(func(o *options) { o.query["word_ids"] = wordIDs })
}

// WithRate rate获取 比例
func (obj *_DictSynonymsMgr) WithRate(rate float32) Option {
	return optionFunc(func(o *options) { o.query["rate"] = rate })
}

// WithProjectID project_id获取 项目组id默认0
func (obj *_DictSynonymsMgr) WithProjectID(projectID uint32) Option {
	return optionFunc(func(o *options) { o.query["project_id"] = projectID })
}

// WithIsDel is_del获取
func (obj *_DictSynonymsMgr) WithIsDel(isDel bool) Option {
	return optionFunc(func(o *options) { o.query["is_del"] = isDel })
}

// WithCreateTime create_time获取
func (obj *_DictSynonymsMgr) WithCreateTime(createTime uint32) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_DictSynonymsMgr) WithUpdateTime(updateTime uint32) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_DictSynonymsMgr) GetByOption(opts ...Option) (result DictSynonyms, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DictSynonyms{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DictSynonymsMgr) GetByOptions(opts ...Option) (results []*DictSynonyms, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DictSynonyms{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_DictSynonymsMgr) GetFromID(id uint32) (result DictSynonyms, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictSynonyms{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_DictSynonymsMgr) GetBatchFromID(ids []uint32) (results []*DictSynonyms, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictSynonyms{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromWordIDs 通过word_ids获取内容
func (obj *_DictSynonymsMgr) GetFromWordIDs(wordIDs string) (results []*DictSynonyms, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictSynonyms{}).Where("`word_ids` = ?", wordIDs).Find(&results).Error

	return
}

// GetBatchFromWordIDs 批量查找
func (obj *_DictSynonymsMgr) GetBatchFromWordIDs(wordIDss []string) (results []*DictSynonyms, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictSynonyms{}).Where("`word_ids` IN (?)", wordIDss).Find(&results).Error

	return
}

// GetFromRate 通过rate获取内容 比例
func (obj *_DictSynonymsMgr) GetFromRate(rate float32) (results []*DictSynonyms, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictSynonyms{}).Where("`rate` = ?", rate).Find(&results).Error

	return
}

// GetBatchFromRate 批量查找 比例
func (obj *_DictSynonymsMgr) GetBatchFromRate(rates []float32) (results []*DictSynonyms, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictSynonyms{}).Where("`rate` IN (?)", rates).Find(&results).Error

	return
}

// GetFromProjectID 通过project_id获取内容 项目组id默认0
func (obj *_DictSynonymsMgr) GetFromProjectID(projectID uint32) (results []*DictSynonyms, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictSynonyms{}).Where("`project_id` = ?", projectID).Find(&results).Error

	return
}

// GetBatchFromProjectID 批量查找 项目组id默认0
func (obj *_DictSynonymsMgr) GetBatchFromProjectID(projectIDs []uint32) (results []*DictSynonyms, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictSynonyms{}).Where("`project_id` IN (?)", projectIDs).Find(&results).Error

	return
}

// GetFromIsDel 通过is_del获取内容
func (obj *_DictSynonymsMgr) GetFromIsDel(isDel bool) (results []*DictSynonyms, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictSynonyms{}).Where("`is_del` = ?", isDel).Find(&results).Error

	return
}

// GetBatchFromIsDel 批量查找
func (obj *_DictSynonymsMgr) GetBatchFromIsDel(isDels []bool) (results []*DictSynonyms, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictSynonyms{}).Where("`is_del` IN (?)", isDels).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_DictSynonymsMgr) GetFromCreateTime(createTime uint32) (results []*DictSynonyms, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictSynonyms{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_DictSynonymsMgr) GetBatchFromCreateTime(createTimes []uint32) (results []*DictSynonyms, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictSynonyms{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_DictSynonymsMgr) GetFromUpdateTime(updateTime uint32) (results []*DictSynonyms, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictSynonyms{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_DictSynonymsMgr) GetBatchFromUpdateTime(updateTimes []uint32) (results []*DictSynonyms, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictSynonyms{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_DictSynonymsMgr) FetchByPrimaryKey(id uint32) (result DictSynonyms, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictSynonyms{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByWord  获取多个内容
func (obj *_DictSynonymsMgr) FetchIndexByWord(isDel bool) (results []*DictSynonyms, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictSynonyms{}).Where("`is_del` = ?", isDel).Find(&results).Error

	return
}
