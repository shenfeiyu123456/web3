package main

/**
结构体是将零个或多个任意类型的变量，组合在一起的聚合数据类型，也可以看作数据的集合
结构体的属性名称可以是大写或小写，但是建议大写，因为JSON序列化一定得是大写
*/
/**
实体类型转JSON
实体类型的属性名称必须是大写 Age,Name
*/
import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Age  int
	Name string
}

type Result struct {
	Code int
	Msg  string
}

// 对类型 Result 设置 set/print/方法
func setter(result *Result) {
	//*是指针类型，获取地址进行内容修改，而非副本修改
	result.Code = 0
	result.Msg = "ok"
	fmt.Println("setter", result)
}
func printJSON(result *Result) {
	jsos, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("printJSON", string(jsos))
}

func main() {
	//1.第一种初始化方式
	var person Person
	person.Age = 23
	person.Name = "Jack"
	fmt.Println(person)
	//第二种初始化方式
	var person2 = Person{Age: 23, Name: "Jack"}
	fmt.Println(person2)
	//第三种初始化方式
	person3 := Person{Age: 23, Name: "Jack"}
	fmt.Println(person3)
	//匿名结构体
	person4 := struct {
		Age  int
		Name string
	}{Age: 23, Name: "Jack"}
	fmt.Println(person4)
	result := Result{Code: 200, Msg: "success"}
	fmt.Println(result)
	//将结果序列话
	jsons, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsons))
	result1 := Result{Code: 200, Msg: "success"}
	//重新设置
	setter(&result1)
	//JSON序列化打印
	printJSON(&result1)
}
