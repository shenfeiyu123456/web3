package main

import (
	"encoding/json"
	"fmt"
)

/*
*
Map集合是无序的key-value键值对
Map中的key/value可以是任意类型，但是所有的key都必须是统一类型，所有的value也必须是统一类型，key和value可以不是统一中类型
*/
func main() {
	//第一种声明方式
	var p1 map[int]string
	p1 = make(map[int]string)
	p1[1] = "hello"
	fmt.Println("p1", p1)
	//第二种声明方式
	var p2 map[int]string = map[int]string{}
	p2[1] = "tom"
	fmt.Println("p2", p2)
	//第三种声明方式
	var p3 map[int]string = make(map[int]string)
	p3[1] = "jerry"
	fmt.Println("p3", p3)
	//第四种声明方式
	p4 := make(map[int]string)
	p4[1] = "jack"
	fmt.Println("p4", p4)
	//第五种声明方式
	p5 := map[int]string{}
	p5[1] = "john"
	fmt.Println("p5", p5)
	//第六种声明方式
	p6 := map[int]string{
		1: "fisher",
	}
	fmt.Println("p6", p6)
	//JSON化Map数据信息
	result := make(map[string]interface{})
	result["code"] = 0
	result["message"] = "success"
	result["data"] = map[string]interface{}{
		"name":  "张三",
		"age":   30,
		"hobby": []string{"爬山", "打篮球"},
	}
	fmt.Println("result", result)
	//JSON序列化
	jsons, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsons))
	//根据key获取val
	fmt.Println(result["code"])
	fmt.Println(result["message"])
	fmt.Println(result["data"])
}
