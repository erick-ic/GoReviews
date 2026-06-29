package main

import (
	"errors"
	"fmt"
)

func handleErr() {
	//使用defer+recover捕获和处理异常
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err:", err)
			//将错误发送给管理员...
		}
	}()

	num1 := 1
	num2 := 0
	res := num1 / num2
	fmt.Println("res:", res)
	//err: runtime error: integer divide by zero
}

func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	}

	//返回自定义错误
	return errors.New("读取文件错误")
}

func handleCustomErr() {
	err := readConf("config.ini")
	if err != nil {
		//文件读取错误，输出错误并终止程序
		panic(err)
	}
	fmt.Println("after err...")
}

func main() {
	//在默认情况下，当发生错误后（panic）, 程序就会退出（崩溃）

	//示例：
	//num1 := 1
	//num2 := 0
	//res := num1 / num2
	//fmt.Println("res:", res)
	////panic: runtime error: integer divide by zero

	//(1)Go的错误处理：
	//处理流程：抛出一个panic异常，在defer中通过recover捕获异常，进行处理。
	//handleErr()
	//fmt.Println("其他执行逻辑...")
	//其他执行逻辑...

	//(2)自定义错误：
	handleCustomErr()
	fmt.Println("其他执行逻辑...")
	//其他执行逻辑...
}
