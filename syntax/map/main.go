package main

import (
	"fmt"
	"strings"
)

func main() {
	//类型：引用类型
	//定义方式：var map变量名[keytype]valuetype
	//key通常为int，string。slice、map和function不能作为key，因为无法用==判断
	//key是无序的

	////声明不会分配内存，make初始化之后才能赋值、使用
	//var a map[string]string
	//a["name"] = "Jack"
	//fmt.Println(a)
	////panic: assignment to entry in nil map

	//(1)基本语法：
	//方式1:
	var users map[string]string
	users = make(map[string]string, 10)
	users["name"] = "Jack"
	fmt.Printf("%v, %T\n", users, users)
	//map[name:Jack], map[string]string

	//方式2:
	users2 := make(map[string]int)
	users2["id"] = 1
	users2["age"] = 21
	fmt.Printf("%v, %T\n", users2, users2)
	//map[age:21 id:1], map[string]int

	//方式3:
	users3 := map[string]string{
		"name": "Jack",
		"age":  "21",
	}
	fmt.Printf("%v, %T\n", users3, users3)
	//map[age:21 name:Jack], map[string]string

	//(2)常用方法
	//（2.1）增：
	userMaps := map[string]string{
		"name":   "Jack",
		"age":    "21",
		"gender": "male",
	}
	fmt.Println(userMaps)
	//map[age:21 gender:male name:Jack]
	userMaps["nickname"] = "Tom"
	fmt.Println(userMaps)
	//map[age:21 gender:male name:Jack nickname:Tom]

	//(2.2) 删：
	delete(userMaps, "nickname")
	fmt.Println(userMaps)
	//map[age:21 gender:male name:Jack]

	////全部删除：
	//userMaps = make(map[string]string)
	//fmt.Println("userMaps:", userMaps)
	////userMaps: map[]

	//(2.3) 改：
	userMaps["age"] = "18"
	fmt.Println(userMaps)
	//map[age:18 gender:male name:Jack]

	//(2.4) 查：
	v1, ok1 := userMaps["nickname"]
	fmt.Println(v1, ok1)
	//false
	v2, ok2 := userMaps["age"]
	fmt.Println(v2, ok2)
	//18 true

	//(3)map循环
	mapLoop := map[string]string{
		"name":   "Jack",
		"age":    "21",
		"gender": "male",
	}
	for key, value := range mapLoop {
		fmt.Printf("k = %v, v = %v \n", key, value)
		//k = name, v = Jack
		//k = age, v = 21
		//k = gender, v = male
	}

	//(4)map与slice
	//(4.1)map切片 => 值为map的切片
	mapSlice := make([]map[string]string, 3)
	if mapSlice[0] == nil {
		mapSlice[0] = make(map[string]string)
		mapSlice[0]["name"] = "Jack"
	}
	if mapSlice[1] == nil {
		mapSlice[1] = make(map[string]string)
		mapSlice[1]["name"] = "Tom"
	}
	fmt.Printf("mapSlice: %v, type: %T \n", mapSlice, mapSlice)
	//mapSlice: [map[name:Jack] map[name:Tom] map[]], type: []map[string]string

	//(4.2)切片map => 值为切片的map
	sliceMap := make(map[string][]string)
	sliceMap["name"] = []string{"Jack", "Tom"}
	sliceMap["hobby"] = []string{"sing", "dance", "rap"}
	fmt.Printf("sliceMap: %v, type: %T \n", sliceMap, sliceMap)
	//sliceMap: map[hobby:[sing dance rap] name:[Jack Tom]], type: map[string][]string

	//案例1: 统计字符串中每个单词出现的次数。如："how do you do"中how = 1, do = 2, you = 1
	str := "how do you do"
	strSlice := strings.Split(str, " ")
	//fmt.Printf("strSlice: %v, type: %T \n", strSlice, strSlice)
	////strSlice: [how do you do], type: []string
	strMap := make(map[string]int)
	for _, v := range strSlice {
		strMap[v]++
	}
	fmt.Println("strMap: ", strMap)
	//strMap:  map[do:2 how:1 you:1]
}
