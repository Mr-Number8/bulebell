package models

import "github.com/jinzhu/gorm"

const (
	userName = "zxq"
	password = "zz123456*"
	host     = "(127.0.0.1:3306)"
	db       = "db02"
	setting  = "charset=utf8mb4&parseTime=True&loc=Local"
)
var (
	DB *gorm.DB
)
func InitMsql(){
	dsn  	 := userName+":"+password+ "@tcp"+host+"/"+db+"?"+setting
	DB,_=gorm.Open("mysql",dsn)
	DB.AutoMigrate(&Todo{}) //绑定模型
}
func CloseDB()error{
	return DB.Close()
}
