package main

import "fmt"

type Person01 struct {
	name string
	age  int
}

func demoStructSay() {

	//直接声明
	var person Person
	person.Name = "hanmei"
	person.Age = "18"
	fmt.Println(person)

	//声明 {}
	//var person01 Person01 = Person01{"hanmei01", 19}
	person01 := Person01{"hanmei01", 19}
	fmt.Println(person01)

	//声名 &

	var person02 *Person01 = new(Person01) //person02是一个指针
	person02.name = "HanMei02"             //等价 (*person02).name = "HanMei02";底层把person02.name 的写法转化为(*person02)
	person02.age = 16

	fmt.Println(*person02) //*号为取值符
	fmt.Println(&person02) //&为取址符

	//声明 {}
	var person03 *Person01 = &Person01{"HanMei03", 17} //person03是一个指针
	fmt.Println(*person03)
}
