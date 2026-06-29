package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// writeData:
func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
	}
	close(intChan)
}

// readData:
func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("readData: ", v)
	}
	exitChan <- true
	close(exitChan)
}

// 写入全部数据
func putNums(intChan chan int) {
	for i := 1; i <= 100; i++ {
		intChan <- i
	}
	close(intChan)
}

// 读取全部数据，并判断素数
func primeNums(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		n, ok := <-intChan
		if !ok {
			break
		}

		if n <= 1 {
			//1不是素数，直接跳过
			continue
		}

		flag = true
		//判断素数，只能被1和自身整除。
		//遍历 [2, n-1] 这个区间，if n%i == 0 就是检查“能否被整除”，如果通过则不是素数。
		for i := 2; i < n; i++ {
			if n%i == 0 {
				flag = false
				break
			}
		}
		//写入素数
		if flag {
			primeChan <- n
		}
	}
	//此时还不能关闭primeChan，其他的primeChan可能还在执行
	exitChan <- true
}

func main() {
	//channel本质是一个数据结构：队列
	//只能存放同类型的数据，如string类型的channel只能存放string。
	//数据先进先出【FIFO：first in first out】
	//线程安全，多个goroutine访问时，无需加锁。
	//类型：引用类型
	//定义方式：var 变量名 chan 数据类型
	//make初始化后使用
	var intChan chan int
	fmt.Printf("%v, %T \n", intChan, intChan)
	//<nil>, chan int

	//================================(1)基本使用================================
	//(1.1)创建channel:
	//一个存放3个string类型的channel
	var strChan chan string
	strChan = make(chan string, 3)
	fmt.Printf("值：%v, 地址：%p \n", strChan, &strChan)
	//值：0x1642d0712070, 地址：0x1642d0702040

	//(1.2)写入数据:
	strChan <- "s1"
	strChan <- "s2"
	strChan <- "s3"

	//放弃存入"s3"
	//<- strChan

	//strChan <- "s4"
	//// 写入数据不能超出其容量  fatal error: all goroutines are asleep - deadlock!
	fmt.Printf("value:%v, len: %d, cap: %d \n", strChan, len(strChan), cap(strChan))
	//value:0x4923cef30070, len: 3, cap: 3

	//(1.3)读取数据: 取出
	res1 := <-strChan
	res2 := <-strChan
	res3 := <-strChan
	////数据取完后不能再取 //fatal error: all goroutines are asleep - deadlock!
	//res4 := <-strChan
	fmt.Printf("res1:%v - %T, res2:%v - %T, res3:%v - %T \n", res1, res1, res2, res2, res3, res3)
	//res1:s1 - string, res2:s2 - string, res3:s3 - string

	//(1.4)关闭channel:
	str1Chan := make(chan string, 3)
	str1Chan <- "s1"
	str1Chan <- "s2"
	//关闭channel
	close(str1Chan)
	//关闭后无法写入
	////panic: send on closed channel
	//str1Chan <- "s3"

	//关闭后可以取出
	str1Data := <-str1Chan
	fmt.Printf("str1Data:%v - %T\n", str1Data, str1Data)
	//str1Data:s1 - string

	//(1.5)channel读写声明
	//默认可读可写：var chan1 chan int
	//只写：
	var chan2 chan<- int
	chan2 = make(chan int, 3)
	chan2 <- 1
	//chan2Data := <-chan2
	////invalid operation: cannot receive from send-only channel chan<- int chan2 (variable of type chan<- int)

	////只读：
	//var chan3 <-chan int
	//chan3Data := <-chan3

	//================================(2)各种类型的channel================================
	//(2.1)map
	var mapChan chan map[string]string
	mapChan = make(chan map[string]string, 10)
	m1 := make(map[string]string, 20)
	m1["name"] = "Jack"
	m2 := make(map[string]string, 20)
	m2["name"] = "Tome"

	mapChan <- m1
	mapChan <- m2

	m22 := <-mapChan
	m33 := <-mapChan
	fmt.Println(m22, m33)
	//map[name:Jack] map[name:Tome]

	//(2.2)struct
	var userChan chan User
	userChan = make(chan User, 10)
	u1 := User{Name: "Jack", Age: 21}
	u2 := User{Name: "Tome", Age: 18}

	userChan <- u1
	userChan <- u2

	u11 := <-userChan
	u22 := <-userChan
	fmt.Println(u11, u22)
	//{Jack 21} {Tome 18}

	//(2.3)interface 任意类型
	//var allChan chan interface{}
	//allChan = make(chan interface{}, 10)
	allChan := make(chan interface{}, 10)

	allChan <- "hello"
	allChan <- 21
	allChan <- true
	allChan <- []int{1, 2, 3}
	allChan <- map[string]int{"a": 1, "b": 2}
	allChan <- u1

	//想要取到目标元素，需要先将取前方的元素推出
	//默认得到的是第一个元素
	//allData := <-allChan
	//fmt.Printf("allData: %v, type: %T \n", allData, allData)
	////allData: hello, type: string

	////推出
	//<-allChan
	//allData := <-allChan
	//fmt.Printf("allData: %v, type: %T \n", allData, allData)
	////allData: 21, type: int
	<-allChan
	<-allChan
	<-allChan
	<-allChan
	<-allChan
	allData := <-allChan
	fmt.Printf("allData: %v, type: %T \n", allData, allData)
	//allData: {Jack 21}, type: main.User
	//运行的type为User，编译时仍为Interface
	newUser := allData.(User)
	fmt.Println(newUser.Name, newUser.Age)
	//Jack 21

	//================================(3)channel遍历================================
	newIntChan := make(chan int, 5)
	newIntChan <- 1
	newIntChan <- 2
	newIntChan <- 3
	newIntChan <- 4
	newIntChan <- 5

	//遍历时，若channel未关闭，则会出现deadlock错误
	//正常遍历，但会报错
	//fatal error: all goroutines are asleep - deadlock!

	//遍历时，若channel已关闭，则正常遍历，完成后推出。
	close(newIntChan)
	for v := range newIntChan {
		fmt.Println("v = ", v)
	}

	////案例1:
	///*
	//	开启一个writeData协程，向管道intChan中写入50个数
	//	开启一个readData协程，从管道intChan中取出数据。
	//	两个协程操作的是同意channel
	//	主线程需要等writeData、readData协程完成后结束
	//*/
	//intChan2 := make(chan int, 50)
	//exitChan := make(chan bool, 1)
	//
	//go writeData(intChan2)
	//go readData(intChan2, exitChan)
	//
	////取完结束主线程
	//for {
	//	_, ok := <-exitChan
	//	if !ok {
	//		break
	//	}
	//}

	//案例2: 统计1-100中的素数(质数)。
	//素数（质数）是指在大于 1 的自然数中，除了 1 和它本身以外不再有其他因数的数。
	//全部数据管道
	intChan3 := make(chan int, 1000)
	//素数管道
	primeChan := make(chan int, 2000)
	//标识退出管道
	exitChan3 := make(chan bool, 4)

	//开启1个协程写入1-8000个数
	go putNums(intChan3)

	//开启4个协程取出全部数据，判断是否为素数，若是，则写入primeChan
	for i := 0; i < 4; i++ {
		go primeNums(intChan3, primeChan, exitChan3)
	}

	//主线程
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan3
		}
		//4个协程结束，关闭primeChan
		close(primeChan)
	}()

	//遍历结果
	for {
		data, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println("素数: ", data)
	}
	fmt.Println("主线程退出...")
}
