package main

import "fmt"

func main() {
	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>1.for循环>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//方式1:
	for i := 0; i <= 5; i++ {
		fmt.Println("i:", i)
	}
	//i: 0
	//i: 1
	//i: 2
	//i: 3
	//i: 4
	//i: 5

	//方式2:
	j := 1
	for j <= 5 {
		fmt.Println("j:", j)
		j++
	}
	//i: 0
	//i: 1
	//i: 2
	//i: 3
	//i: 4
	//i: 5

	//方式3:
	k := 1
	for { // 等价于 for ; ; {
		if k <= 5 {
			fmt.Println("k:", k)
		} else {
			break
		}
		k++
	}
	//k: 1
	//k: 2
	//k: 3
	//k: 4
	//k: 5

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>2.for range>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//遍历字符串-传统方式：
	str := "hello"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c \n", str[i])
	}
	//h
	//e
	//l
	//l
	//o
	//遍历字符串-for range：
	for inx, val := range str {
		fmt.Printf("index: %d, value: %c\n", inx, val)
	}
	//index: 0, value: h
	//index: 1, value: e
	//index: 2, value: l
	//index: 3, value: l
	//index: 4, value: o

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>3.for实现while、do while>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//go语言中没有while、do while语法

	//3.1.for循环实现while语法：
	//for {
	//	if 循环表达式 {
	//		break
	//	}
	//	循环操作
	//	循环变量迭代
	//}
	var m int = 1
	for {
		if m >= 5 {
			break
		}
		fmt.Println("m:", m)
		m++
	}
	//m: 1
	//m: 2
	//m: 3
	//m: 4

	//3.2.for循环实现do while语法：至少执行一次循环操作
	//for {
	//	循环操作
	//	循环变量迭代
	//	if 循环表达式 {
	//		break
	//	}
	//}
	var n int = 1
	for {
		fmt.Println("n:", n)
		n++
		if n <= 5 {
			break
		}
	}
	//n: 1

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>4.break跳转控制语句>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//当break语句出现在多层嵌套循环语句中时，可以通过标签指明终止的语句块
	for i := 0; i < 3; i++ {
	label1: //设置标签
		for j := 0; j <= 2; j++ {
			if j == 2 {
				//break //跳出当前循环
				break label1 // 等同于break，跳出当前循环
			}
			fmt.Println("j=", j)
		}
	}
	//j= 0
	//j= 1
	//j= 0
	//j= 1
	//j= 0
	//j= 1

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>5.continue跳转控制语句>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//结束本次循环，继续执行下一次循环
	for i := 0; i < 2; i++ {
		for j := 0; j <= 4; j++ {
			if j == 2 {
				continue
			}
			fmt.Println("j:", j)
		}
	}
	//j: 0
	//j: 1
	//j: 3
	//j: 4
	//j: 0
	//j: 1
	//j: 3
	//j: 4

here:
	for x := 0; x < 2; x++ {
		for y := 1; y < 4; y++ {
			if y == 2 {
				continue here
			}
			fmt.Printf("x: %v, y: %v \n", x, y)
		}
	}
	//x: 0, y: 1
	//x: 1, y: 1

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>6.goto跳转控制语句>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//不推荐使用goto语句
	fmt.Println("text1")
	fmt.Println("text2")
	if num := 30; num < 100 {
		goto lable1
	}
	fmt.Println("text3")
	fmt.Println("text4")
lable1:
	fmt.Println("text5")
	fmt.Println("text6")

	//text1
	//text2
	//text5
	//text6

}
