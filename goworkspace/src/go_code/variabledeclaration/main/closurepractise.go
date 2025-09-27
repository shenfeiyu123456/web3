package main

/**
闭包也称为匿名函数，没有函数名，通常在函数内定义，或者作为参数、返回值进行传递，
匿名函数的优势是可以直接使用当前函数内在匿名函数声明之前声明的变量
*/

import "fmt"

type A struct {
	i int
}

func (a *A) add(v int) int {
	return a.i + v
}

// 声明函数变量
var func1 func(int) int

// 声明闭包
var closure = func(p int) int {
	p *= p
	return p
}

func main() {
	a := A{1}
	//把方法赋值给函数变量
	func1 = a.add
	//声明一个闭包并直接运行,入参和出参都是两个数
	//() 的作用‌：
	//外层匿名函数 func() {...} 定义后，通过 () 立即执行它，相当于调用了这个函数。
	//执行后，外层函数返回内层的闭包函数（即 func(i int, s string) (int, string)），并将其赋值给 returnFun
	//如果去掉 ()，returnFun 会变成外层匿名函数本身（未被调用），而非内层的闭包函数。
	returnFun := func() func(int, string) (int, string) {
		fmt.Println("这是一个匿名函数内部")
		return func(i int, s string) (int, string) {
			return i, s
		}
	}()
	//执行匿名函数
	res1, res2 := returnFun(1, "test")
	fmt.Println(res1, res2)
	fmt.Println(a.i)
	//调用函数
	fmt.Println(func1(1))
	fmt.Println(a.i)
	//调用函数
	fmt.Println(closure(10))
}
