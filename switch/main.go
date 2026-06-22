package main

import "fmt"

func main() {
	//var key byte
	//fmt.Println("请输入一个0-9的数字:")
	//_, error := fmt.Scanf("%d", &key)
	//if error != nil {
	//	return
	//}
	//
	//switch key {
	//case 0:
	//	fmt.Println("data: 0")
	//case 1:
	//	fmt.Println("data: 1")
	//case 2:
	//	fmt.Println("data: 2")
	//case 3:
	//	fmt.Println("data: 3")
	//case 4:
	//	fmt.Println("data: 4")
	//case 5, 6, 7, 8, 9:
	//	fmt.Println("data: >5")
	//default:
	//	fmt.Println("超出范围")
	//}

	////score := 69
	//switch score := 69; {
	//case score < 60:
	//	fmt.Println("不合格")
	//case score >= 60 && score <= 80:
	//	fmt.Println("良好")
	//case score > 80:
	//	fmt.Println("优秀")
	//default:
	//	fmt.Println("格式错误")
	//}

	//var num int = 10
	//switch num {
	//case 10:
	//	fmt.Println("10")
	//	fallthrough // 默认穿透到下一层
	//case 20:
	//	fmt.Println("20")
	//case 30:
	//	fmt.Println("30")
	//default:
	//	fmt.Println("格式错误")
	//}
	////10
	////20

	var x interface{}
	var y = 10
	x = y
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case float64, float32:
		fmt.Println("float64或float32")
	default:
		fmt.Println("格式错误")
	}
	//int
}
