package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _DictFestivalMgr struct {
	*_BaseMgr
}

// DictFestivalMgr open func
func DictFestivalMgr(db *gorm.DB) *_DictFestivalMgr {
	if db == nil {
		panic(fmt.Errorf("DictFestivalMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DictFestivalMgr{_BaseMgr: &_BaseMgr{DB: db.Table("dict_festival"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DictFestivalMgr) GetTableName() string {
	return "dict_festival"
}

// Reset 重置gorm会话
func (obj *_DictFestivalMgr) Reset() *_DictFestivalMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_DictFestivalMgr) Get() (result DictFestival, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictFestival{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DictFestivalMgr) Gets() (results []*DictFestival, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictFestival{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_DictFestivalMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(DictFestival{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_DictFestivalMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithWordID word_id获取
func (obj *_DictFestivalMgr) WithWordID(wordID uint32) Option {
	return optionFunc(func(o *options) { o.query["word_id"] = wordID })
}

// WithProjectID project_id获取
func (obj *_DictFestivalMgr) WithProjectID(projectID uint32) Option {
	return optionFunc(func(o *options) { o.query["project_id"] = projectID })
}

// WithName name获取 名称
func (obj *_DictFestivalMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithSunDate sun_date获取 阳历
func (obj *_DictFestivalMgr) WithSunDate(sunDate string) Option {
	return optionFunc(func(o *options) { o.query["sun_date"] = sunDate })
}

// WithLunarDate lunar_date获取 阴历
func (obj *_DictFestivalMgr) WithLunarDate(lunarDate string) Option {
	return optionFunc(func(o *options) { o.query["lunar_date"] = lunarDate })
}

// WithIsDel is_del获取
func (obj *_DictFestivalMgr) WithIsDel(isDel bool) Option {
	return optionFunc(func(o *options) { o.query["is_del"] = isDel })
}

// WithCreateTime create_time获取
func (obj *_DictFestivalMgr) WithCreateTime(createTime uint32) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_DictFestivalMgr) WithUpdateTime(updateTime uint32) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_DictFestivalMgr) GetByOption(opts ...Option) (result DictFestival, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DictFestival{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DictFestivalMgr) GetByOptions(opts ...Option) (results []*DictFestival, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DictFestival{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_DictFestivalMgr) GetFromID(id uint32) (result DictFestival, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictFestival{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_DictFestivalMgr) GetBatchFromID(ids []uint32) (results []*DictFestival, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictFestival{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromWordID 通过word_id获取内容
func (obj *_DictFestivalMgr) GetFromWordID(wordID uint32) (results []*DictFestival, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictFestival{}).Where("`word_id` = ?", wordID).Find(&results).Error

	return
}

// GetBatchFromWordID 批量查找
func (obj *_DictFestivalMgr) GetBatchFromWordID(wordIDs []uint32) (results []*DictFestival, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictFestival{}).Where("`word_id` IN (?)", wordIDs).Find(&results).Error

	return
}

// GetFromProjectID 通过project_id获取内容
func (obj *_DictFestivalMgr) GetFromProjectID(projectID uint32) (results []*DictFestival, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictFestival{}).Where("`project_id` = ?", projectID).Find(&results).Error

	return
}

// GetBatchFromProjectID 批量查找
func (obj *_DictFestivalMgr) GetBatchFromProjectID(projectIDs []uint32) (results []*DictFestival, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictFestival{}).Where("`project_id` IN (?)", projectIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *_DictFestivalMgr) GetFromName(name string) (results []*DictFestival, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictFestival{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 名称
func (obj *_DictFestivalMgr) GetBatchFromName(names []string) (results []*DictFestival, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictFestival{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromSunDate 通过sun_date获取内容 阳历
func (obj *_DictFestivalMgr) GetFromSunDate(sunDate string) (results []*DictFestival, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictFestival{}).Where("`sun_date` = ?", sunDate).Find(&results).Error

	return
}

// GetBatchFromSunDate 批量查找 阳历
func (obj *_DictFestivalMgr) GetBatchFromSunDate(sunDates []string) (results []*DictFestival, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictFestival{}).Where("`sun_date` IN (?)", sunDates).Find(&results).Error

	return
}

// GetFromLunarDate 通过lunar_date获取内容 阴历
func (obj *_DictFestivalMgr) GetFromLunarDate(lunarDate string) (results []*DictFestival, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictFestival{}).Where("`lunar_date` = ?", lunarDate).Find(&results).Error

	return
}

// GetBatchFromLunarDate 批量查找 阴历
func (obj *_DictFestivalMgr) GetBatchFromLunarDate(lunarDates []string) (results []*DictFestival, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictFestival{}).Where("`lunar_date` IN (?)", lunarDates).Find(&results).Error

	return
}

// GetFromIsDel 通过is_del获取内容
func (obj *_DictFestivalMgr) GetFromIsDel(isDel bool) (results []*DictFestival, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictFestival{}).Where("`is_del` = ?", isDel).Find(&results).Error

	return
}

// GetBatchFromIsDel 批量查找
func (obj *_DictFestivalMgr) GetBatchFromIsDel(isDels []bool) (results []*DictFestival, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictFestival{}).Where("`is_del` IN (?)", isDels).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_DictFestivalMgr) GetFromCreateTime(createTime uint32) (results []*DictFestival, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictFestival{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_DictFestivalMgr) GetBatchFromCreateTime(createTimes []uint32) (results []*DictFestival, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictFestival{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_DictFestivalMgr) GetFromUpdateTime(updateTime uint32) (results []*DictFestival, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictFestival{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_DictFestivalMgr) GetBatchFromUpdateTime(updateTimes []uint32) (results []*DictFestival, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictFestival{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_DictFestivalMgr) FetchByPrimaryKey(id uint32) (result DictFestival, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DictFestival{}).Where("`id` = ?", id).First(&result).Error

	return
}
