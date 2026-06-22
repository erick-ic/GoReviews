package main

import (
	"fmt"
	"strconv"
	"time"
)

const (
	Nanosecond  = 1                  //纳秒
	Microsecond = 1000               //微秒
	Millisecond = 1000 * Microsecond //毫秒
	Second      = 1000 * Millisecond //秒
	Minute      = 60 * Second        //分
	Hour        = 60 * Minute        //时
)

func main() {
	//(1)获取当前时间
	now := time.Now()
	fmt.Printf("now = %v, type = %T\n", now, now)
	//now = 2026-06-22 11:24:36.224215 +0800 CST m=+0.000115126, type = time.Time

	//(2)获取年月日时分秒
	fmt.Println("年：", now.Year())
	//年： 2026
	fmt.Println("月：", now.Month())
	//月： June
	fmt.Println("月：", int(now.Month()))
	//月： 6
	fmt.Println("日：", now.Day())
	//日： 22
	fmt.Println("时：", now.Hour())
	//时： 11
	fmt.Println("分：", now.Minute())
	//分： 28
	fmt.Println("秒：", now.Second())
	//秒： 23

	//(3)格式化日期时间
	//方式1: Printf、Sprintf
	fmt.Printf("当前日期：%d-%d-%d %d:%d:%d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	//当前日期：2026-6-22 11:32:40
	dateStr := fmt.Sprintf("当前日期：%d-%d-%d %d:%d:%d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println("dateStr:", dateStr)
	//dateStr: 当前日期：2026-6-22 11:34:36

	//方式2: time.Format()
	fmt.Println("当前日期：", now.Format("2006-01-02 15:04:05"))
	//当前日期： 2026-06-22 11:36:44
	fmt.Println(now.Format("2006-01-02"))
	//2026-06-22
	fmt.Println(now.Format("15:04:05"))
	//11:38:16

	//(4)时间戳
	//Unix表示从1970-01-01到现在所经过的时间（单位秒）
	//Unixnano表示从1970-01-01到现在所经过的时间（单位纳秒）
	fmt.Printf("unix时间戳=%v, uinxnano时间戳=%v\n", now.Unix(), now.UnixNano())
	//unix时间戳=1782100727, uinxnano时间戳=1782100727833956000

	////案例1：
	////每隔1秒打印一个数字，打印到5退出
	//i := 0
	//for {
	//	i++
	//	//每次打印前重新获取当前时间，把 time.Now() 放到循环内部即可
	//	fmt.Printf("当前时间：%v, 数字：%v \n", time.Now().Format("15:04:05"), i)
	//	time.Sleep(time.Second * 1)
	//	if i == 5 {
	//		break
	//	}
	//}

	//案例2:
	//统计函数的执行时间

	//需要重新获取当前时间
	start := time.Now().Unix()
	test()
	end := time.Now().Unix()
	fmt.Println("函数执行时间：", end-start)
	//函数执行时间： 2
}

func test() {
	str := "golang"
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}
