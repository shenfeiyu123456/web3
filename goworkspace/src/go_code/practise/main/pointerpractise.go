package main

/**
指针是一个变量，它存储了另一个变量的内存地址，通过指针可以访问存储在指定内存地址中的变量
*/
import "fmt"

// 在Go中，声明一个指针类型变量需使用‘*’标识,&用于取地址,两者相互搭配进行传参
var p1 *int
var p2 *string

func main() {
	i := 1
	s := "hello"
	p1 = &i
	p2 = &s
	fmt.Println(*p1, *p2)

}
