package main

import (
	"fmt"
	"time"
)

/*
*
在 go 关键字后面加一个函数，就可以创建一个线程，函数可以为已经写好的函数，也可以是匿名函数。
*/
func main() {
	fmt.Println("main函数执行开始")
	//go func() {
	//	fmt.Println("goroutine")
	//}()
	/**
	为什么没有输出 goroutine ？
	首先，我们清楚 Go 语言的线程是并发机制，不是并行机制。
	那么，什么是并发，什么是并行？
	并发是不同的代码块交替执行，也就是交替可以做不同的事情。
	并行是不同的代码块同时执行，也就是同时可以做不同的事情。
	举个生活化场景的例子：
	你正在家看书，忽然电话来了，然后你接电话，通话完成后继续看书，这就是并发，看书和接电话交替做。
	如果电话来了，你一边看书一遍接电话，这就是并行，看书和接电话一起做。
	说回上面的例子，为什么没有输出 goroutine ？
	main 函数是一个主线程，是因为主线程执行太快了，子线程还没来得及执行，所以看不到输出。
	现在让主线程休眠 1 秒钟，再试试。
	*/
	//time.Sleep(1 * time.Second)
	//fmt.Println("main函数执行结束")
	//声明channal,遵循先进先出的规则
	//声明不带缓冲的通道
	//ch := make(chan string)
	//ch := make(chan string,1)
	//ch <- "a" //写入 "a"
	////阻塞的通道是不能走到这里的
	//fmt.Println("走这里了")
	//go func() {
	//	val := <-ch //出 "ch"
	//	fmt.Println("val", val)
	//}()
	//time.Sleep(1 * time.Second)
	//fmt.Println("main方法结束")
	//如果启动两个线程一个进行写入，一个进行输出，此时即使使用阻塞渠道，也没有问题，因为已经输入的同时已经输出，渠道是空的
	//ch := make(chan string)
	////写入线程
	//go func() {
	//	//写入字母”a“
	//	ch <- "a"
	//	fmt.Println("写入成功", ch)
	//}()
	////写出线程
	//go func() {
	//	val := <-ch
	//	fmt.Println("写出成功", val)
	//}()
	//time.Sleep(1 * time.Second)
	//fmt.Println("main函数结束")
	//ch := make(chan string, 3)
	////启动线程
	//go producer(ch)
	//time.Sleep(1 * time.Second)
	//fmt.Println("main函数结束")
	//创建渠道，没有缓冲，因为此时有消费者处理信息，不会影响阻塞
	ch := make(chan string)
	go producer(ch)
	go consumer(ch)
	time.Sleep(1 * time.Second)
	fmt.Println("main函数结束")
}

// 生产者队列
// 字符串队列
// 渠道阻塞，可以使用缓冲channel，如果此时缓冲channel的长度小于输入的长度，此时渠道依旧会阻塞
// 消除渠道阻塞：1、缓冲channel长度足够大，2、增加输出渠道，将信息及时输出即可
func producer(ch chan string) {
	fmt.Println("producer start	")
	ch <- "a"
	ch <- "b"
	ch <- "c"
	ch <- "d"
	fmt.Println("producer end	")
}
func consumer(ch chan string) {
	fmt.Println("consumer start	")
	//此步骤使用死循环，无限制的将生产者的信息进行消费
	for {
		msg := <-ch
		fmt.Println("consumer", msg)
	}
}
