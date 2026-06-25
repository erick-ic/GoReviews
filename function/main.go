package main

import (
	newUtils "GoReviews/utils"
	"fmt"
)

var (
	Func1 = func(a int, b int) int {
		return a + b
	}
)

func calculate(a, b float64, operate byte) float64 {
	var result float64
	switch operate {
	case '+':
		result = a + b
	case '-':
		result = a - b
	case '*':
		result = a * b
	case '/':
		result = a / b
	default:
		fmt.Println("操作错误")
	}
	return result
}

func testReturnValue() (name string, age int) {
	name = "Jack"
	age = 21
	return
}

func recursion(n int) {
	if n > 2 {
		n--
		recursion(n)
	}
	fmt.Println("n:", n)
	//后进先出
	//n = 2 => fmt 2
	//n = 3, n-- => fmt 2
	//n = 4, n-- => fmt 3
}

func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}

func getSum(n1 int, args ...int) int {
	sum := n1
	//args是slice切片，通过args[index]访问其中的元素
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func useDefer(n1 int, n2 int) int {
	//当函数执行到defer时，暂不执行，会将defer后方的语句压入独立的栈（defer栈）,也会将对应的值拷贝进栈。
	//当函数执行完毕后，再从defer栈，按照先入后出的方式出栈
	defer fmt.Println("n1:", n1)
	defer fmt.Println("n2:", n2)
	res := n1 + n2
	fmt.Println("res:", res)
	return res

	//res: 3
	//n2: 2
	//n1: 1
}

func main() {
	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>1.基本实现>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//基本语法：
	//func 函数名(形参列表) (返回值列表){
	//	执行语句
	//	return 返回值列表
	//}

	fmt.Printf("%.2f \n", calculate(1.1, 2.2, '+'))
	//3.30

	//返回值
	res1, res2 := testReturnValue()
	fmt.Printf("res1: %s, res2: %d \n", res1, res2)
	//res1: Jack, res2: 21

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>2.包的使用>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	newUtils.SayHello("Jack")
	//hello Jack

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>3.递归>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//栈运行机制：后进先出
	recursion(4)
	//n: 2
	//n: 2
	//n: 3

	//斐波那契数：1，1，2，3，5，8，13...
	fmt.Println("fbn():", fbn(3))
	fmt.Println("fbn():", fbn(5))
	//fbn(): 2
	//fbn(): 5

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>4.可变参数>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	fmt.Println("getSum():", getSum(10, 0, -1, 10, 1))
	//getSum(): 20

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>5.匿名函数>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	data := func(n1, n2 int) int {
		return n1 + n2
	}(1, 2)
	fmt.Println("data:", data)
	//data: 3

	a := func(n1, n2 int) int {
		return n1 + n2
	}
	fmt.Println("a:", a(1, 2))
	//a: 3

	//全局匿名函数
	fmt.Println("Func1:", Func1(1, 2))
	//Func1: 3

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>6.defer关键字>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	useDefer(1, 2)
}
