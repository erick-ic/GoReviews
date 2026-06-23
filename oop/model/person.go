package model

type Person struct {
	Name string
	age  int
}

// NewPerson 工厂模式的函数，类似于构造函数
func NewPerson(name string) *Person {
	return &Person{
		Name: name,
	}
}

func (p *Person) SetAge(age int) {
	p.age = age
}

func (p *Person) GetAge() int {
	return p.age
}
