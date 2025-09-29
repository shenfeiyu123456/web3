package updatetable

import (
	"fmt"
	"goworkspace/gorm/src/createtable"

	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	//user := createtable.User{}
	//db.First(&user)
	//user.Name = "申飞宇"
	//user.Age = 36
	// 1、根据id修改
	//UPDATE `users` SET `name`='申飞宇',`email`=NULL,`age`=36,`birthday`=NULL,`member_number`='',`activated_at`=NULL,`created_at`='2025-09-29 09:52:17.675',`updated_at`='2025-09-29 17:37:31.866' WHERE `id` = 1
	//result := db.Debug().Save(&user)
	// 2、条件更新
	//UPDATE `users` SET `age`=18,`updated_at`='2025-09-29 17:42:25.437' WHERE `name` = '申飞宇'
	result := db.Debug().Model(&createtable.User{}).Where("name", "申飞宇").Update("age", 18)
	fmt.Println(result.RowsAffected)
}
