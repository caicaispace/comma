package admin

import (
	"comma/pkg/library/db"
	"comma/pkg/library/util"
	"errors"

	"github.com/caicaispace/gohelper/business"
)

type Project struct {
	ID         int    `gorm:"id" json:"id"`
	Name       string `gorm:"name" json:"name"`
	IsDel      int    `gorm:"is_del" json:"is_del"`
	CreateTime int    `gorm:"create_time" json:"create_time"`
	UpdateTime int    `gorm:"update_time" json:"update_time"`
}

type ProjectList struct {
	Project
}

const projectTableName = "dict_project"

type projectService struct{}

func NewProject() *projectService {
	return &projectService{}
}

func (ps *projectService) ProjectGetList(pager *business.Pager) ([]ProjectList, int64) {
	list := make([]ProjectList, 0)
	total := int64(0)
	table := db.DB().Table(projectTableName)
	table.Select("count(*)").Count(&total)
	pager.SetTotal(int(total))
	table.Select("*")
	table.Where("is_del", 0)
	table.Order("id DESC")
	table.Limit(pager.GetLimit()).Offset(pager.GetOffset())
	table.Find(&list)
	return list, total
}

func (ps *projectService) ProjectGetInfoById(id int) (*Project, error) {
	modelOut := &Project{}
	model := db.DB().Table(projectTableName)
	ret := model.Where("id", id).First(modelOut)
	if ret.Error != nil {
		return modelOut, ret.Error
	}
	return modelOut, nil
}

type ProjectCreateForm struct {
	Name  string `json:"name"`
	IsDel int    `json:"is_del"`
}

func (ps *projectService) ProjectCreate(inData ProjectCreateForm) (*Project, error) {
	outData := Project{
		Name:       inData.Name,
		IsDel:      inData.IsDel,
		CreateTime: int(util.NowTimestamp()),
	}
	ret := db.DB().Table(projectTableName).Create(&outData)
	if ret.Error != nil {
		return nil, ret.Error
	}
	return &outData, nil
}

type ProjectUpdateForm struct {
	Name       string `json:"name"`
	IsDel      int    `json:"is_del"`
	UpdateTime int    `json:"update_time"`
}

func (ps *projectService) ProjectUpdateById(id int, inData ProjectUpdateForm) (*ProjectUpdateForm, error) {
	updateData, _ := util.StructToMap(inData, "json")
	updateData["update_time"] = int(util.NowTimestamp())
	ret := db.DB().Table(projectTableName).Where("id = ?", id).Updates(updateData)
	if ret.Error != nil {
		return nil, ret.Error
	}
	if ret.RowsAffected <= 0 {
		return nil, errors.New("update err")
	}
	err := util.MapToStruct(updateData, &inData)
	if err != nil {
		return nil, err
	}
	return &inData, nil
}

type ProjectMultipleDeleteForm struct {
	Ids []int `json:"ids"`
}

func (ps *projectService) ProjectDeleteByIds(inData ProjectMultipleDeleteForm) bool {
	ret := db.DB().Table(projectTableName).Where("id", inData.Ids).Update("is_del", 1)
	if ret.Error != nil {
		return false
	}
	return ret.RowsAffected > 0
}
