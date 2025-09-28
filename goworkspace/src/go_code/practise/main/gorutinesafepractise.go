package main

/*
go中并发同样存在线程安全问题，因为GO也是共享内存使用多个gorutine之间通信，并且大部分时间为了性能，
所以go的大多数标准库的数据结构默认是非线程安全的
尽量避免程序使用共享变量的操作，或者使用局部变量进行替换
*/
import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter 线程安全计数器
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

// Inc 增加计数
func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// GetCount 获取当前计数
func (c *SafeCounter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

// UnsafeCounter 定义unsafecounter结构体
type UnsafeCounter struct {
	count int
}

func (c *UnsafeCounter) Inc() {
	c.count += 1
}

func (c *UnsafeCounter) GetCount() int {
	return c.count
}

func main() {
	//counter := UnsafeCounter{}
	counter := SafeCounter{}
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				counter.Inc()
			}
		}()
	}
	//等待goroutine时间完成
	time.Sleep(3 * time.Second)
	//输出最终数据完成
	fmt.Println("最终获取的计数为:", counter.GetCount())
}
