package main

import (
	"fmt"
	"time"
)

/**
go 支持并发的方式，就是通过gorutine和channel提供的简洁且高效的方式实现的
*/

/**
gorutine是轻量级线程，创建一个gorutine所需的资源开销较小，所以可以创建许多的gorutine进行并发工作
他们是有Go运行时调度的，调度过程就是Go运行时把gorutine任务分配给CPU执行的过程
但是gorutine不是通常意义上的线程，线程都是由操作系统进行调度的
在GO中，想让某个任务并发或者异步执行，只需要把任务封装成一个函数或闭包即可，交给gorutine执行即可
*/

func sayHello(str string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}

}

func main() {
	//声明方式1把方法或者函数交给gorutain执行
	go sayHello("nihao")
	//声明方式2把闭包交给gorutain执行
	go func() {
		fmt.Println("good morning")
	}()
	//声明方式3把闭包交给gorutain执行
	go func(string2 string) {
		fmt.Println(string2)
	}("hello world")
	//只有方法调用，gorutain执行才有意义
	sayHello("hello")
}
