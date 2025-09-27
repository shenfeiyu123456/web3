package main

import "fmt"

/**
与函数相比，方法是一个包含接收者的函数，大部分可以通过类型的实例调用，也可以将方法赋值给一个函数变量
使用函数变量调用这个方法，调用方法类似闭包
*/

// B 声明结构体
type B struct {
	i int
}

// add 声明方法 func + (接受体[指针类型])+方法名+(入参)+返回值+{方法体}
// 这个接受体就是指明了是它的方法
func (B *B) add(v int) int {
	return B.i + v
}

// 声明函数
func add1(B *B, b int) int {
	return B.i + b
}

func main() {
	B := B{1}
	// 不用使用&的指针获取方式，直接调用
	fmt.Println(B.add(2))
	//函数中传入指针，并且使用& 可以替代方法的使用
	fmt.Println(add1(&B, 2))
}
