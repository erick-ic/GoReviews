package closure

import (
	"fmt"
	"strings"
)

// 累加器
func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
	//闭包内容：变量n与以下代码块构成闭包
	/*
		var n int = 10
		return func(x int) int {
			n = n + x
			return n
		}
	*/
}

func makeSuffix(suffix string) func(string) string {
	//若没有指定后缀，则加上，否则返回原名
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
	//闭包内容：变量suffix与以下代码块构成闭包
	/*
		return func(name string) string {
			if !strings.HasSuffix(name, suffix) {
				return name + suffix
			}
			return name
		}
	*/
}

func DeferClosureLoopV1() {
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func DeferClosureLoopV2() {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			fmt.Println(val)
		}(i)
	}
}

func DeferClosureLoopV3() {
	var j int
	for i := 0; i < 10; i++ {
		j = i
		defer func() {
			fmt.Println(j)
		}()
	}
}
