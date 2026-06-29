package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// A 空接口：任意类型都可实现空接口，无任何约束
type A interface{}

// Hero 结构体
type Hero struct {
	Name string
	Age  int
}

// HeroSlice 结构体切片类型
// [{Name, Age}, {Name, Age}, ...]
type HeroSlice []Hero

// Len 对结构体切片HeroSlice实现interface
func (hs HeroSlice) Len() int {
	return len(hs)
}

// Less 按照年龄从小到大排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
	////姓名升序
	//return hs[i].Name < hs[j].Name
}

// Swap 交换元素位置
func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

type Monkey struct {
	Name string
}

func (m *Monkey) Climbing() {
	fmt.Println(m.Name, "Climbing...")
}

// Bird 声明接口，实现其他方法
type Bird interface {
	Flying()
}

type LittleMonkey struct {
	Monkey //继承
}

func (l *LittleMonkey) Flying() {
	fmt.Println(l.Name, "Flying...")
}

func main() {
	//类型：引用类型，未初始化就使用会输出nil
	//含义：接口是一组行为的抽象。使用接口，实现面向接口编程。
	//优点：定义了对象行为规范，只定义不实现，体现程序设计的多态、高内聚低耦合思想。
	//步骤：一个变量含有接口类型中的全部方法，该变量就实现了这个接口。

	//(1)基本语法
	//(1.1)实现空接口
	var a A

	str := "hello interface"
	//string类型实现了A接口
	a = str
	fmt.Printf("%v, %T\n", a, a)
	//hello interface, string

	flag := true
	//bool类心实现了A接口
	a = flag
	fmt.Printf("%v, %T\n", a, a)
	//true, bool

	//map实现空接口
	user := make(map[string]interface{})
	user["name"] = "jack"
	user["age"] = 18
	fmt.Printf("%v, %T\n", user, user)
	//map[age:18 name:jack], map[string]interface {}

	//slice实现空接口
	typeSlice := make([]interface{}, 6)
	typeSlice[0] = "string"
	typeSlice[1] = true
	typeSlice[2] = 1
	typeSlice[3] = 21.21
	typeSlice[4] = [3]int{1, 2, 3}
	fmt.Printf("%v, %T\n", typeSlice, typeSlice)
	//[string true 1 21.21 [1 2 3] <nil>], []interface {}
	//这里切片的元素类型是 interface{}，即空接口。
	//空接口在 Go 中不规定任何方法，因此所有类型都实现了空接口（包括基本类型、结构体、数组等）。
	//实际上每个元素存储的都是一个 interface{} 类型的值，但这个 interface{} 内部包裹了具体类型的值。
	//所以不是“元素类型混用”，而是所有元素类型统一为 interface{}，只不过每个元素持有的动态类型和值不同。

	//(1.2)接口中不能出现变量
	/*
		type AInterface interface {
			//syntax error: unexpected name string in interface type
			Name string
			test()
		}

		func test() {
			fmt.Println("test...")
		}
	*/

	//最佳实践：
	////切片排序
	//intSlice := []int{0, -1, 9, 6}
	////方式1: 冒泡排序
	////方式2: 工具方法
	//sort.Ints(intSlice)
	//fmt.Printf("%v, %T\n", intSlice, intSlice)
	////[-1 0 6 9], []int

	//结构体切片排序
	//方式1: 冒泡排序
	//方式2: sort方法
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			//rand.Intn(100)随即返回[0, 100)的数值
			Name: fmt.Sprintf("Hero-%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}
	for _, hero := range heroes {
		fmt.Println("hero: ", hero)
		//hero:  {Hero-32 30}
		//hero:  {Hero-32 17}
		//hero:  {Hero-8 71}
		//hero:  {Hero-28 30}
		//hero:  {Hero-12 68}
		//hero:  {Hero-48 51}
		//hero:  {Hero-86 30}
		//hero:  {Hero-76 78}
		//hero:  {Hero-63 86}
		//hero:  {Hero-7 22}
	}
	fmt.Println("----------sort begin----------")
	//因为heroes实现了Len、Less、Swap，也就是实现了sort.interface
	sort.Sort(heroes)
	for _, hero := range heroes {
		fmt.Println("hero: ", hero)
		//年龄升序
		//hero:  {Hero-90 10}
		//hero:  {Hero-56 27}
		//hero:  {Hero-49 40}
		//hero:  {Hero-69 43}
		//hero:  {Hero-15 43}
		//hero:  {Hero-20 56}
		//hero:  {Hero-46 60}
		//hero:  {Hero-20 89}
		//hero:  {Hero-56 96}
		//hero:  {Hero-89 98}
	}
	//(2)接口与继承
	//接口是对继承的补充，在不破坏继承关系的情况下，对功能进行扩展
	//继承的价值：解决代码的复用性、可维护性
	//接口的价值：设计，设计各种规范，让自定义类型去实现这些方法
	//接口更加灵活，继承 ==> is，接口 ==> like
	lm := LittleMonkey{
		Monkey{
			Name: "Monkey01",
		},
	}
	lm.Climbing()
	//Monkey01 Climbing...
	//通过接口方式实现其他方法
	lm.Flying()
	//Monkey01 Flying...
}
