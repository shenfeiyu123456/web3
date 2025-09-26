package main

/**
切片是一种动态数组，比操作数组更加灵活，长度可以动态变动，可以追加和删除,len()和cap()返回的结果可以不相同
*/
import "fmt"

func main() {
	//和数组的声明方式区别是不指定容器中的元素个数和具体值
	var slice []int
	fmt.Println("slice", slice)
	var slice1 = []int{}
	fmt.Println("slice1", slice1)
	slice2 := []int{1, 2, 3, 4, 5} //[]中没有指定元素个数即为切片，指定元素个数了就是数组了
	fmt.Println("slice2", slice2)
	//指定内容长度创建切片
	//make的作用是创建一个切片
	//参数说明：第一个参数[]int 表示创建一个int类型的切片
	//第二个参数5表示切片的初始长度为5，即包含5个默认值0
	//第三个参数8表示切片的容量，即底层数组可以容纳的最大元素数量
	//与new不同，make会初始化内存并准备好直接使用
	//对于切片，make会创建底层数组并返回切片引用
	//容量可以省略，此时容量=长度
	//简单说，这行代码创建了一个长度为5（可访问5个元素）、容量为8（可扩展至8个元素）的int切片。
	var slice3 []int = make([]int, 5, 8)
	fmt.Println("slice3", slice3)
	fmt.Println(len(slice3), cap(slice3))
	slice4 := make([]int, 3, 5)
	fmt.Println("slice4", slice4)
	fmt.Println(len(slice4), cap(slice4))
	//截取切片(包含开始不包含结尾)
	slice5 := []int{1, 2, 3, 4, 5}
	fmt.Println("slice5[1] == ", slice5[1])
	//[:]指的是从开头到结尾
	fmt.Println("slice5[:] == ", slice5[:])
	//[1:]指的是第二个元素到末尾元素
	fmt.Println("slice5[1:] == ", slice5[1:])
	//[:4]指的是第一个元素到第四个
	fmt.Println("slice5[:4] == ", slice5[:4])
	//切片的追加
	slice6 := []int{1, 2, 3, 4, 5}
	fmt.Println("slice6 == ", slice6)
	slice6 = append(slice6, 6)
	fmt.Println("slice6追加后的数据 == ", slice6)
	//切片的删除
	slice7 := []int{1, 2, 3, 4, 5}
	fmt.Println("slice7 == ", slice7)
	fmt.Println("slice7[:1] == ", slice7[:1])
	fmt.Println("slice7[1:] == ", slice7[1:])
	//删除头部元素 [1:]表示 第二个元素到最后一个
	slice7 = append(slice7[1:])
	fmt.Println("slice7删除头部元素为 == ", slice7)
	//删除尾部元素 [:3]表示0到第四个元素，不包含第四个元素
	slice7 = append(slice7[:3])
	fmt.Println("slice7删除尾部元素为 == ", slice7)
}
