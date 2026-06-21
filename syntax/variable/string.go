package main

import (
	"fmt"
	"unsafe"
)

func main() {
	////////////////////////////////////3.字符串//////////////////////////////////

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>3.1字符类型>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//字符常量通常是用单引号包裹的单个字符，英文1个字节，汉字3个字节。

	//字符型在计算机中的运行过程：
	//存储：字符 => 对应码值 => 二进制 => 存储完成
	//读取：二进制 => 码值 => 字符 => 读取完成

	//使用byte保存单个字母字符
	//字符的本质是一个整数，直接输出时对应UTF-8的码值（UTF-8是Unicode编码的具体实现）
	//%c:格式化输出字符
	var b1 byte = 'a'
	fmt.Printf("b1:%d, 格式化输出：%c, 类型: %T, 所占字节: %d\n", b1, b1, b1, unsafe.Sizeof(b1))
	//b1:97, 格式化输出：a, 类型: uint8, 所占字节: 1

	var b2 = '中'
	fmt.Printf("b2:%d, 格式化输出：%c, 类型: %T, 所占字节: %d\n", b2, b2, b2, unsafe.Sizeof(b2))
	//rune类型
	//b2:20013, 格式化输出：中, 类型: int32, 所占字节: 4

	//var b3 = '123'
	//fmt.Printf("b3:%d ,格式化输出：%c, 类型: %T, 所占字节: %d\n", b3, b3, b3, unsafe.Sizeof(b3))
	//more than one character in rune literal

	//数字变量%c输出时，会输出该数字对应的Unicode字符
	var b4 int = 22269
	fmt.Printf("b4:%d ,格式化输出：%c, 类型: %T, 所占字节: %d\n", b4, b4, b4, unsafe.Sizeof(b4))
	//b4:22269 ,格式化输出：国, 类型: int, 所占字节: 8

	//字符类型可以进行运算，相当于一个整数，因为都有对应的Unicode码
	var b5 = 10 + 'a'
	fmt.Printf("b5:%d ,格式化输出：%c, 类型: %T, 所占字节: %d \n", b5, b5, b5, unsafe.Sizeof(b5))
	//b5:107 ,格式化输出：k, 类型: int32, 所占字节: 4

	//3.2字符串
	//go中的字符串是由单个字节连接组成的

	var s1 = "中"
	fmt.Printf("s1:%v, 类型: %T, 所占字节: %d\n", s1, s1, unsafe.Sizeof(s1))
	//s1:中, 类型: string, 所占字节: 16

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>3.2字符串类型>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//字符串的字节使用UTF-8编码标识Unicode文本，解决中文乱码问题
	str1 := "hello"
	str1 = "world"
	fmt.Printf("str1:%v \n", str1)
	//str1:world

	//字符串不可变，一旦赋值，就无法更改
	//str1[0] = "a"
	//fmt.Printf("str1:%v \n", str1)
	//cannot assign to str1[0] (neither addressable nor a map index expression)
}
