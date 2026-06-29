package main

import (
	"fmt"
	"reflect"
)

func reflectInt(n interface{}) {
	//通过反射获取变量的type、kind、value等

	//获取reflect.Type
	nType := reflect.TypeOf(n)
	fmt.Printf("nType:%v, type: %T \n", nType, nType)
	//nType:int, type: *reflect.rtype

	//获取reflect.Value
	nValue := reflect.ValueOf(n)
	fmt.Printf("nValue:%v, type: %T \n", nValue, nValue)
	//nValue:100, type: reflect.Value

	//a := 1
	////fmt.Println(a + nValue)
	////invalid operation: a + nValue (mismatched types int and reflect.Value)
	//n2 := a + int(nValue.Int())
	//fmt.Printf("n2:%v, type: %T \n", n2, n2)
	////n2:101, type: int

	//将reflect.Value转换成interface{}
	iV := nValue.Interface()
	n3 := iV.(int)
	fmt.Printf("n3:%v, type: %T \n", n3, n3)

}

type Student struct {
	Name string
	Age  int
}

func reflectStruct(n interface{}) {
	nType := reflect.TypeOf(n)
	fmt.Printf("nType:%v, type: %T \n", nType, nType)
	//nType:main.Student, type: *reflect.rtype

	nValue := reflect.ValueOf(n)
	fmt.Printf("nValue:%v, type: %T \n", nValue, nValue)
	//nValue:{Jack 20}, type: reflect.Value

	iV := nValue.Interface()
	fmt.Printf("iV:%v, type: %T \n", iV, iV)
	//iV:{Jack 20}, type: main.Student
	//打印的是运行时的类型，和编译时的类型不同。
	//fmt.Println(iV.Name)
	//iV.Name undefined (type any has no field or method Name)
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu:%v, type: %T \n", stu, stu)
		fmt.Println(stu.Name)
		//Jack
	}
}

func main() {
	//reflect包实现了运行时的反射
	//变量 <==> interface{} <==> reflect.Value, 可以相互转换

	//(1)基本使用：
	//获取变量类型：reflect.TypeOf(变量名)
	//获取变量值：reflect.ValueOf(变量名)

	//var num int = 100
	//reflectInt(num)

	//stu := Student{
	//	Name: "Jack",
	//	Age:  20,
	//}
	//reflectStruct(stu)

	//(2)最佳实践：
	//使用反射遍历结构体的字段，调用结构体的方法，并获取结构体标签的值。
	m := Monster{
		Name:  "jack",
		Age:   20,
		Score: 99.9,
		Sex:   "male",
	}
	TestStruct(m)
}

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

func (m Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(m)
	fmt.Println("---end---")
}

func (m Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (m Monster) Set(name string, age int, score float32, sex string) {
	m.Name = name
	m.Age = age
	m.Score = score
	m.Sex = sex
}

func TestStruct(m interface{}) {
	//获取refelect.Type
	mType := reflect.TypeOf(m)
	//获取reflect.Value
	mValue := reflect.ValueOf(m)
	//获取类别
	mKind := mType.Kind()
	//判断是否为struct
	if mKind != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	//获取struct字段
	numField := mValue.NumField()
	fmt.Println("numField:", numField)
	//numField: 4

	for i := 0; i < numField; i++ {
		//fmt.Printf("Field %d:%v\n", i, mValue.Field(i))
		//Field 0:jack
		//Field 1:20
		//Field 2:99.9
		//Field 3:male
		//获取struct标签，通过reflect.Type获取tag标签的值
		tagVal := mType.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d:%v\n", i, tagVal)
			//Field 0:name
			//Field 1:monster_age
		}
	}
	//获取struct有多少方法
	numOfMethod := mValue.NumMethod()
	fmt.Println("numOfMethod:", numOfMethod)
	//numOfMethod: 3

	//获取并调用struct第2个方法：
	//顺序默认按照函数名的ASCII码：GetSum()、Print()、Set()
	mValue.Method(1).Call(nil)
	//Print():
	//---start---
	//{jack 20 99.9 male}
	//---end---

	//获取第1个方法：
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(20))
	//传入的参数：[]reflect.Value
	//返回的结果：[]reflect.Value
	res := mValue.Method(0).Call(params)
	fmt.Println("res:", res[0].Int())
	//res: 30
}
