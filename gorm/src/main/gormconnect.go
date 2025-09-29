package main

import (
	"fmt"
	"goworkspace/gorm/src/deletetable"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/web3?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("Successfully connected to MySQL with GORM")
	//创建user表
	//createtable.Run(db)
	//插入user表数据
	//inserttable.Run(db)
	//查询数据
	//querytable.Run(db)
	//修改数据
	//updatetable.Run(db)
	//删除数据
	deletetable.Run(db)
}
