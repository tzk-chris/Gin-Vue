package common

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/tzk-chris/Gin-Vue/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := ""
	port := ""
	database := "gingooo"
	username := "root"
	password := ""
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}
	db.AutoMigrate(&model.User{}) // 自动创建数据表

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
