package db

import (
     "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func DbOpen(){
	db, err := gorm.Open("mysql", "root:123456@/dbname?charset=utf8")
	if err !=nil{
		
	}
	defer db.Close()
}