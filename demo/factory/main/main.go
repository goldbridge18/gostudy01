package main

import (
	"demo/factory/model"
	"fmt"
)

func main() {

	var stu model.Student //model的Student 的首字母小写的话就会报错
	//var stu = model.Student{
	//	id : 1002,
	// Name: "Lily"
	// Score:90
	//}
	stu.Id = 1002
	stu.Name = "lily"
	stu.Score = 90

	fmt.Println(stu)

	//class结构体首字母是小写，可以通过工厂模式来解决

	var class = model.FactoryClass(100) //model里的FactoryClass返回的是个指针
	fmt.Println(*class)                 //使用(*号)取值
	fmt.Println(class.Level)

	//结构体字段 首字母为小写
	var class01 = model.Factoryclass01(101, "tom")
	fmt.Println(class01.GetLevel(), class01.GetName())
}
