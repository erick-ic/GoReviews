package testing

func AddUpper(n int) int {
	result := 0
	for i := 0; i < n; i++ {
		result = result + i
	}
	return result
}

func main() {
	//传统方式：
	//1.需要在main函数中调用，需要修改main函数，可能需要重启项目。
	//2.不利于管理，测试多个函数、模块都在main函数中。

	//单元测试：
	//1.确保每个函数可运行，且运行正确。
	//2.确保写出的代码性能良好。
	//3.及时发现程序设计、实现的逻辑错误，便于定位。如性能测试，高并发等场景。
}
