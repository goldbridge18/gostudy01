package main

import "fmt"

type Ainterface interface {
	Say()
}

type Binterface interface {
	Hello()
}

type Students struct {
	Name string
}

func (students Students) Say() {
	fmt.Println("Student 结构体实现了Say() ")
}

type integer int

func (i integer) Say() {
	fmt.Println("integer say...")
}

//一个自定义类型实现多个接口
type Mutl struct {
	Id int
}

func (m Mutl) Say() {
	fmt.Println("Ainterface.....Say(一个自定义实现多个接口)")
}

func (m Mutl) Hello() {
	fmt.Println("Binterface....Hello(一个自定义实现多个接口)")
}

func demoInterfaceDetails01() {
	var students Students
	//接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量（实例)
	//如果 没有students ，直接var a Ainterface则会报错
	var a Ainterface = students
	a.Say()

	//只要是自定义数据类型，就可以实现接口，不仅是结构体类型。
	var i integer
	var a1 Ainterface = i
	a1.Say()

	//一个自定义实现多个接口

	var m Mutl
	var a2 Ainterface = m
	var b2 Binterface = m

	a2.Say()
	b2.Hello()

}
