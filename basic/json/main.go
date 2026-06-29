package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//JSON: Javascript Object Notation，一种轻量级的数据交换格式，易于机器解析和生成。
	//JSON 序列化与反序列化（数据交换通用格式）
	//通常情况下，程序在网络传输时，会先将数据(结构体、map等)序列化成JSON字符串，接收方得到后，再反序列化恢复成原数据类型。
	//过程：
	//    Go 结构体 / map / 切片
	//        ↓ 序列化（json.Marshal）
	//    JSON 字符串（纯文本，可网络传输）
	//        ↓ 网络传输 / 存储
	//    JSON 字符串（接收方得到）
	//        ↓ 反序列化（json.Unmarshal）
	//    目标语言的数据结构（如 Go 结构体、Python dict、Java 对象等）

	//关键点：
	//    - 序列化：将 Go 数据 -> JSON 字符串（用于发送或存储）。
	//    - 反序列化：将 JSON 字符串 -> Go 数据（用于接收或读取）。
	//    - JSON 是跨语言、跨平台的通用数据交换格式，几乎所有主流语言都支持。

	//1.序列化
	//将有key-value结构的数据类型（结构体、map、切片）序列化成JSON字符串的操作
	p := Person{
		Name:    "John Doe",
		Age:     18,
		Score:   100.0,
		Hobbies: []string{"fitness", "swimming", "coding"},
	}
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Marshal err: ", err)
	}
	fmt.Println(string(data))
	//{"Name":"John Doe","Age":18,"Score":100,"Hobbies":["fitness","swimming","coding"]}

	//2.反序列化
	//将JSON字符串反序列化成对应有key-value结构的数据类型的操作。
	jsonStr := "{\"Name\":\"John Doe\",\"Age\":18,\"Score\":100,\"Hobbies\":[\"fitness\",\"swimming\",\"coding\"]}\n"
	var p2 Person
	err = json.Unmarshal([]byte(jsonStr), &p2)
	if err != nil {
		fmt.Println("Unmarshal err: ", err)
	}
	fmt.Println("p2: ", p2)
	//p2:  {John Doe 18 100 [fitness swimming coding]}
}

type Person struct {
	Name    string
	Age     int
	Score   float64
	Hobbies []string
}
