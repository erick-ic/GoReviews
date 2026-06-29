package main

import (
	"fmt"
	"strconv"
)

func main() {
	//////////////////////////////////一、变量声明//////////////////////////////////
	//1.声明后不赋值，使用默认值
	var num01 int
	fmt.Printf("num01:%d, type:%T \n", num01, num01)
	//num:01, type:int

	//2.类型推导
	var num02 = 21
	fmt.Printf("num2:%d, type:%T \n", num02, num02)
	//num02:21, type:int

	//3.省略var关键字，:=左侧变量不应是已声明的，否则编译错误
	num03 := 3
	fmt.Println("num03:", num03)
	//num03: 3

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

	// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>二、基本数据类型>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//类型：整型、浮点型、字符串、布尔型
	//%v：按照变量的值输出
	var age int
	var score float32
	var salary float64
	var isAdulated bool
	var name string
	fmt.Printf("age:%d \n", age)
	fmt.Printf("score:%f \n", score)
	fmt.Printf("salary:%v \n", salary)
	fmt.Printf("isAdulated:%v \n", isAdulated)
	fmt.Printf("name:%v \n", name)
	//age:0
	//score:0.000000
	//salary:0
	//isAdulated:false
	//name:

	// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>三、基本类型转换>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

	//1.基本数据类型转string
	var num int = 100
	var num2 float64 = 23.456
	var flag bool = true
	var myChar byte = 'a'
	var nickname string = "Jack"

	//1.1fmt.Sprintf()方法
	//%v 值的默认格式
	//%T 值的类型
	//%d 十进制
	//%f 浮点数
	//%t 布尔值
	//%c 值对应的Unicode码值
	//%s 字符串或字节

	//int => string
	str := fmt.Sprintf("%d", num)
	fmt.Printf("str:%v, type:%T \n", str, str)
	//str:100, type:string

	//被转换的是变量存储的值，原变量数据类型未发生改变
	fmt.Printf("num:%v, type:%T \n", num, num)
	//num:100, type:int

	//float => string
	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str:%v, type:%T \n", str, str)
	//str:23.456000, type:string

	//bool => string
	str = fmt.Sprintf("%t", flag)
	fmt.Printf("str:%v, type:%T \n", str, str)
	//str:true, type:string

	//byte => string
	//%c（字符动词）：专门用于将 整数（byte/rune） 格式化为对应的 ASCII 或 Unicode 字符。
	//把数字 97 传进去，它吐出字符 'a'。
	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str:%v, type:%T \n", str, str)
	//str:a, type:string

	str = string(myChar)
	fmt.Printf("str:%v, type:%T \n", str, str)
	//str:a, type:string

	//string => string
	str = fmt.Sprintf("%s", nickname)
	fmt.Printf("str:%v, type:%T \n", str, str)
	//str:Jack, type:string

	//1.2.strconv函数
	// int => string
	// base: 进制，取值范围2 ～ 36
	str = strconv.FormatInt(int64(num), 10)
	fmt.Printf("str:%v, type:%T \n", str, str)
	//str:100, type:string

	str = strconv.Itoa(num)
	fmt.Printf("str:%v, type:%T \n", str, str)
	//str:100, type:string

	//float => string
	//fmt: 输出格式，f(-ddd.ddd)
	//prec: 精度，小数点位数，排除指数部分。-1，代表使用最少数量，但又必需的数字来表示。
	//bitsize: 数据来源类型
	str = strconv.FormatFloat(num2, 'g', -1, 64)
	fmt.Printf("str:%v, type:%T \n", str, str)
	//str:23.456, type:string

	//bool => string
	str = strconv.FormatBool(flag)
	fmt.Printf("str:%v, type:%T \n", str, str)
	//str:true, type:string

	//2.string转基本数据类型
	//注意点：需要转为有效值，否则输出目标类型，但取零值。

	var str2 string = "21"
	var str3 string = "21.21"
	var str4 string = "true"

	//string => int
	//base: 进制，取值范围2 ～ 36，若为0，则根据字符串前置判断，0X为16进制，0为8进制，剩下10进制
	//bitsize: 类型，0，8，16，32，64
	int2, _ := strconv.ParseInt(str2, 10, 64)
	fmt.Printf("str:%v, type:%T \n", int2, int2)
	//str:21, type:int64

	//string => float
	float3, _ := strconv.ParseFloat(str3, 64)
	fmt.Printf("str:%v, type:%T \n", float3, float3)
	//str:21.21, type:float64

	//string => bool
	bool4, _ := strconv.ParseBool(str4)
	fmt.Printf("str:%v, type:%T \n", bool4, bool4)
	//str:true, type:bool
}
