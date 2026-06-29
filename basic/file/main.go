package file

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

func HandleOpenFile() {
	//wd, _ := os.Getwd()
	//fmt.Println("当前工作目录:", wd)

	//file: file对象、file指针、file文件句柄
	file, err := os.Open("/Users/erick/Code/Golang/GoReviews/file/demo")
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
}

func HandleReadFileLine() {
	file, err := os.Open("/Users/erick/Code/Golang/GoReviews/file/hello")
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
}

func HandleReadFileAll() {
	filePath := "/Users/erick/Code/Golang/GoReviews/file/hello"
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
}

func HandleWriteFile() {
	filePath := "/Users/erick/Code/Golang/GoReviews/file/c.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	//及时关闭file
	defer file.Close()

	//\r\n表示换行，有些编辑器\n换行无效
	str := "hello,test6\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//因为writer是带缓存的，因此在调用writeString方法时，内容先写入内存。
	//所以需要调用Flush方法，将缓冲数据写入文件，否则文件中无数据。
	writer.Flush()
}

func JudgeFileExists() {
	filePath := "/Users/erick/Code/Golang/GoReviews/file/c.txt"
	_, err := PathExists(filePath)
	if err != nil {
		fmt.Println("file check err: ", err)
	}
	fmt.Println("file exists: ", filePath)
	//file exists:  /Users/erick/Code/Golang/GoReviews/file/c.txt
}

func ReWriteFile(dst, src string) {
	data, err := os.ReadFile(src)
	if err != nil {
		fmt.Println("read file a err: ", err)
		return
	}
	err = os.WriteFile(dst, data, 0666)
	if err != nil {
		fmt.Println("write file b err: ", err)
		return
	}
	fmt.Println("a ==> b ok!")
}

func HandleCopyFile(dst, src string) {
	_, err := CopyFile(dst, src)
	if err != nil {
		fmt.Println("copy file err: ", err)
		return
	}
	fmt.Println("copyFile ok!")
}
