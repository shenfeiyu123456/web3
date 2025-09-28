package main

/**
channel是GO语言中定义的一种类型，专门用来在多个goroutine之间通信的线程安全的数据结构，可以
在一个goroutine中向一个channel中发送数据，从另一个channel中接收数据，channel类似队列，
满足先进先出的原则,channel有三种操作，发送数据，接收数据，关闭通道
在GO语言中，当需要goroutine之间协作地方，更常见的是使用channel,而不是sync包中的Mutex或RWMutex的互斥锁
但是他们各有侧重点，大部分时候流程是根据数据驱动的，channel会被使用的更频繁
channel擅长的领域是数据流动的场景
1、传递数据权，把某个数据发送给其他协程
2、分发任务，每个任务都是一个数据
3、交流异步结果，结果是一个数据
而锁使用的场景更偏向同一个协程访问数据的权限
1、访问缓存
2、状态管理
结论：channel适合广播将数据一对多进行传递，锁更多倾向于一对一的传递
*/

import (
	"fmt"
	"time"
)

// 只管接收功能的渠道函数
func receiveOnly(ch <-chan int) {
	for v := range ch {
		fmt.Printf("接收的数据%d\n", v)
	}
}

// 只管发送功能的渠道函数
func sendOnly(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("发送的数据%d\n", i)
	}
	close(ch)
}

// 缓冲长度的使用：没有缓冲则必须一边消费一边生产，即同步，如果有缓冲长度则实现了异步，
// 即可以在缓冲长度的情况下，一边随便放，一边随便消费
func main() {
	//创建渠道
	ch := make(chan int, 3)
	go sendOnly(ch)
	go receiveOnly(ch)
	//使用selector进行多路复用
	timeout := time.After(2 * time.Second)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("channel已关闭")
				return
			}
			fmt.Printf("主goroutine接收到:%d\n", v)
		case <-timeout:
			fmt.Println("操作超时")
			return
		default:
			fmt.Println("没有数据等待中....")
			time.Sleep(500 * time.Millisecond)
		}

	}
}
