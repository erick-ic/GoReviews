package main

import "fmt"

func main() {
	//1.制表符
	fmt.Println("Hello\tWorld")
	//Hello		World

	//2.换行符
	fmt.Println("Hello\nWorld")
	//Hello
	//World

	//3.\\
	fmt.Println("Hello \\ World")
	//Hello \ World

	//4.\"
	fmt.Println("Hello\"World")
	//Hello"World

	//5.\r
	fmt.Println("Hello\r World")
	// 从\r前面开始输出
	//World
}
