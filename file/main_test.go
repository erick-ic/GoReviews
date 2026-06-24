package file

import (
	"testing"
)

// 1.基本操作：
// =============================(1.1)打开文件：=============================
func TestHandleOpenFile(t *testing.T) {
	HandleOpenFile()
}

// =============================(1.2)读取文件：=============================
// 方式1: 每次读取一行
func TestHandleReadFileLine(t *testing.T) {
	HandleReadFileLine()
}

// 方式2: 一次性全部读取
func TestHandleReadFileAll(t *testing.T) {
	HandleReadFileAll()
}

// =============================(1.3)写入文件：=============================
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
func TestHandleWriteFile(t *testing.T) {
	HandleWriteFile()
}

// =============================(1.4)判断文件、文件夹是否存在:=============================
func TestJudgeFileExists(t *testing.T) {
	JudgeFileExists()
}

// =============================实践1:=============================
// 覆盖操作:
// 将a.txt内容导入到b.txt

func TestReWriteFile(t *testing.T) {
	aFilePath := "/Users/erick/Code/Golang/GoReviews/file/a.txt"
	bFilePath := "/Users/erick/Code/Golang/GoReviews/file/b.txt"
	ReWriteFile(bFilePath, aFilePath)
}

// =============================实践2:=============================
// 文件拷贝：
// 将a.txt拷贝到assets文件夹
func TestHandleCopyFile(t *testing.T) {
	aFilePath := "/Users/erick/Code/Golang/GoReviews/file/a.txt"
	bFilePath := "/Users/erick/Code/Golang/GoReviews/file/assets/a.txt"
	HandleCopyFile(bFilePath, aFilePath)
}
