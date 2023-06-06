package main

import (
	"fmt"
	"test/demo"
)

func main() {

	demo.Cal() //demo ： package调用

	// demo.IntToString(5) //int 转string

	// demo.FordemoWithRune() //demo：for range

	// demo.SwitchCaseDemo() //demo:switch case

	// demo.PrintPyramid(5, true) //打印n层金字塔（是否需要空心）

	// demo.Recursion(4) //demo：递归调用

	// fmt.Println(demo.SumALl(10, 20, 15))

	// //使用指针，调换两个数
	// n1, n2 := 10, 20
	// demo.Swap(&n1, &n2)
	// fmt.Printf("n1=%v,n2=%v \n", n1, n2)

	// //闭包的使用：给文件名加后缀，如果有就忽略，没有就加上
	// f := demo.Makesuffix(".jpg")
	// fmt.Println("dog 文件处理后文件名为：", f("dog"))                           //dog.jpg
	// fmt.Println("winter.jpg 文件处理后文件名为：", f("winter"))                 //winter.jpg
	// fmt.Println("phone 文件处理后文件名为：", demo.Makesuffix(".exe")("phone")) //phone.exe

	// demo.BubbleSort(&[]int{34, 21, 6, 30, 16, 9}) //demo：冒泡排序

	// //二分法查找
	// var arr = []int{6, 9, 12, 30, 46, 99}
	// demo.BinarySearch(&arr, 0, len(arr), 99)

	// //interface接口demo，实现切片排序
	// demo.SortSlice()

	/*
		CRM案例，包括增删改查
		用到了封装、切片、
	*/
	// view.CrmRun()

	/*
		json序列化
	*/
	// demo.JsonParse()

	monster := Monster{
		Name: "tom",
		Age:  18,
	}
	interfaceTest(monster)
}

type Monster struct {
	Name string
	Age  int
}

func interfaceTest(inter interface{}) {

	v, err := inter.(Monster)
	if err {
		fmt.Printf("type is %T,value.Name is %v", v, v.Name)
	}

}
