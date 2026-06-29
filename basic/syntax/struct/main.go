package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	Gender string
}

func (p Person) GetInfos() {
	fmt.Printf("Name: %#v, Age: %#v, Gender: %#v\n", p.Name, p.Age, p.Gender)
}

func (p *Person) SetInfos(name string, age int) {
	p.Name = name
	p.Age = age
}

type Author struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int
	slice  []int
	map1   map[string]string
}

type User struct {
	name    string
	age     int
	hobbies []string
	userMap map[string]string
	company Company
}

type Company struct {
	companyName string
	companyType string
}

type Monster struct {
	Name  string `json:"name"`
	Age   int
	Skill string
}

func main() {
	//类型：值类型

	//(1)基本语法：
	//方式1:
	var p1 Person
	p1.Name = "Jack"
	p1.Age = 21
	p1.Gender = "male"
	fmt.Printf("%v, %T\n", p1, p1)
	//{Jack 21 male}, main.Person
	fmt.Printf("%#v, %T\n", p1, p1)
	//main.Person{Name:"Jack", Age:21, Gender:"male"}, main.Person

	//方式2:
	var p2 *Person = new(Person)
	p2.Name = "Mary"
	p2.Age = 18
	p2.Gender = "female"
	fmt.Printf("%#v, %T\n", p2, p2)
	//&main.Person{Name:"Mary", Age:18, Gender:"female"}, *main.Person

	//方式3:
	var p3 *Person = &Person{}
	p3.Name = "Tom"
	p3.Age = 25
	p3.Gender = "male"
	fmt.Printf("%#v, %T\n", p3, p3)
	//&main.Person{Name:"Tom", Age:25, Gender:"male"}, *main.Person

	//方式4:
	var p4 Person = Person{
		Name:   "Jack",
		Age:    21,
		Gender: "female",
	}
	fmt.Printf("%#v, %T\n", p4, p4)
	//main.Person{Name:"Jack", Age:21, Gender:"female"}, main.Person

	//方式5:
	var p5 Person = Person{
		Name: "Jack",
		//部分参数可省略
	}
	fmt.Printf("%#v, %T\n", p5, p5)
	//main.Person{Name:"Jack", Age:0, Gender:""}, main.Person

	//方式6: 不推荐
	p6 := Person{
		"Jack",
		28,
		"male",
	}
	fmt.Printf("%#v, %T\n", p6, p6)
	//main.Person{Name:"Jack", Age:28, Gender:"male"}, main.Person

	//(2)常用方法：
	//(2.1)访问成员：
	//在go中支持对结构体指针直接使用，来访问结构体成员。
	//p7 := new(Person)
	p7 := &Person{}
	p7.Name = "Jack"
	//p7.Name = "Jack" ==> (*p).name = "Jack"
	fmt.Printf("%#v, %T\n", p7, p7)
	//&main.Person{Name:"Jack", Age:0, Gender:""}, *main.Person
	fmt.Printf("Name: %#v\n", p7.Name)
	//Name: "Jack"

	p8 := Person{}
	p8.Name = "Jack"
	fmt.Printf("Name=%#v\n", p8.Name)
	//Name="Jack"

	//(2.2)自定义方法：
	p9 := Person{
		Name: "Jack",
		Age:  21,
	}
	p9.GetInfos()
	//Name: "Jack", Age: 21, Gender: ""

	//修改struct值，接收者需要定义为指针类型
	p9.SetInfos("Tome", 18)
	fmt.Printf("%#v\n", p9)
	//main.Person{Name:"Tome", Age:18, Gender:""}

	//(2.3)属性为指针、slice、map时，由于零值都是nil，即未分配内存，使用前需要make初始化。
	a := Author{}
	fmt.Printf("a: %#v\n", a)
	//a: main.Author{Name:"", Age:0, Scores:[5]float64{0, 0, 0, 0, 0}, ptr:(*int)(nil), slice:[]int(nil), map1:map[string]string(nil)}

	a.Name = "Jack"
	a.Age = 21
	a.Scores = [5]float64{5, 6, 7, 8}

	a.ptr = new(int)

	////slice需要先make初始化，才能使用
	////index out of range [0] with length 0
	//a.slice[0] = 1
	a.slice = make([]int, 3)
	a.map1 = make(map[string]string)

	//(2.4)结构体嵌套：
	u := User{}
	u.name = "Jack"
	u.age = 21
	u.hobbies = []string{"sing", "dance", "rap"}
	u.userMap = make(map[string]string)
	u.userMap["address"] = "No.1"
	u.userMap["date"] = "2026-06-22"
	u.company.companyName = "apple"
	u.company.companyType = "IT"
	fmt.Println("hobbies: ", u.hobbies)
	//hobbies:  [sing dance rap]
	fmt.Println("userMap: ", u.userMap)
	//userMap:  map[address:No.1 date:2026-06-22]
	fmt.Println("company: ", u.company)
	//company:  {apple IT}

	//(2.5)标签：
	//使用场景：序列化、反序列化
	m := Monster{"牛魔王", 500, "无敌冲撞"}

	//变量序列化：
	//json.Marshal使用了反射
	jsonStr, err := json.Marshal(m)
	if err != nil {
		fmt.Println("err:", err)
	}
	//返回的是[]byte，需要string转换
	fmt.Println(string(jsonStr))
	//{"name":"牛魔王","Age":500,"Skill":"无敌冲撞"}
}
