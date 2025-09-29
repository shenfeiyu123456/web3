package deletetable

import (
	"fmt"
	"goworkspace/gorm/src/createtable"

	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	user := createtable.User{}
	//db.First(&user)
	//1、查询后删除
	//DELETE FROM `users` WHERE `users`.`id` = 1
	//result := db.Debug().Delete(&user)
	//2、条件删除
	//DELETE FROM `users` WHERE `name` = '李四'
	//result := db.Debug().Where("name", "李四").Delete(&user)
	//3.多个条件删除
	//DELETE FROM `users` WHERE `users`.`age` = 18 AND `users`.`name` = '王五'
	//result := db.Debug().Where(map[string]interface{}{"name": "王五", "age": 18}).Delete(&user)
	//DELETE FROM `users` WHERE `name` = ''
	result := db.Debug().Where("name", "").Delete(&user)
	fmt.Println(result.RowsAffected)

}
