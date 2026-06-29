package main

import "fmt"

func checkType(x interface{}) {
	// if _, ok := x.(string); ok {
	// 	fmt.Println("string类型")
	// } else if _, ok := x.(int); ok {
	// 	fmt.Println("int类型")
	// } else if _, ok := x.(bool); ok {
	// 	fmt.Println("bool类型")
	// } else {
	// 	fmt.Println("断言失败")
	// }
	switch x.(type) {
	// switch语句中x.(type)判断一个变量的类型
	case int:
		fmt.Println("int类型")
		break
	case string:
		fmt.Println("string类型")
		break
	case bool:
		fmt.Println("bool类型")
		break
	default:
		fmt.Println("断言失败")
	}
}

type Leader struct {
	Name string
	Age  int
}

func main() {
	//使用场景：由于接口可以接受任意类型，若转成具体类型，则需使用类型断言。

	//1.基本语法：
	var a interface{}
	a = 21
	v, ok := a.(int)
	if ok {
		fmt.Println(v, ok)
		//21 true
	} else {
		fmt.Println("断言失败")
	}

	var x interface{}
	//x = 1
	//fmt.Println("x + 1 = ", x+1)
	//编译时x为interface{}类型，无法进行运算
	////invalid operation: x + 1 (mismatched types interface{} and untyped int)

	var num int = 1
	x = num //空接口接受任意类型
	//var x interface{} 这个声明，决定了编译器在编译时看到的永远是 interface{}。
	fmt.Printf("x: %v, type: %T\n", x, x)
	//运行时结果：x: 1, type: int
	/*
		Printf 中的 %T 是在程序运行时去读取接口内部存储的动态类型信息，
		所以它告诉你“里面是 int”。
		但在编写 value := x.(int) 这行代码时，编译器还不知道里面一定是 int。
	*/

	value := x.(int)
	fmt.Printf("value: %v, type: %T\n", value, value)
	//value: 1, type: int
	//类型断言 value := x.(int) 的作用，就是把 x 从通用的 interface{} 转换成具体的 int 类型

	//2.多分支断言：
	checkType(21)
	//int类型

	//3.结构体断言：
	user := make(map[string]interface{})
	user["name"] = "Jack"
	user["age"] = 18
	user["hobbies"] = []string{"sing", "code"}
	fmt.Println("user:", user)
	//user: map[age:18 hobbies:[sing code] name:Jack]
	fmt.Println("hobbies:", user["hobbies"])
	//hobbies: [sing code]

	leader := Leader{
		Name: "Tome",
		Age:  18,
	}
	fmt.Println("leader.Name:", leader.Name)
	//leader.Name: Tome

	user["leader"] = leader
	fmt.Println("user:", user)
	//user: map[age:18 hobbies:[sing code] leader:{Tome 18} name:Jack]

	//leaderObj, ok := user["leader"].(Leader)
	//if ok {
	//	fmt.Println("leaderObj", leaderObj)
	//	//leaderObj {Tome 18}
	//} else {
	//	fmt.Println("0")
	//}
	leaderObj := user["leader"].(Leader)
	fmt.Println("leaderObj:", leaderObj)
	//leaderObj: {Tome 18}

}
