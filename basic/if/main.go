package main

import "fmt"

func main() {
	//1.基本形式：
	//score := 60
	//if score < 60 {
	//	fmt.Println("不合格")
	//} else {
	//	fmt.Println("合格")
	//}

	//if score := 60; score >= 60 {
	//	fmt.Println("合格")
	//} else {
	//	fmt.Println("不合格")
	//}

	//2.多分支
	if score := 60; score < 60 {
		fmt.Println("不合格")
	} else if score >= 60 && score < 80 {
		fmt.Println("良好")
	} else {
		fmt.Println("优秀")
	}
	//3.嵌套分支
	//嵌套分支推荐控制在3层以内
	//if 条件表达式 {
	//	if 条件表达式	{
	//		代码块...
	//	}else{
	//		代码块...
	//	}
	//}else{
	//	代码块...
	//}
}
