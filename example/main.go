package example

import "fmt"

func getPeachNum(d int) int {
	//day10: 1
	//day9: (day10 + 1) * 2
	if d > 10 || d < 0 {
		fmt.Println("getPeachNum fail")
		return 0
	}
	if d == 10 {
		return 1
	} else {
		return (getPeachNum(d+1) + 1) * 2
	}

}

func swap(n1, n2 *int) {
	t := *n1
	*n1 = *n2
	*n2 = t
}

func exchange(a, b int) (int, int) {
	////方式1:
	//temp := a
	//a = b
	//b = temp
	//return a, b

	////方式2:
	//a = a + b
	//b = a - b // ==> b = a + b - b = a
	//a = a - b // ==> a = a + b - a = b
	//return a, b

	//方式3:
	a, b = b, a
	return a, b
}

func calculateNine(target, num int) (count int, sum int) {
	for i := 1; i < num; i++ {
		if i%target == 0 {
			count++
			sum += i
		}
	}
	return count, sum
}

func handlePrint(target int) {
	for i := 0; i <= target; i++ {
		fmt.Printf("%d + %d = %d \n", i, target-i, target)
	}
}

func handleClass(classNum, studentNum int) (classAvg, allClassScore float64) {
	//班级情况
	for i := 1; i <= classNum; i++ {
		var classScore float64
		//每个班级的学生情况
		for j := 1; j <= studentNum; j++ {
			var studentScore float64
			fmt.Printf("请输入第%d个班级第%d个学生成绩：", i, j)
			_, err := fmt.Scanln(&studentScore)
			if err != nil {
				return
			}
			fmt.Println(studentScore)
			classScore += studentScore
		}
		classAvg = classScore / float64(studentNum)
		fmt.Printf("第%v个班级平均分：%v \n", i, classAvg)
		allClassScore += classScore
	}
	allClassAvg := allClassScore / (float64(classNum * studentNum))
	return allClassAvg, allClassScore
}

func printGoldTower(n int) {
	////思路：打印矩形：4*3
	//for i := 0; i < 3; i++ {
	//	for j := 0; j < 4; j++ {
	//		fmt.Print("*")
	//	}
	//	fmt.Println()
	//}
	////****
	////****
	////****

	////思路：打印半个金字塔
	////第一层：1
	////第二层：2
	////第三层：3
	////第n层：n
	//for i := 1; i <= n; i++ {
	//	//i 当前的层数
	//	//j 每层打印的*
	//	for j := 1; j <= i; j++ {
	//		fmt.Printf("*")
	//	}
	//	fmt.Printf("\n")
	//}
	////*
	////**
	////***

	////思路：打印完整的金字塔
	////第一层：1个*，2个空格
	////第二层：3个*，1个空格
	////第三层：5个*，0个空格
	////第n层：2n - 1个*, 空格：总层数-当前层数
	//for i := 1; i <= n; i++ {
	//	//i 当前的层数
	//	//j 每层要打印的*
	//	//打印空格
	//	for k := 1; k <= n-i; k++ {
	//		fmt.Print(" ")
	//	}
	//	//打印*
	//	for j := 1; j <= 2*i-1; j++ {
	//		fmt.Print("*")
	//	}
	//	fmt.Println()
	//}
	////  *
	//// ***
	////*****

	//思路：打印空心金字塔
	//规律：每层第一个和最后一个打印*，其余为空格
	for i := 1; i <= n; i++ {

		for k := 1; k <= n-i; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			//每层第一个和最后一个打印*，其余为空格
			if j == 1 || j == 2*i-1 || i == n {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	//  *
	// * *
	//*****

}

func multiTable(n int) {
	for i := 1; i <= n; i++ {
		//i 表示层数
		//j 表示当前个数
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d \t", j, i, j*i)
		}
		fmt.Println()
	}
}

func getSumWithinNum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
		if sum > 20 {
			fmt.Println("i:", i)
			fmt.Println("sum:", sum)
			break
		}
	}
	return sum
}

func findOddNumber(n int) {
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("i:", i)
	}
}
func main() {
	//10.编写函数swap(n1 *int, n2 *int) 可以交换n1、n2的值
	a := 10
	b := 20
	swap(&a, &b)
	fmt.Printf("a: %v, b: %v \n", a, b)
	//a: 20, b: 10
}
