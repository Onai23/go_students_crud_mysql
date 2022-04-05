package config

import "github.com/jinzhu/gorm"

var (
	db *gorm.DB
)

//set fungsi connect
func Connect() {
	d, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/db_go_mysql?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
