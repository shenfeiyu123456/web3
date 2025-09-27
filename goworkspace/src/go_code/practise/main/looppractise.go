package main

import "fmt"

func main() {
	//数组
	arr := [5]int{1, 2, 3, 4, 5}
	//切片
	slice := []int{6, 7, 8}
	//Map
	map1 := make(map[int]string)
	map1[0] = "张三"
	map1[1] = "郑州"
	map1[2] = "中国"
	//fmt.Println(arr)
	//fmt.Println(slice)
	//fmt.Println(map1)
	//循环三种数据结构
	//1、利用下标进行循环
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	for i := 0; i < len(map1); i++ {
		fmt.Println(map1[i])
	}
	fmt.Println("第二种循环开始")
	//2、使用kv键值对进行循环 range关键字
	for k, v := range arr {
		//下标和值
		fmt.Println(k, v)
	}
	for k, v := range slice {
		fmt.Println(k, v)
	}
	for k, v := range map1 {
		fmt.Println(k, v)
	}
	//Map结构如果键不是int类型，for按照坐标循环是会报错的
	map2 := make(map[string]interface{})
	map2["name"] = "张三"
	map2["age"] = 30
	map2["hobby"] = []string{"打篮球", "踢足球", "乒乓球"}
	for k, v := range map2 {
		fmt.Println(k, v)
	}
	fmt.Println("使用第三种循环方式")
	//使用空白符进行循环便利
	for _, v := range arr {
		fmt.Println(v)
	}
	for _, v := range slice {
		fmt.Println(v)
	}
	for _, v := range map2 {
		fmt.Println(v)
	}
	for k, _ := range map2 {
		fmt.Println(k)
	}

}
