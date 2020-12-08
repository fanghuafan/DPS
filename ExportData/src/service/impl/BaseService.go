package impl

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

// 类定义
type BaseService struct {
}

// 获取初始化数据库对象
func (baseService *BaseService) GetDB() *gorm.DB {
	db, err := gorm.Open("mysql",
		"root:root@(localhost)/export?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	sdb := db.DB()
	err = sdb.Ping()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return db
}

func (baseService *BaseService) CloseDB(db *sql.DB) {
	db.Close()
}
