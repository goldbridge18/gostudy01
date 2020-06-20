package main

import "fmt"

type School struct {
	Name  string
	Level int
}

//给*School 实现方法String()
func (sch *School) String() string {

	str := fmt.Sprintf("Name=[%v]  Level=[%v]", sch.Name, sch.Level)
	return str
}

func demoMethodedetails() {

	var school School
	school.Name = "zheng qing"
	school.Level = 2

	fmt.Println(school)
	//如果实现了 *School 类型的string方法，就会自动调用string
	fmt.Println(&school) //传入一个地址
	fmt.Println(school.String())

}
