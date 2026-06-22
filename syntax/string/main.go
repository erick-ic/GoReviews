package main

import (
	"fmt"
	"strconv"
	"strings"
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

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>3.3字符串常用函数>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

	//(1)统计字符串长度:
	//golang的编码统一为utf-8，ASCII的字母和数字占1个字符，汉字占3个字符
	str2 := "hello北"
	fmt.Println("str2 len:", len(str2))
	//str2 len: 8

	//(2)遍历字符串:
	str3 := "hi成"
	//for i := 0; i < len(str3); i++ {
	//	fmt.Println("字符: ", str3[i])
	//}
	////字符= 104
	////字符= 105
	////字符= 230
	////字符= 136
	////字符= 144

	//for i := 0; i < len(str3); i++ {
	//	fmt.Printf("字符: %c \n", str3[i])
	//}
	////字符: h
	////字符: i
	////字符: æ
	////字符: 
	////字符:

	//遍历字符串，同时处理带有中文的问题：
	//r = []rune(str)
	r := []rune(str3)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符: %c \n", r[i])
	}
	//字符: h
	//字符: i
	//字符: 成

	//(3.1)字符串转整型：
	//res, error := strconv.Atoi(str)
	res, err := strconv.Atoi("123")
	//res, error := strconv.Atoi("hello")
	//strconv.Atoi: parsing "hello": invalid syntax
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("res:", res)
	}

	//(3.2)整型转字符串：
	fmt.Println("newStr:", strconv.Itoa(123))
	//newStr: 123

	//(4)10进制转2，8，16进制：
	//str = strconv.FormatInt(123, 2), 返回对应字符串
	str4 := strconv.FormatInt(123, 2)
	fmt.Printf("123对应的2进制 = %v, type = %T\n", str4, str4)
	//123对应的2进制 = 1111011, type = string

	//(5.1)字符串转[]byte：
	bytes := []byte("hi123成")
	fmt.Printf("str = %v, type = %T\n", bytes, bytes)
	//str = [104 105 49 50 51 230 136 144], type = []uint8
	fmt.Printf("str = %c, type = %T\n", bytes, bytes)
	//str = [h i 1 2 3 æ  ], type = []uint8

	//(5.2)[]byte转字符串：
	str5 := string([]byte{97, 98, 99})
	fmt.Printf("str = %v, type = %T\n", str5, str5)
	//str = abc, type = string

	//(6)判断字符串是否包含子串：
	b := strings.Contains("hello", "he")
	fmt.Printf("b = %v, type = %T\n", b, b)
	//b = true, type = bool

	//(7)统计字符串中字串数量：
	num := strings.Count("hello", "l")
	fmt.Printf("num = %v, type = %T\n", num, num)
	//num = 2, type = int

	//(8)不区分大小写的字符串比较：
	//== 区分大小写
	b = strings.EqualFold("hello", "Hello")
	fmt.Printf("b = %v, type = %T\n", b, b)
	//b = true, type = bool
	fmt.Println("区分大小写比较: ", "abc" == "Abc")
	//区分大小写比较:  false

	//(9)判断字符串中字串首次和最后出现的index：
	//未找到返回-1
	firstIndex := strings.Index("hello", "ll")
	lastIndex := strings.LastIndex("hello", "o")
	fmt.Printf("firstIndex = %v, type = %T\n", firstIndex, firstIndex)
	//firstIndex = 2, type = int
	fmt.Printf("lastIndex = %v, type = %T\n", lastIndex, lastIndex)
	//lastIndex = 4, type = int

	//(10)替换字符串：
	str6 := "go go hello"
	//n：表示替换的数量，-1表示全部替换
	str66 := strings.Replace(str6, "go", "golang", -1)
	fmt.Printf("str6 = %v, type = %T\n", str6, str6)
	//str6 = go go hello, type = string
	fmt.Printf("str66 = %v, type = %T\n", str66, str66)
	//str66 = golang golang hello, type = string

	//(11)字符串大小写转换：
	str7 := "golang"
	str77 := strings.ToUpper(str7)
	fmt.Printf("str7 = %v, type = %T\n", str7, str7)
	//str7 = GOLANG, type = string
	fmt.Printf("str77 = %v, type = %T\n", str77, str77)
	//str77 = GOLANG, type = string
	str777 := strings.ToLower(str77)
	fmt.Printf("str777 = %v, type = %T\n", str777, str777)
	//str777 = golang, type = string

	//(12)字符串分割：
	strList := strings.Split("hello-world", "-")
	for inx, str := range strList {
		fmt.Printf("str-%v: %v \n", inx, str)
	}
	//str-0: hello
	//str-1: world

	//(13)判断字符串开头和结束：
	str8 := "hello-world"
	fmt.Println("判断字符串开头：", strings.HasPrefix(str8, "he"))
	//判断字符串开头： true
	fmt.Println("判断字符串结束：", strings.HasSuffix(str8, "l"))
	//判断字符串结束： false
}
