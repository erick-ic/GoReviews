package main

import "fmt"

const (
	A = 0
	B = 1
	C = 2
)

const (
	X = iota
	Y
	Z
)

type Color int

const (
	Red = iota
	Green
	Blue
)

const (
	a = iota
	b
	c
	d = 33
	e
	f
	g = iota
	h
	i
)

func main() {
	//(1)基本用法：
	fmt.Printf("A=%d,  B=%d,  C=%d \n", A, B, C)
	//A=0,  B=1,  C=2
	fmt.Printf("typeA=%T,  typeB=%T, typeC=%T \n", A, B, C)
	//typeA=int,  typeB=int, typeC=int

	//(2)枚举：
	//iota是一个内置的常量生成器，用于创建递增常量
	//无类型常量
	fmt.Printf("X=%d,  Y=%d,  Z=%d \n", X, Y, Z)
	//X=0,  Y=1,  Z=2
	fmt.Printf("typeX=%T,  typeY=%T,  typeZ=%T \n", X, Y, Z)
	//typeX=0,  typeY=1,  typeZ=2

	//有类型的常量
	fmt.Printf("Red=%d,  Green=%d,  Blue=%d \n", Red, Green, Blue)
	//Red=0,  Green=1,  Blue=2
	fmt.Printf("typeRed=%T,  typeGreen=%T,  typeBlue=%T \n", Red, Green, Blue)
	//typeRed=0,  typeGreen=1,  typeBlue=2

	//常量顺序
	//const (
	//	a = iota
	//	b
	//	c
	//	d = 33
	//	e
	//	f
	//	g = iota
	//	h
	//	i
	//)
}
