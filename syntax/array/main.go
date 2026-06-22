package main

import "fmt"

func test1(arr [3]int) {
	arr[0] = 100
}

func test2(arr *[3]int) {
	(*arr)[0] = 100
}

func main() {
	//数组：存放多个同一类型的元素，一旦定义，其长度固定，无法动态变化。
	//数组是连续的内存块，元素按顺序紧密排列，编译时大小固定。
	//元素：任意类型，但不能混用。
	//类型：值类型。
	//组成：长度，索引，值
	//定义方式：var 数组名 [数组大小]类型

	//数组的地址=数组第1个元素的地址
	//第2个元素地址=第1个元素地址+8（int占8个字节）
	list := [3]int{1, 2, 3}
	fmt.Printf("数组地址：%p \n", &list)
	fmt.Printf("第1个元素地址：%p \n", &list[0])
	fmt.Printf("第2个元素地址：%p \n", &list[1])
	//数组地址：0x620a6f446090
	//第1个元素地址：0x620a6f446090
	//第2个元素地址：0x620a6f446098

	//(1)基本语法：
	//方式1: 先声明后赋值
	var arr1 [3]int
	fmt.Printf("arr1: %v, type: %T \n", arr1, arr1)
	//arr1: [0 0 0], type: [3]int
	arr1 = [3]int{1, 2, 3}
	arr1[0] = 11
	fmt.Printf("arr1: %v, type: %T \n", arr1, arr1)
	//arr1: [11 2 3], type: [3]int

	//方式2: 字面量语法
	var arr2 = [3]int{1, 2, 3}
	fmt.Printf("arr2: %v, type: %T \n", arr2, arr2)
	//arr2: [1 2 3], type: [3]int

	//方式3: 不定长度
	//根据初始值的个数自动推断数组长度
	var arr3 = [...]int{1, 2, 3, 4, 5}
	fmt.Printf("arr3: %v, type: %T \n", arr3, arr3)
	//arr3: [1 2 3 4 5], type: [5]int

	//方式4: 指定元素值对应的索引
	var arr4 = [...]int{1: 10, 0: 0, 3: 30}
	fmt.Printf("arr4: %v, type: %T \n", arr4, arr4)
	//arr4: [0 10 0 30], type: [4]int

	//(2)遍历数组：
	arr5 := [...]int{1, 2, 3}
	//方式1:
	for i := 0; i < len(arr5); i++ {
		fmt.Printf("arr5-%v: %v \n", i, arr5[i])
	}
	//arr5-0: 1
	//arr5-1: 2
	//arr5-2: 3

	//方式2:
	for inx, val := range arr5 {
		fmt.Printf("arr5->%v: %v \n", inx, val)
	}
	//arr5->0: 1
	//arr5->1: 2
	//arr5->2: 3

	//(3)注意点：
	//数组属于值类型，默认是值传递。进行拷贝时，数组之间不会相互影响。
	arr6 := [...]int{11, 22, 33}
	test1(arr6)
	fmt.Println("arr6:", arr6)
	//arr6: [11 22 33]

	//若想在其他函数中修改原数组，使用引用传递（指针方式）。
	test2(&arr6)
	fmt.Println("arr6:", arr6)
	//arr6: [100 22 33]

	//(4)常用方法：
	//(4.1)数组截取：
	arr8 := [...]int{1, 2, 3, 4, 5}
	arr88 := arr8[:]
	fmt.Println("arr88:", arr88)
	//arr88: [1 2 3 4 5]
	arr88 = arr8[:3]
	fmt.Println("arr88:", arr88)
	//arr88: [1 2 3]
	arr88 = arr8[3:]
	fmt.Println("arr88:", arr88)
	//arr88: [4 5]
	arr88 = arr8[1:3]
	fmt.Println("arr88:", arr88)
	//arr88: [2 3]

	////案例1：创建一个byte类型的26个元素的数组，分别放置“A-Z”，打印所有元素。
	////提示：字符数据运算'A'+1='B'
	//var charsArr [26]byte
	////for i := 0; i < 26; i++ {
	////	charsArr[i] = 'A' + byte(i)
	////	fmt.Printf("charsArr-%v: %c \n", i+1, charsArr[i])
	////	//charsArr-1: A
	////	//charsArr-2: B
	////	//charsArr-3: C
	////}
	////fmt.Printf("charsArr: %c \n", charsArr)
	//////charsArr: [A B C D E F G H I J K L M N O P Q R S T U V W X Y Z]
	//
	//for i, v := range charsArr {
	//	charsArr[i] = 'A' + byte(i)
	//	fmt.Printf("charsArr->%v: %c, v: %v \n", i+1, charsArr[i], v)
	//	//charsArr->1: A, v: 0
	//	//charsArr->2: B, v: 0
	//	//charsArr->3: C, v: 0
	//	//...
	//}
	////fmt.Printf("charsArr: %c \n", charsArr)
	//////charsArr: [A B C D E F G H I J K L M N O P Q R S T U V W X Y Z]
	//
	////案例2: 求数组的最大值，并找到对应下标。
	//arr7 := [...]int{1, 2, 3, 4, 5}
	//maxVal := arr7[0]
	//maxValIndex := 0
	//for i, _ := range arr7 {
	//	if arr7[i] > maxVal {
	//		maxVal = arr7[i]
	//		maxValIndex = i
	//	}
	//}
	//fmt.Printf("maxVal:%v, maxValIndex: %v \n", maxVal, maxValIndex)
	////maxVal:5, maxValIndex: 4
	//
	//// 案例3：从数组[1,3,5,7,8]中找到和为8的两个元素的下标
	//var numList = [...]int{1, 3, 5, 7, 8}
	//for i := 0; i < len(numList); i++ {
	//	// i+1 内层循环体从外层循环体之后的索引开始查找
	//	for j := i + 1; j < len(numList); j++ {
	//		if numList[i]+numList[j] == 8 {
	//			fmt.Printf("(%v, %v)\n", i, j)
	//			// (0, 3)
	//			// (1, 2)
	//		}
	//	}
	//}
}
