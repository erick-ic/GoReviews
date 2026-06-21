package main

import "fmt"

func main() {
	//////////////////////////////////变量声明//////////////////////////////////
	//1.声明后不赋值，使用默认值
	var num1 int
	fmt.Printf("num:%d, type:%T \n", num1, num1)
	//num:0, type:int

	//2.类型推导
	var num2 = 21
	fmt.Printf("num2:%d, type:%T \n", num2, num2)
	//num2:21, type:int

	//3.省略var关键字，:=左侧变量不应是已声明的，否则编译错误
	num3 := 3
	fmt.Println("num:", num3)
	//num: 3

	//4.多变量
	n1, s1, f1 := 21, "hello", 21.21
	fmt.Printf("n1:%d, n2:%s, n3:%f, \n", n1, s1, f1)
	//n1:21, n2:hello, n3:21.210000,
	fmt.Printf("n1:%T, n2:%T, n3:%T, \n", n1, s1, f1)
	//n1:int, n2:string, n3:float64,

	var (
		t1 = 100
		t2 = 200
		t3 = 300
	)
	fmt.Println(t1, t2, t3)
	//100 200 300

	//5.变量可以重复赋值，但需要符合类型
	var n4 int = 100
	n4 = 200
	fmt.Println("n4:", n4)
	//n4: 200

	//n4 = 1.1
	//fmt.Println("n4:", n4)
	// cannot use 1.1 (untyped float constant) as int value in assignment (truncated)
}
