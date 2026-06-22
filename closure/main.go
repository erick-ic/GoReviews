package main

import (
	"fmt"
	"strings"
)

// 累加器
func addUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
	//闭包内容：变量n与以下代码块构成闭包
	/*
		var n int = 10
		return func(x int) int {
			n = n + x
			return n
		}
	*/
}

func makeSuffix(suffix string) func(string) string {
	//若没有指定后缀，则加上，否则返回原名
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
	//闭包内容：变量suffix与以下代码块构成闭包
	/*
		return func(name string) string {
			if !strings.HasSuffix(name, suffix) {
				return name + suffix
			}
			return name
		}
	*/
}

func main() {
	fmt.Println("hello main()...")
	//闭包定义：一个函数和与其相关的引用环境组成的整体
	//闭包内容：返回一个匿名函数，且引用了匿名函数外的变量(n)。
	//变量只会初始化一次，后续调用只会在原基础上叠加

	f := addUpper()
	fmt.Println(f(1)) // 11
	fmt.Println(f(2)) // 13
	fmt.Println(f(3)) // 16

	//该实践中闭包的优势：后缀只需传入一次，闭包可以保留上次引用的某个值，便于重复使用。
	f2 := makeSuffix(".jpg")
	fmt.Println("f2():", f2("winter"))
	// f2(): winter.jpg
	fmt.Println("f2():", f2("winter.jpg"))
	// f2(): winter.jpg
}
