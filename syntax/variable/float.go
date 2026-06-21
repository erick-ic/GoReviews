package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// //////////////////////////////////2.浮点型,小数//////////////////////////////////
	//取值：1.2、0.23、-1.023
	//类型：float32（单精度）、float64（默认类型，双精度）

	var n6 float32 = 21.21
	fmt.Printf("n6: %f, 类型: %T, 所占字节: %d\n", n6, n6, unsafe.Sizeof(n6))
	//n6: 21.209999, 类型: float32, 所占字节: 4

	var n7 float64 = 22.22
	fmt.Printf("n7: %f, 类型: %T, 所占字节: %d\n", n7, n7, unsafe.Sizeof(n7))
	//n7: 22.220000, 类型: float64, 所占字节: 8

}
