package createtable

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

// User 创建user结构体对应的数据库表
type User struct {
	ID           uint    `gorm:"primarykey"`
	Name         string  //默认值传入数据库是空串""
	Email        *string //默认值传入数据库是null
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString //默认值传入数据库是null
	ActivatedAt  sql.NullTime   //默认值传入数据库是null
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ignored      string
}

// Model 抽象出公共结构体
type Model struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// User 将gorm中设置的Model嵌入到User中使用
//type User struct {
//	gorm.Model
//	Name string
//}

// Author 正常结构体内嵌
type Author struct {
	Name  string
	Email string
}

// Blog 正常内嵌结构体
type Blog struct {
	ID      int    `gorm:"primary"`
	Author  Author `gorm:"embedded"`
	upVotes int32
}

// Run 执行Run方法创建user表 约定将结构名转换为复数即加s
func Run(db *gorm.DB) {
	err := db.AutoMigrate(&User{})
	if err != nil {
		return
	}
	user := &User{}
	user.MemberNumber.Valid = true //设置true 此时sql.NullString传入的值为""
	err = db.Create(user).Error
	//创建blog表
	//err := db.AutoMigrate(&Blog{})
	//if err != nil {
	//	return
	//}
	//blog := &Blog{}
	//err = db.Create(blog).Error
	//if err != nil {
	//	return
	//}
}
