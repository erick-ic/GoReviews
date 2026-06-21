package main

import "fmt"

func main() {
	////1.两个变量，取值互换
	//a := 1
	//b := 2
	//
	////temp := a
	////a = b
	////b = temp
	////fmt.Printf("a: %d, b: %d \n", a, b)
	//////a: 2, b: 1
	//
	//a = a + b
	//b = a - b // b = a + b - b => b = a
	//a = a - b // a = a + b - a => a = b
	//fmt.Printf("a: %d, b: %d\n", a, b)
	////a: 2, b: 1

	////2.打印1-100之间所有9的倍数的整数的个数，总和
	//var count uint64 = 0
	//var sum uint64 = 0
	//var i uint64 = 1
	//for ; i < 100; i++ {
	//	if i%9 == 0 {
	//		count++
	//		sum += i
	//	}
	//}
	//fmt.Printf("count: %d, sum: %d \n", count, sum)
	////count: 11, sum: 594

	////3.按要求打印下列数字
	///*
	//	0 + 6 = 6
	//	1 + 5 = 6
	//	2 + 4 = 6
	//	3 + 3 = 6
	//	4 + 2 = 6
	//	5 + 1 = 6
	//	6 + 0 = 6
	//*/
	//var target int = 6
	//for i := 0; i <= 6; i++ {
	//	fmt.Printf("%d + %d = %d \n", i, target-i, target)
	//}
	////0 + 6 = 6
	////1 + 5 = 6
	////2 + 4 = 6
	////3 + 3 = 6
	////4 + 2 = 6
	////5 + 1 = 6
	////6 + 0 = 6

	////4.统计3个班成绩情况，每个班有5名学生，求出各个班级的平均分和所有班级的平均分
	//var class_num int = 3
	//var student_num int = 5
	//var all_class_sum float64 = 0
	////统计3个班成绩情况，j 代表班级
	//for j := 1; j <= class_num; j++ {
	//	var class_sum float64 = 0
	//	//统计一个班的成绩情况, i 代表学生
	//	for i := 1; i <= 5; i++ {
	//		var score float64
	//		fmt.Printf("请输入第%d个班级第%d个学生成绩：", j, i)
	//		_, err := fmt.Scanln(&score)
	//		if err != nil {
	//			return
	//		}
	//		fmt.Println(score)
	//		class_sum += score
	//	}
	//	fmt.Printf("第%v个班级平均分：%v \n", j, class_sum/float64(student_num))
	//	all_class_sum += class_sum
	//}
	//fmt.Printf("所有班级的总成绩：%v, 平均分：%v \n", all_class_sum, all_class_sum/float64((class_num*student_num)))

	////5.打印金字塔：编写一个程序，接收一个整数表示金字塔的层数。
	//
	////思路：打印矩形
	////var x int = 3
	////var y int = 4
	//
	////for j := 1; j <= y; j++ {
	////	for i := 1; i <= x; i++ {
	////		fmt.Print("*")
	////	}
	////	fmt.Println()
	////}
	////***
	////***
	////***
	////***
	//
	////思路：打印半个金字塔
	////第一层：1
	////第二层：2
	////第三层：3
	////第n层：n
	////var num int = 3
	//
	////for i := 1; i <= num; i++ {
	////  //i 当前的层数
	////	//j 每层打印的*
	////	for j := 1; j <= i; j++ {
	////		fmt.Print("*")
	////	}
	////	fmt.Println()
	////}
	////*
	////**
	////***
	//
	////思路：打印完整的金字塔
	////第一层：1个*，2个空格
	////第二层：3个*，1个空格
	////第三层：5个*，0个空格
	////第n层：2n - 1个*, 空格：总层数-当前层数
	////var num2 int = 3
	//
	////for i := 1; i <= num2; i++ {
	////	//i 当前的层数
	////	//j 每层要打印的*
	////	for k := 1; k <= num2-i; k++ {
	////		fmt.Print(" ")
	////	}
	////	for j := 1; j <= 2*i-1; j++ {
	////		fmt.Print("*")
	////	}
	////	fmt.Println()
	////}
	////  //  *
	////	// ***
	////	//*****
	//
	////思路：打印空心金字塔
	////规律：每层第一个和最后一个打印*，其余为空格
	//var floor int = 5
	//for i := 1; i <= floor; i++ {
	//
	//	for k := 1; k <= floor-i; k++ {
	//		fmt.Print(" ")
	//	}
	//	for j := 1; j <= 2*i-1; j++ {
	//		//每层第一个和最后一个打印*，其余为空格
	//		if j == 1 || j == 2*i-1 || i == floor {
	//			fmt.Print("*")
	//		} else {
	//			fmt.Print(" ")
	//		}
	//	}
	//	fmt.Println()
	//}
	////    *
	////   * *
	////  *   *
	//// *     *
	////*********

	////6.打印九九乘法表
	//var num int = 9
	//for i := 1; i <= num; i++ {
	//	//i 表示层数
	//	//j 表示当前个数
	//	for j := 1; j <= i; j++ {
	//		fmt.Printf("%v * %v = %v \t", j, i, i*j)
	//	}
	//	fmt.Println()
	//}
	//1 * 1 = 1
	//1 * 2 = 2 	2 * 2 = 4
	//1 * 3 = 3 	2 * 3 = 6 	3 * 3 = 9
	//1 * 4 = 4 	2 * 4 = 8 	3 * 4 = 12 	4 * 4 = 16
	//1 * 5 = 5 	2 * 5 = 10 	3 * 5 = 15 	4 * 5 = 20 	5 * 5 = 25
	//1 * 6 = 6 	2 * 6 = 12 	3 * 6 = 18 	4 * 6 = 24 	5 * 6 = 30 	6 * 6 = 36
	//1 * 7 = 7 	2 * 7 = 14 	3 * 7 = 21 	4 * 7 = 28 	5 * 7 = 35 	6 * 7 = 42 	7 * 7 = 49
	//1 * 8 = 8 	2 * 8 = 16 	3 * 8 = 24 	4 * 8 = 32 	5 * 8 = 40 	6 * 8 = 48 	7 * 8 = 56 	8 * 8 = 64
	//1 * 9 = 9 	2 * 9 = 18 	3 * 9 = 27 	4 * 9 = 36 	5 * 9 = 45 	6 * 9 = 54 	7 * 9 = 63 	8 * 9 = 72 	9 * 9 = 81

	////7.100以内数求和，求出当和第一次大于20的当前数
	//sum := 0
	//for i := 1; i <= 100; i++ {
	//	sum += i
	//	if sum > 20 {
	//		fmt.Println("i:", i)
	//		fmt.Println("sum:", sum)
	//		break
	//	}
	//}

	//8.10以内的奇数-for循环+continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("i:", i)
	}
	//i: 1
	//i: 3
	//i: 5
	//i: 7
	//i: 9
}
