package querytable

import (
	"fmt"
	"goworkspace/gorm/src/createtable"

	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	var users []createtable.User
	//查询全部对象
	//db.Debug().Find(&users)
	//for _, user := range users {
	//	fmt.Println(user.Name)
	//}
	//user := createtable.User{}
	// 自动带排序关键字orderby
	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1
	//db.Debug().First(&user)
	//用主键(int)搜索
	//SELECT * FROM `users` WHERE `users`.`id` = 10
	//db.Debug().Find(&user, 10)
	//SELECT * FROM `users` WHERE `users`.`id` IN (1,2,3)
	//db.Debug().Find(&user, []int{1, 2, 3})
	//如果主键是UUID字符串
	//db.Debug().Find(&user, "id=?", "1123131231")
	//1、条件查询
	//SELECT * FROM `users` WHERE Name='张三'
	//db.Debug().Find(&user, "Name=?", "张三")
	//SELECT * FROM `users` WHERE `users`.`name` = '张三' AND `users`.`age` = 18
	//db.Debug().Find(&user, &createtable.User{Name: "张三", Age: 18})
	//SELECT * FROM `users` WHERE `users`.`age` = 18 AND `users`.`name` = '张三'
	//db.Debug().Find(&user, map[string]interface{}{"name": "张三", "age": 18})
	//2、条件查询
	//SELECT * FROM `users` WHERE Name = '张三'
	//db.Debug().Where("Name = ?", "张三").Find(&user)
	//多条件查询 实体对象
	//SELECT * FROM `users` WHERE `users`.`name` = '李四' AND `users`.`age` = 18
	//db.Debug().Where(&createtable.User{Name: "李四", Age: 18}).Find(&user)
	//多条件查询 Map结构
	//SELECT * FROM `users` WHERE `users`.`age` = 18 AND `users`.`name` = '李四'
	//db.Debug().Where(map[string]interface{}{"name": "李四", "age": 18}).Find(&user)
	//3、NOT查询
	//SELECT * FROM `users` WHERE NOT name='张三'
	//db.Debug().Not("name=?", "张三").Find(&users)
	//SELECT * FROM `users` WHERE `id` NOT IN (1,2,3)
	//db.Debug().Not("id", []int{1, 2, 3}).Find(&users)
	//4、OR查询
	//SELECT * FROM `users` WHERE `name` = '张三' OR `age` = 18
	//db.Debug().Where("name", "张三").Or("age", 18).Find(&users)
	// 5、特定字段查询返回
	//SELECT `name`,`age` FROM `users`
	//db.Debug().Select("name", "age").Find(&users)
	// 6、ORDER查询
	//SELECT * FROM `users` ORDER BY id asc
	//db.Debug().Order("id asc").Find(&users)
	fmt.Println(users)

}
