package inserttable

import (
	"fmt"
	"goworkspace/gorm/src/createtable"
	"time"

	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	var pt *time.Time
	t := time.Now()
	pt = &t
	var email *string
	str := "1023858367@qq.ocm"
	email = &str
	//单条信息插入
	//user := createtable.User{Name: "张三", Age: 18, Birthday: pt, Email: email}
	//result := db.Create(&user)
	//fmt.Println(result.RowsAffected)
	//批量信息插入
	users := []*createtable.User{
		{Name: "李四", Age: 18, Birthday: pt, Email: email},
		{Name: "王五", Age: 18, Birthday: pt, Email: email},
		{Name: "赵六", Age: 18, Birthday: pt, Email: email},
	}
	result := db.Create(&users)
	fmt.Println(result.RowsAffected)
}
