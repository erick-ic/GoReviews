package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		//文件存在
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CopyFile(dstFileName, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	defer srcFile.Close()

	//获取reader
	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	defer dstFile.Close()

	//获取writer
	writer := bufio.NewWriter(dstFile)

	return io.Copy(writer, reader)
}

func main() {
	//1.基本操作：
	//=============================(1.1)打开文件：=============================
	//file: file对象、file指针、file文件句柄
	file, err := os.Open("file/demo")
	if err != nil {
		fmt.Println("open err: ", err)
	}
	fmt.Println("file: ", file)
	//file:  &{0x4a948f64c180}
	//未找到时报错：file:  <nil>

	err = file.Close()
	if err != nil {
		fmt.Println("close err: ", err)
	}

	//=============================(1.2)读取文件：=============================
	//方式1: 每次读取一行
	file, err = os.Open("file/hello")
	if err != nil {
		fmt.Println("open err: ", err)
	}
	//函数退出时，关闭file
	defer file.Close()

	//创建一个*Reader，带有缓冲，缓冲区为4096
	reader := bufio.NewReader(file)
	//line, err := reader.ReadString('\n')
	//fmt.Println(line)
	//hello--01
	for {
		//读到换行就结束
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			//io.EOF代表文件末尾
			break
		}
		fmt.Print(str)
		//hello--01!
		//hello--02!
		//hello--03!
		//hello--04!
		//hello--05!
		//hello--06!
	}
	fmt.Println()

	//方式2: 一次性全部读取
	filePath := "file/hello"
	content, err := os.ReadFile(filePath) //返回[]byte
	if err != nil {
		fmt.Println("read file err: ", err)
	}
	fmt.Println(string(content))
	//hello--01!
	//hello--02!
	//hello--03!
	//hello--04!
	//hello--05!
	//hello--06!

	//=============================(1.3)写入文件：=============================
	/*
		const (
			// O_RDONLY, O_WRONLY, O_RDWR 三者必须且只能指定其中一个
			O_RDONLY int = syscall.O_RDONLY // 只读方式打开文件
			O_WRONLY int = syscall.O_WRONLY // 只写方式打开文件
			O_RDWR   int = syscall.O_RDWR   // 读写方式打开文件

			// 以下常量可通过按位或（|）组合使用，以控制额外行为
			O_APPEND int = syscall.O_APPEND // 写入时追加到文件末尾
			O_CREATE int = syscall.O_CREAT  // 若文件不存在则创建新文件
			O_EXCL   int = syscall.O_EXCL   // 与 O_CREATE 一起使用，要求文件必须不存在（用于原子性创建）
			O_SYNC   int = syscall.O_SYNC   // 以同步 I/O 方式打开（每次写操作都等待物理写入完成）
			O_TRUNC  int = syscall.O_TRUNC  // 若文件存在且为普通可写文件，则打开时将其截断为空
		)
	*/
	filePath = "file/c.txt"
	file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	//及时关闭file
	defer file.Close()

	//\r\n表示换行，有些编辑器\n换行无效
	str := "hello,test\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//因为writer是带缓存的，因此在调用writeString方法时，内容先写入内存。
	//所以需要调用Flush方法，将缓冲数据写入文件，否则文件中无数据。
	writer.Flush()

	//=============================(1.4)判断文件、文件夹是否存在:=============================
	filePath = "file/c.txt"
	_, err = PathExists(filePath)
	if err != nil {
		fmt.Println("file check err: ", err)
	}

	//=============================实践1:=============================
	//将a.txt内容导入到b.txt
	aFilePath := "file/a.txt"
	bFilePath := "file/b.txt"

	data, err := os.ReadFile(aFilePath)
	if err != nil {
		fmt.Println("read file a err: ", err)
		return
	}
	err = os.WriteFile(bFilePath, data, 0666)
	if err != nil {
		fmt.Println("write file b err: ", err)
		return
	}
	fmt.Println("a ==> b ok!")

	//=============================实践2:=============================
	//文件拷贝：将a.txt拷贝到assets文件夹
	_, err = CopyFile("file/assets/a.txt", "file/a.txt")
	if err != nil {
		fmt.Println("copy file err: ", err)
		return
	}
	fmt.Println("copyFile ok!")
}
