package main

import "fmt"

//文件中包含全局变量时，执行流程：
//变量定义 => init() => main()

var num = test()

func init() {
	//使用场景：完成初始化配置
	fmt.Println("hello init()...")
}

func test() int {
	fmt.Println("hello test()...")
	return 1
}

func main() {
	fmt.Println("hello main()...")
	fmt.Println("num:", num)
}

//hello test()...
//hello init()...
//hello main()...
