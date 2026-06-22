package main

import "fmt"

func fbn(n int) []uint64 {
	//切片大小为n
	//当 n=1 时，只有 fbnSlice[0] 合法，而 fbnSlice[1] 非法，
	//因此触发 index out of range [1] with length 1。
	if n <= 0 {
		return []uint64{}
	}
	if n == 1 {
		return []uint64{1}
	}
	fbnSlice := make([]uint64, n)
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}

func main() {
	//切片是引用类型，在进行传递时，遵循引用传递机制。切片是一个数组的引用。
	//切片的长度是可变的，可以看作一个动态变化数组。
	//元素：任意类型，但不能混用。
	//组成：
	/*
		切片：一个三字段的结构体。
		type sliceHeader struct {
			Data uintptr // 指向底层数组的指针
			Len  int     // 切片长度
			Cap  int     // 切片容量
		}
			Data：指向底层数组的起始内存地址，切片本身不存储数据。
			Len：表示当前切片长度，决定了通过切片访问多少元素。
			Cap：表示切片容量，表示从起始地址到底层数组末尾最多能访问多少元素。
	*/
	//定义方式：var 切片名 []类型
	var slice []int
	fmt.Printf("%v, %T, len = %v\n", slice, slice, len(slice))
	//[], []int, len = 0

	//(1)基本语法：
	//方式1:
	var slice1 = []int{1, 2, 3}
	fmt.Printf("%v, %T, len = %v\n", slice1, slice1, len(slice1))
	//[1 2 3], []int, len = 3

	//方式2:
	var slice2 = []int{0: 1, 1: 11, 3: 33}
	fmt.Printf("%v, %T, len = %v\n", slice2, slice2, len(slice2))
	//[1 11 0 33], []int, len = 4

	//方式3:
	//长度：元素个数
	//容量：从第1个元素到其底层数组元素末尾的个数
	var slice3 = make([]int, 6)
	fmt.Printf("%v, %T\n", slice3, slice3)
	//[0 0 0 0 0 0], []int
	fmt.Printf("长度为%d, 容量为%d\n", len(slice3), cap(slice3))
	//长度为6, 容量为6

	//(3)常用方法：
	//(3.1)切片截取
	slice4 := []int{1, 2, 3, 4, 5}
	slice44 := slice4[:]
	fmt.Println("slice44:", slice44)
	//slice44: [1 2 3 4 5]
	slice44 = slice4[:2]
	fmt.Println("slice44:", slice44)
	//slice44: [1 2]

	//(3.2)添加操作:
	//长度增加：增加元素个数
	//容量增加：元素个数大于容量时，容量翻倍，即旧容量的2倍（一般情况）
	var sliceA []int
	fmt.Printf("%v, %T\n", sliceA, sliceA)           // [], []int
	fmt.Printf("%v, %v\n", len(sliceA), cap(sliceA)) // 0, 0

	sliceA = append(sliceA, 21)
	fmt.Printf("%v, %T\n", sliceA, sliceA)           // [21], []int
	fmt.Printf("%v, %v\n", len(sliceA), cap(sliceA)) // 1, 1

	sliceA = append(sliceA, 21)
	fmt.Printf("%v, %T\n", sliceA, sliceA)           // [21 21], []int
	fmt.Printf("%v, %v\n", len(sliceA), cap(sliceA)) // 2, 2

	sliceA = append(sliceA, 21)
	fmt.Printf("%v, %T\n", sliceA, sliceA)           // [21 21 21], []int
	fmt.Printf("%v, %v\n", len(sliceA), cap(sliceA)) // 3, 4

	//(3.3)合并操作：
	sliceB := []int{1, 2, 3}
	sliceC := []int{4, 5}
	sliceD := append(sliceB, sliceC...)
	fmt.Printf("%v, %T\n", sliceD, sliceD)           //[1 2 3 4 5], []int
	fmt.Printf("%v, %v\n", len(sliceD), cap(sliceD)) //5, 6

	//(3.4)删除操作：
	//go中没有删除切片元素的专用方法，通过切片截取的方式实现。
	sliceE := []int{1, 2, 3, 4, 5, 6}
	//删除索引为2的元素，即3
	sliceE = append(sliceE[:2], sliceE[3:]...)
	fmt.Printf("%v, %T\n", sliceE, sliceE) //[1 2 4 5 6], []int

	//(3.5)修改操作：
	sliceF := []int{1, 2, 3}
	sliceF[0] = 100
	fmt.Printf("%v, %T\n", sliceF, sliceF) //[100 2 3], []int

	//(3.6)复制操作：
	sliceG := []int{1, 2, 3}
	// slice属于引用类型数据，二者共同指向同一存储空间
	sliceH := sliceG
	// 执行此操作会修改存储空间的值，影响原切片（sliceG）
	sliceH[0] = 11
	fmt.Printf("%v, %T\n", sliceG, sliceG) //[11 2 3], []int
	fmt.Printf("%v, %T\n", sliceH, sliceH) //[11 2 3], []int

	//copy方法
	sliceM := []int{1, 2, 3, 4, 5}
	sliceN := make([]int, 4, 4)
	copy(sliceN, sliceM)
	fmt.Printf("%v, %T\n", sliceM, sliceM) //[1 2 3 4 5], []int
	fmt.Printf("%v, %T\n", sliceN, sliceN) //[1 2 3 4], []int

	//(4)遍历切片：
	sliceLoop := []int{1, 2, 3}
	//for i := 0; i < len(sliceLoop); i++ {
	//	fmt.Printf("sliceLoop[%d] = %d\n", i, sliceLoop[i])
	//	//sliceLoop[0] = 1
	//	//sliceLoop[1] = 2
	//	//sliceLoop[2] = 3
	//}

	for index, val := range sliceLoop {
		fmt.Printf("sliceLoop[%d] => %d\n", index, val)
		//sliceLoop[0] => 1
		//sliceLoop[1] => 2
		//sliceLoop[2] => 3
	}

	//(5)string和切片
	//(5.1)字符串截取
	//string底层是一个byte数组，因此string可以进行切片处理
	str1 := "hello"
	strSlice := str1[:2]
	fmt.Printf("%v, %T\n", strSlice, strSlice)
	//he, string

	//(5.2)字符串修改
	//string是不可变的，无法通过str[0] = 'z' 方式修改字符串
	//str1[0] = 'z'
	strArr := []byte(str1)
	strArr[0] = 'Z'
	fmt.Printf("%v, %T\n", strArr, strArr)
	//[90 101 108 108 111], []uint8
	targetStr := string(strArr)
	fmt.Printf("%v, %T\n", targetStr, targetStr)
	//Zello, string

	//案例1: 编写一个函数fbn(n int), 返回对应完整的斐波那契数组合。
	/*
		要求：
			可以接收一个参数n int,
			将斐波那契数的数据放在切片中
			提示：arr[0] = 1, arr[1] = 1, arr[2] = 2, arr[3] = 3, arr[4] = 5...
	*/
	fbnSlice := fbn(3)
	fmt.Printf("%v, %T\n", fbnSlice, fbnSlice)
	//[1 1 2], []uint64
}
