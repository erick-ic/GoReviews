package main

import "fmt"

func main() {
	//获取指针地址：&
	//获取指针类型的值：*
	var num int = 100
	fmt.Printf("num: %v, 指针地址: %v \n", num, &num)
	//num: 100, 指针地址: 0x542f419e020

	//num2：指针变量，类型为*int，取值为nil
	var num2 *int
	fmt.Printf("num2: %v, 指针地址: %v \n", num2, &num2)
	//num2: <nil>, 指针地址: 0x542f4182040

	//num2：指针变量，类型为*int，取值为&num，即num变量的指针地址
	num2 = &num
	fmt.Printf("num2: %v, 指针地址: %v, 本身的值：%v \n", num2, &num2, *num2)
	//num2: 0x542f419e020, 指针地址: 0x542f4182040, 本身的值：100

	//案例：获取一个int类型的变量x的地址，显示到终端，将x的地址赋值给指针ptr，并通过ptr修改x的值
	var x int = 21
	var ptr *int = &x
	fmt.Printf("x: %v, ptr: %v \n", *ptr, *ptr)
	//x: 21, ptr: 21
	*ptr = 100
	fmt.Printf("x: %v, ptr: %v \n", *ptr, *ptr)
	//x: 100, ptr: 100

}
