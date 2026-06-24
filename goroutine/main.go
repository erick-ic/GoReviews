package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// testRoutine 协程
func testRoutine() {
	for i := 0; i <= 10; i++ {
		fmt.Println("hello, testRoutine" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
}

// testThread 线程
func testThread() {
	for i := 0; i <= 3; i++ {
		fmt.Println("hello, testThread" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
}

var (
	myMap = make(map[int]int, 10)
	//声明全局的互斥锁
	lock sync.Mutex
)

func main() {
	//进程：程序在操作系统中的一次执行过程，是系统进行资源分配、调度的基本单位。
	//线程：进程的一个执行实例，程序执行的最小单位。
	//一个进程可以创建、核销多个线程，同一个进程中多个线程可以并发执行。
	//一个程序至少一个进程，一个进程至少一个线程
	//并发：多线程程序在单核上运行。多个任务作用在单个cpu，从微观角度来看，一个时间点上，只有一个任务在执行。
	//并行：多线程程序在多核上运行。多个任务作用在多个CPU，从微观角度来看，一个时间点上，多个任务在同时执行。

	//一个Go主线程上可以起多个协程
	//Go协程的特点： 独立的栈空间、共享程序堆空间、调度由用户控制、协程是轻量级的线程。

	//(1)基本操作：
	//(1.1)main主线程和协程同时进行
	//若主线程退出，即使协程未执行完成也会退出

	////开启协程
	//go testRoutine()
	////主线程
	//testThread()
	////hello, testThread0
	////hello, testRoutine0
	////hello, testThread1
	////hello, testRoutine1
	////hello, testRoutine2
	////hello, testThread2
	////hello, testThread3
	////hello, testRoutine3
	////hello, testRoutine4
	////结束

	////(1.2)设置Go运行的CPU数 (GO1.8+ 不需要单独设置，默认多线程)
	////获取设备的逻辑CPU数，不一定是真实的
	//cpuNum := runtime.NumCPU()
	//fmt.Println("cpuNum:", cpuNum)
	////cpuNum: 18
	//
	////设置同时执行的最大CPU数
	//runtime.GOMAXPROCS(cpuNum)

	//(2)MPG调度模型
	//M：操作系统的主线程（物理线程）
	//P：协程执行的上下文（运行资源）
	//G：协程

	//案例1：计算1-20的阶乘，并把各个数的阶乘放入map中，即key-value形式，goroutine方式实现。
	//出现的问题1: 存在资源竞争的问题
	//出现的问题2: 无法确定主线程结束时间

	////问题1 解决方式1: 存在资源竞争的问题。
	//for i := 1; i <= 20; i++ {
	//	go calcuFactorial1(i)
	//}

	//问题1 解决方式2：加入互斥锁，解决资源竞争问题
	for i := 1; i <= 20; i++ {
		go calcuFactorial2(i)
	}

	//问题2 解决方式1: 等协程结束，再关闭主线程。（临时设置）
	time.Sleep(5 * time.Second)

	lock.Lock()
	for k, v := range myMap {
		fmt.Printf("map[%d] = %d\n", k, v)
	}
	lock.Unlock()
}

func calcuFactorial1(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//缺乏互斥保护机制，存在资源竞争
	myMap[n] = res
	//fatal error: concurrent map writes
	//编译时，添加-race参数可查看具体情况，即go build -race main.go
	//查看：./main
	//资源竞争报错：WARNING: DATA RACE
}

func calcuFactorial2(n int) {
	res := 1
	for i := 1; i < n; i++ {
		res *= i
	}

	//加上互斥锁
	lock.Lock()
	myMap[n] = res
	//解锁
	lock.Unlock()
}
