package main

import (
	"fmt"
	"unsafe"
)

func main() {
	////////////////////////////////////4.布尔//////////////////////////////////
	//取值：true、false
	//类型：bool
	t1 := true
	fmt.Printf("t1:%v, 类型: %T, 所占字节: %d \n", t1, t1, unsafe.Sizeof(t1))
	//t1:true, 类型: bool, 所占字节: 1
}
