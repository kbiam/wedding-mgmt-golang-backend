package db

import (

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB


func Connect() {
	dsn := "root:SJYSXCxrkJQhDgUPYaAExUjolRhChoHq@tcp(centerbeam.proxy.rlwy.net:51871)/wedding_management_golang?parseTime=true"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil{
		panic("failed to connect database")
	}
}