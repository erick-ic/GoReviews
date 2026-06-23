package main

import (
	"GoReviews/oop/model"
	"fmt"
)

type Student struct {
	Name  string
	age   int
	Score float64
}

func (s *Student) GetStudent() {
	fmt.Printf("name: %v, age: %v, score: %v\n", s.Name, s.age, s.Score)
}

func (s *Student) SetScore(score float64) {
	s.Score = score
}

func (s *Student) sayHi() {
	fmt.Println("Student: hi~")
}

// 小学生
type Pupil struct {
	//嵌套Studentn匿名结构体
	Student

	//有名结构体
	//s Student
}

func (p *Pupil) test() {
	fmt.Println("pupil test...")
}

// 大学生
type Undergraduate struct {
	Student
}

func (u *Undergraduate) test() {
	fmt.Println("undergraduate test...")
}

func main() {
	//面向对象编程三大特性：继承、封装和多态

	//一、封装
	//定义：把抽象出的字段和对字段的操作封装组合，程序内其他包只有通过被授权的操作（方法），才能对字段进行操作。
	//优点：隐藏实现细节、对数据进行验证，确保安全合理。
	//步骤：
	/*
		1.将结构体、属性（字段）的首字母小写
		2.给结构体所在包提供一个工厂模式的函数，首字母大写。
		3.提供Set方法，用于属性赋值
		4.提供Get方法，用于获取属性的值
	*/
	p := model.NewPerson("Jack")
	fmt.Printf("p: %#v\n", *p)
	//p: model.Person{Name:"Jack", age:0}
	fmt.Printf("name: %v\n", p.Name)
	//name: "Jack"
	fmt.Printf("age: %v\n", p.GetAge())
	//age: 0

	p.SetAge(30)
	fmt.Printf("age: %v\n", p.GetAge())
	//age: 30

	//二、继承
	//定义：当多个结构体存在相同属性和方法时，可从这些结构体中抽象出结构体，并在该结构体中定义相同的属性和方法。
	//在Go中，一个结构体嵌套了另一个结构体，该结构体可访问嵌套结构的属性和方法，即实现了继承。
	//结构体A嵌套了结构体B，A可访问B中的属性和方法，则A继承了B。
	//优点：减少代码冗余，提高复用性、扩展性和可维护性。
	//步骤：嵌套抽象出的结构体
	pu := &Pupil{}
	pu.Student.Name = "Jack"
	//pu.Student.age = 18，其中的Student可省略
	pu.age = 18
	pu.test()
	//pupil test...
	pu.SetScore(100)
	pu.GetStudent()
	//name: Jack, age: 18, score: 100
	pu.sayHi()
	//Student: hi~

	u := &Undergraduate{}
	u.Name = "Mary"
	u.age = 21
	u.test()
	//undergraduate test...
	u.SetScore(100)
	u.GetStudent()
	//name: Mary, age: 21, score: 100

	//三、多态
	//定义：同一个接口，多种不同的实现。“一个命令，多种响应”
	//多态通过接口实现，可以按照统一的接口，调用不同的实现。
	//优势：解耦、可扩展、代码复用、维护方便。
	//步骤：
	/*
		统一接口：Usb接口定义了Start、Stop方法
		多种实现：Phone和Camera分别实现了这两个方法（输出不同内容）
		统一处理：Computer.Working(usb Usb)方法的参数是Usb接口类型，不关心传入的是什么，只关心传入的能Start、Stop
		运行时动态绑定：传入phone，usb.Start()执行Phone的Start，在运行时根据传入的实际类型决定，这就是多态的核心。
	*/

	computer := &Computer{}
	phone := &Phone{}
	camera := &Camera{}

	computer.Working(phone)
	//Phone Start
	//Phone Stop
	computer.Working(camera)
	//Camera Start
	//Camera Stop
}

// Usb 接口声明了Start、Stop方法，意味着任何类型只要实现了这两个方法，就自动成为Usb类型。
type Usb interface {
	Start()
	Stop()
}

/*
如果一个类型实现接口时用的是值接收者（如 (p Phone)），那么该类型的值和指针都实现了该接口。
但如果用的是指针接收者（如 (p *Phone)），那么只有指针类型（*Phone）实现了该接口，值类型（Phone）不满足接口。

// 假设把 Phone 的方法改成指针接收者
func (p *Phone) Start() { ... }
func (p *Phone) Stop() { ... }

// 调用时
computer.Working(phone)   // 编译报错！Phone 没有实现 Usb
computer.Working(&phone)  // 编译通过，*Phone 实现了 Usb
*/

// Phone 实现Usb接口：Start、Stop ==> Phone是Usb
type Phone struct{}

func (p *Phone) Start() {
	fmt.Println("Phone Start")
}

func (p *Phone) Stop() {
	fmt.Println("Phone Stop")
}

// Camera 实现Usb接口：Start、Stop ==> Camera是Usb
type Camera struct{}

func (c *Camera) Start() {
	fmt.Println("Camera Start")
}
func (c *Camera) Stop() {
	fmt.Println("Camera Stop")
}

type Computer struct{}

// Working 只要实现了Usb接口，就实现了Usb接口的所有方法
// Working 方法接收的是“能 Start 和 Stop 的东西”，至于它具体是什么，由调用者决定，这就是接口赋予的灵活性和多态性。
// 参数 usb 是一个接口类型的变量， 它不关心传入的具体是什么，只关心传入的东西有没有 Start 和 Stop 方法。
func (com *Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}
