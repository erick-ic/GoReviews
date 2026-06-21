package main

import (
	"fmt"
	"unsafe"
)

func main() {
	////////////////////////////////////1.int 整型//////////////////////////////////
	//取值：0、-1、123等数值
	//类型：int（默认类型）、uint、rune、byte

	n1 := 21
	fmt.Printf("n1: %d, 类型: %T, 所占字节: %d\n", n1, n1, unsafe.Sizeof(n1))
	//n1: 21, 类型: int, 所占字节: 8

	//有符号整型
	//int8:		1字节，-2^7 ~ 2^7-1(-128 ~ 127)
	//int16:	2字节，-2^15 ~ 2^15-1
	//int32:	4字节，-2^31 ~ 2^31-1
	//int64:	8字节，-2^63 ~ 2^63-1

	var n2 int8 = -128
	n2 = 127
	fmt.Printf("n2: %d, 类型: %T, 所占字节: %d\n", n2, n2, unsafe.Sizeof(n2))
	//n2: 127, 类型: int8, 所占字节: 1
	//n2 = -129
	fmt.Printf("n2: %d, 类型: %T, 所占字节: %d\n", n2, n2, unsafe.Sizeof(n2))
	//cannot use -129 (untyped int constant) as int8 value in assignment (overflows)

	//无符号整型
	//uint8:	1字节，0 ~ 2^8-1(0 ~ 255)
	//uint16:	2字节，0 ~ 2^16-1
	//uint32:	4字节，0 ~ 2^32-1
	//uint64:	8字节，0 ~ 2^64-1

	//var n3 uint = 255
	//fmt.Printf("n3: %d, 类型: %T, 所占字节: %d\n", n3, n3, unsafe.Sizeof(n3))
	////n3: 255, 类型: uint, 所占字节: 8
	var n3 uint8 = 255
	fmt.Printf("n3: %d, 类型: %T, 所占字节: %d\n", n3, n3, unsafe.Sizeof(n3))
	//n3: 255, 类型: uint8, 所占字节: 1
	//n3 = 256
	//fmt.Printf("n3: %d, 类型: %T, 所占字节: %d\n", n3, n3, unsafe.Sizeof(n3))
	//cannot use 256 (untyped int constant) as uint8 value in assignment (overflows)

	//rune等价于int32，表示一个Unicode码
	//int32:	4字节，-2^31 ~ 2^31-1
	var n4 rune = 10
	fmt.Printf("n4: %d, 类型: %T, 所占字节: %d\n", n4, n4, unsafe.Sizeof(n4))
	//n4: 10, 类型: int32, 所占字节: 4

	//byte等价于uint8，存储字符时使用
	//byte是uint8的别名
	//uint8:	1字节，0 ~ 2^8-1(0 ~ 255)
	//1字节=8bit
	var n5 byte = 10
	fmt.Printf("n5: %d, 类型: %T, 所占字节: %d\n", n5, n5, unsafe.Sizeof(n5))
	//n5: 10, 类型: uint8, 所占字节: 1
}
