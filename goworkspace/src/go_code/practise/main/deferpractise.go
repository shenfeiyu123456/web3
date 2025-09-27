package main

import (
	"fmt"
)

/*
*
defer 函数大家肯定都用过，它在声明时不会立刻去执行，而是在函数 return 后去执行的。
它的主要应用场景有异常处理、记录日志、清理数据、释放资源 等等。
这篇文章不是分享 defer 的应用场景，而是分享使用 defer 需要注意的点。
*/
func main() {
	fmt.Println("main start")
	defer calc("0", 1, 3)
	defer calc("1", 2, 4)
	fmt.Println("main end")
	fmt.Println("main end")
	//结论：defer 函数定义的顺序 与 实际执的行顺序是相反的，也就是最先声明的最后才执行。
}

func calc(index string, a, b int) int {
	sum := a + b
	fmt.Println(index, a, b, sum)
	return sum
}
