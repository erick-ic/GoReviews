package example

import (
	"fmt"
	"testing"
)

// 1.两个变量，取值互换
func TestExchange(t *testing.T) {
	a, b := exchange(10, 20)
	fmt.Printf("a: %d, b: %d \n", a, b)
	// a: 20, b: 10
}

// 2.打印1-100之间所有9的倍数的整数的个数，总和
func TestCalculateNine(t *testing.T) {
	count, sum := calculateNine(9, 100)
	fmt.Printf("count: %d, sum: %d\n", count, sum)
	//count: 11, sum: 594
}

// 3.按要求打印下列数字
func TestHandlePrint(t *testing.T) {
	handlePrint(6)
	/*
		0 + 6 = 6
		1 + 5 = 6
		2 + 4 = 6
		3 + 3 = 6
		4 + 2 = 6
		5 + 1 = 6
		6 + 0 = 6
	*/
}

//// 4.统计3个班成绩情况，每个班有5名学生，求出各个班级的平均分和所有班级的总分
//func TestHandleClass(t *testing.T) {
//	classAvg, allClassScore := handleClass(3, 5)
//	fmt.Printf("classAvg: %v, allClassAvg: %v \n", classAvg, allClassScore)
//}

// 5.打印金字塔：编写一个程序，接收一个整数表示金字塔的层数。
func TestPrintGoldTower(t *testing.T) {
	printGoldTower(3)
}

// 6.打印九九乘法表
func TestMultiTable(t *testing.T) {
	multiTable(9)
	//1 * 1 = 1
	//1 * 2 = 2 	2 * 2 = 4
	//1 * 3 = 3 	2 * 3 = 6 	3 * 3 = 9
	//1 * 4 = 4 	2 * 4 = 8 	3 * 4 = 12 	4 * 4 = 16
	//1 * 5 = 5 	2 * 5 = 10 	3 * 5 = 15 	4 * 5 = 20 	5 * 5 = 25
	//1 * 6 = 6 	2 * 6 = 12 	3 * 6 = 18 	4 * 6 = 24 	5 * 6 = 30 	6 * 6 = 36
	//1 * 7 = 7 	2 * 7 = 14 	3 * 7 = 21 	4 * 7 = 28 	5 * 7 = 35 	6 * 7 = 42 	7 * 7 = 49
	//1 * 8 = 8 	2 * 8 = 16 	3 * 8 = 24 	4 * 8 = 32 	5 * 8 = 40 	6 * 8 = 48 	7 * 8 = 56 	8 * 8 = 64
	//1 * 9 = 9 	2 * 9 = 18 	3 * 9 = 27 	4 * 9 = 36 	5 * 9 = 45 	6 * 9 = 54 	7 * 9 = 63 	8 * 9 = 72 	9 * 9 = 81
}

// 7.100以内数求和，求出当和第一次大于20的当前数
func TestGetSumWithinNum(t *testing.T) {
	data := getSumWithinNum(100)
	fmt.Println(data)
	//i: 6
	//sum: 21
	//21
}

// 8.10以内的奇数-for循环+continue
func TestFindOddNumber(t *testing.T) {
	findOddNumber(10)
	//i: 1
	//i: 3
	//i: 5
	//i: 7
	//i: 9
}

// 9.猴子吃桃子问题：
// 有一堆桃子，猴子第一天吃一半，并多吃了一个，以后每天猴子都吃剩余的一半，并多吃一个。
// 当第10天时，再想吃时（还没吃），发现只有1个桃子了，问最初有多少桃子？
func TestGetPeachNum(t *testing.T) {
	fmt.Printf("getPeachNum:%v \n", getPeachNum(1))
	// getPeachNum:1534
}
