package main

import (
	"fmt"
	"unsafe"
)

/*
*
基础的简单数据类型包括字符串、布尔、数字
*/
func main() {
	//1、字符串string 类型 使用""或``声明，
	var name = "Alice"
	fmt.Println("name", name)
	fmt.Println("name占的实际内存", unsafe.Sizeof(name))
	//2、布尔类型，只有true或false
	var flag = true
	fmt.Println("flag", flag)
	fmt.Println("name占的实际内存", unsafe.Sizeof(flag))
	//3、数字类型
	//3.1 整数
	//3.1.1 正负整数类型 默认 int
	//int8 int16 int32 int64
	//int8 占位1个byte 8个字节bit
	//int16 占位2个byte 16个字节bite
	//int32 占位4个byte 32个字节bite
	//int64 占位8个byte 64个字节bite
	//int 占位和操作系统有关系，32位操作系统就是int32 占4个byte;64位操作系统就是int64 占8个byte,可以使用unsafe.Sizeof()判断使用内存
	var age int8 = 25
	fmt.Println("age", age)
	fmt.Println("age占的实际内存", unsafe.Sizeof(age))
	var age1 = 32
	fmt.Println("age1", age1)
	fmt.Println("age1占的实际内存", unsafe.Sizeof(age1))
	//3.1.2 纯非负整数类型
	//uint8 uint16 uint32 uint64
	var age2 uint8 = 21
	fmt.Println("age2", age2)
	fmt.Println("age2占的实际内存", unsafe.Sizeof(age2))
	var age4 uint = 21
	fmt.Println("age4", age4)
	fmt.Println("age4占的实际内存", unsafe.Sizeof(age4))
	//3.2浮点型包括float32 float64 默认 float64
	var paiVal float32 = 3.141592
	fmt.Println("paiVal", paiVal)
	fmt.Println("paiVal", unsafe.Sizeof(paiVal))
	var sumMoney = 11.05
	fmt.Println("sumMoney", unsafe.Sizeof(sumMoney))
	fmt.Println("sumMoney", sumMoney)
	//4.常量的声明
	const address = "郑州"
	fmt.Println("address", address)
	const email, youbian string = "1023", "111111"
	fmt.Println("email,youbian", email, youbian)
	const nameDoughter, ageDoughter = "申诗怡", 7
	fmt.Println("nameDoughter,ageDoughter", nameDoughter, ageDoughter)
}
