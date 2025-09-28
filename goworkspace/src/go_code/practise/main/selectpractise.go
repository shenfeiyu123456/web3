package main

import (
	"fmt"
	"time"
)

/**
select 语义是和channel绑定在一起使用的，select可以实现多个channel收发数据，可以使得一个goroutine就可以处理多个channel的通信
语法与switch相似，有case和default分支，只不过select的每个case后面跟的是chinnel的收发工作
*/

func main() {
	//创建三个channel
	ch := make(chan string, 10)
	ch1 := make(chan string, 10)
	ch2 := make(chan string, 10)
	ch3 := make(chan string, 10)
	//启动一个协程 循环将数据写入到三个渠道中
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- "hello"
			ch2 <- "world"
			ch <- "info"
		}
	}()
	//使用select将写入渠道的消息进行获取
	timeOut := time.After(2 * time.Second)
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		case msg3 := <-ch:
			fmt.Println(msg3)
		case ch3 <- "world1":
			fmt.Println("使用select将数据写道ch3中", ch3)
		case msg4 := <-ch3:
			fmt.Println("使用select读取ch3中的数据", msg4)
		case <-timeOut:
			fmt.Println("timeout")
			return
		}
	}
}
