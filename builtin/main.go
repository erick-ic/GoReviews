package main

import "fmt"

func main() {
	//(1)new方法：
	//使用场景：分配内存。主要用来分配值类型，如int、float32、struct等，返回指针。
	num1 := 100
	fmt.Printf("num1 = %v, type = %T, address = %v \n", num1, num1, &num1)
	//num1 = 100, type = int, address = 0x4f3c73306020

	num2 := new(int)
	fmt.Printf("num2 = %v, type = %T, address = %v, 指向的值 = %v \n", num2, num2, &num2, *num2)
	//new生成的是指针类型：*int
	//值为指针地址：0x237ca4f300d0
	//本身地址：0x237ca4f20040
	//指向的值：0
	//num2 = 0x237ca4f300d0, type = *int, address = 0x237ca4f20040, 指向的值 = 0

	*num2 = 21
	fmt.Printf("num2 = %v, type = %T, address = %v, 指向的值 = %v \n", num2, num2, &num2, *num2)
	//new生成的是指针类型：*int
	//值为指针地址：0x237ca4f300d0
	//本身地址：0x237ca4f20040
	//指向的值：21
	//num2 = 0x237ca4f300d0, type = *int, address = 0x237ca4f20040, 指向的值 = 21

	//(2)make方法：
	//使用场景：分配内存。只能用于分配引用类型，如channel、map、slice。
}
