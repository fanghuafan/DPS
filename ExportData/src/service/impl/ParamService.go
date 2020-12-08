package impl

import (
	"ExportData/src/models"
	"log"
)

type ParamService struct {
	BaseService
}

func (paramService *ParamService) Get(templateId int32) []models.Param {
	db := paramService.BaseService.GetDB()

	list := []models.Param{}
	//调用原生sql语句
	db.Raw("SELECT * FROM t_param WHERE template_id=?", templateId).Find(&list)
	log.Println("query export param data by template_id: ", list)
	defer paramService.BaseService.CloseDB(db.DB())

	return list
}

func (paramService *ParamService) Add(param *models.Param) bool {
	db := paramService.BaseService.GetDB()

	db.Create(param)
	defer paramService.BaseService.CloseDB(db.DB())

	return db.NewRecord(param)
}
