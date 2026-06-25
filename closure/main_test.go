package closure

import (
	"fmt"
	"testing"
)

//闭包定义：一个函数和与其相关的引用环境组成的整体
//闭包内容：返回一个匿名函数，且引用了匿名函数外的变量(n)。
//变量只会初始化一次，后续调用只会在原基础上叠加

func TestAddUpper(t *testing.T) {
	f := AddUpper()
	fmt.Println(f(10)) //20
	fmt.Println(f(20)) //40
	fmt.Println(f(30)) //70
}

// 该实践中闭包的优势：后缀只需传入一次，闭包可以保留上次引用的某个值，便于重复使用。
func TestMakeSuffix(t *testing.T) {
	f := makeSuffix(".jpg")
	fmt.Println(f("winter"))     // winter.jpg
	fmt.Println(f("spring.jpg")) //spring.jpg
}

func TestDeferClosureLoopV1(t *testing.T) {
	//每次迭代都会 重新声明一个新的 i 变量（作用域为本次循环体），
	//每个 defer 闭包捕获的是本次迭代独立的 i
	DeferClosureLoopV1()
	//9
	//8
	//7
	//6
	//5
	//4
	//3
	//2
	//1
	//0
}

func TestDeferClosureLoopV2(t *testing.T) {
	DeferClosureLoopV2()
	//9
	//8
	//7
	//6
	//5
	//4
	//3
	//2
	//1
	//0
}

func TestDeferClosureLoopV3(t *testing.T) {
	DeferClosureLoopV3()
	//9
	//9
	//9
	//9
	//9
	//9
	//9
	//9
	//9
	//9
}
