package testing

import (
	"testing"
)

//===============================(1)基本使用===============================
//文件名：以_test.go结尾， 如cal_test.go, a_test.go。
//测试用例函数：以Test开头，一般为Test+被测试函数名，如TestAddUpper
//形式：TestAddUpper(t *testing.T)

func TestAddUpper(t *testing.T) {
	data := AddUpper(10)
	//fmt.Println(data) // 45

	if data != 45 {
		//输出错误日志
		t.Fatal("Add Upper fail")
	}

	//t.Logf()输出相应日志
	t.Logf("Add Upper success")
	//=== RUN   TestAddUpper
	//main_test.go:38: Add Upper success
	//--- PASS: TestAddUpper (0.00s)
	//PASS

	//if data != 55 {
	//	t.Fatal("Add Upper fail")
	//	//=== RUN   TestAddUpper
	//	//main_test.go:45: Add Upper fail
	//	//--- FAIL: TestAddUpper (0.00s)
	//	//
	//	//FAIL
	//}
}

//===============================(2)测试单个文件===============================
//命令行：go test -v 单元测试文件 被测试员文件
//如：go test -v main_test.go main.go

//erick@erickdeMacBook-Pro testing % go test -v main_test.go main.go
//=== RUN   TestAddUpper
//main_test.go:31: Add Upper success
//--- PASS: TestAddUpper (0.00s)
//PASS
//ok      command-line-arguments  0.475s

//===============================(3)测试单个方法===============================
//命令行：go test -v -test.run 测试用例函数
//如：go test -v -test.run TestAddUpper

//erick@erickdeMacBook-Pro testing % go test -v -test.run TestAddUpper
//=== RUN   TestAddUpper
//main_test.go:22: Add Upper success
//--- PASS: TestAddUpper (0.00s)
//PASS
//ok      GoReviews/testing       0.477s
