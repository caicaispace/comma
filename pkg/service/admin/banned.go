package admin

import (
	"errors"

	"comma/pkg/library/db"
	"comma/pkg/library/util"
	"comma/pkg/library/util/business"

	service "comma/pkg/service/banned"

	"gorm.io/gorm"
)

type Banned struct {
	ID         int `gorm:"id" json:"id"`
	WordId     int `gorm:"word_id" json:"word_id"`
	ProjectId  int `gorm:"project_id" json:"project_id"`
	IsDel      int `gorm:"is_del" json:"is_del"`
	CreateTime int `gorm:"create_time" json:"create_time"`
	UpdateTime int `gorm:"update_time" json:"update_time"`
}

type BannedList struct {
	Banned
	Word        string `gorm:"word" json:"word"`
	ProjectName string `gorm:"project_name" json:"project_name"`
}

type bannedService struct{}

const bannedTableName = "dict_banned"

func NewBanned() *bannedService {
	return &bannedService{}
}

func (bs *bannedService) BannedGetList(pager *business.Pager, filter *Word) ([]BannedList, int64) {
	list := make([]BannedList, 0)
	table := db.DB().Table(bannedTableName)
	total := int64(0)
	table.Select("count(*)").Count(&total)
	pager.SetTotal(int(total))
	table.Select(`
dict_banned.id AS id,
dict_banned.word_id AS word_id,
dict_word.word AS word,
dict_banned.project_id AS project_id,
dict_project.name AS project_name,
dict_banned.is_del as is_del,
dict_banned.create_time as create_time,
dict_banned.update_time AS update_time
`)
	table.Joins("left join dict_word ON dict_banned.word_id = dict_word.id")
	table.Joins("left join dict_project ON dict_banned.project_id = dict_project.id")
	if (Word{} != *filter) {
		if filter.Word != "" {
			table.Where("dict_word.word LIKE ?", filter.Word+"%")
		}
		if filter.IsDel > 0 {
			table.Where("dict_banned.is_del", filter.IsDel)
		}
	}
	table.Order("dict_banned.id DESC")
	table.Limit(pager.GetLimit()).Offset(pager.GetOffset())
	table.Find(&list)
	return list, total
}

type BannedCreateForm struct {
	WordId    int     `json:"word_id"`
	Weight    float32 `json:"weight"`
	ProjectId int     `json:"project_id"`
}

func (bs *bannedService) BannedCreate(inData BannedCreateForm) (*Banned, error) {
	outData := Banned{
		WordId:     inData.WordId,
		ProjectId:  inData.ProjectId,
		IsDel:      1,
		CreateTime: int(util.NowTimestamp()),
	}
	err := db.DB().Transaction(func(tx *gorm.DB) error {
		ret := tx.Table(bannedTableName).Create(&outData)
		if ret.Error != nil {
			return ret.Error
		}
		// word, err := NewWord().WordGetInfoById(inData.WordId)
		// if err != nil {
		// 	return err
		// }
		// service.GetInstance().Add(word.Word)
		// 返回 nil 提交事务
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &outData, nil
}

type BannedUpdateForm struct {
	WordId     int `json:"word_id"`
	ProjectId  int `json:"project_id"`
	IsDel      int `json:"is_del"`
	UpdateTime int `json:"update_time"`
}

func (bs *bannedService) BannedUpdateById(id int, inData BannedUpdateForm) error {
	updateData, _ := util.StructToMap(inData, "json")
	updateData["update_time"] = int(util.NowTimestamp())
	ret := db.DB().Table(bannedTableName).Where("id = ?", id).Updates(updateData)
	if ret.Error != nil {
		return ret.Error
	}
	// var err error
	// var word Word
	// word, err = NewWord().WordGetInfoById(inData.WordId)
	// if err != nil {
	// 	return err
	// }
	// service.GetInstance().Del(word.Word)
	// word, err = NewWord().WordGetInfoById(inData.WordId)
	// if err != nil {
	// 	return err
	// }
	// service.GetInstance().Add(word.Word)
	// if ret.RowsAffected <= 0 {
	// 	return errors.New("update error")
	// }
	return nil
}

type BannedMultipleDeleteForm struct {
	Ids []int `json:"ids"`
}

func (bs *bannedService) BannedDeleteByIds(inData BannedMultipleDeleteForm) error {
	ret := db.DB().Table(bannedTableName).Where("id", inData.Ids).Update("is_del", 1)
	if ret.Error != nil {
		return ret.Error
	}
	word, err := NewWord().WordGetInfoByIds(inData.Ids)
	if err != nil {
		return err
	}
	service.GetInstance().Del(word.Word)
	if ret.RowsAffected <= 0 {
		return errors.New("delete error")
	}
	return nil
}
