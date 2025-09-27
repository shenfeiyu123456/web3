package main

/**
数组是一个由固定长度的特定类型元素组成的序列，一个数组中可以由零个或多个元素组成，一旦声明了，数组长度就固定了，不能动态变化
len()和cap()返回的结果是一样的
*/
import (
	"fmt"
)

func main() {
	//一维数组
	//声明方式1、只声明长度
	var arr1 [5]int
	fmt.Println("arr_1", arr1)
	//声明方式2、声明长度的同时初始化元素
	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr_2", arr2)
	fmt.Println(arr2[2])
	//声明方式3、不使用var声明
	arr3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr_3", arr3)
	//声明方式4、[...]中的...表示让编译器自动计算数组的长度
	arr4 := [...]int{1, 2, 3, 4, 5}
	fmt.Println("arr_4", arr4)
	//声明方式5、指定元素位置进行初始化
	arr5 := [...]int{0: 3, 1: 5, 4: 6}
	fmt.Println("arr_5", arr5)
	//循环遍历
	for i := 0; i < len(arr5); i++ {
		fmt.Println(arr5[i])
	}
	//二维数组
	var arr6 = [3][5]int{{1, 2, 3, 4, 5}, {5, 4, 3, 2, 1}, {1, 1, 1, 1, 1}}
	fmt.Println("arr_6", arr6)
	arr7 := [3][5]int{{1, 2, 3, 4, 5}, {5, 4, 3, 2, 1}, {1, 1, 1, 1, 1}}
	fmt.Println("arr_7", arr7)
	arr8 := [...][5]int{{1, 2, 3, 4, 5}, {0: 3, 1: 4, 2: 5}, {1, 2, 3, 4, 5}}
	fmt.Println("arr_8", arr8)
	//person3 := Person{Age: 23, Name: "Jack"}
	//fmt.Println("aaaaa", person3)
}
