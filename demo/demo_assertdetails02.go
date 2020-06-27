package main

import (
	"fmt"
)

func TypeJudge(items ...interface{}) {

	for index, x := range items {

		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数是bool类型，值是%v \n", index, x)
		case float32:
			fmt.Printf("第%v个参数是float32类型，值是%v \n", index, x)
		case float64:
			fmt.Printf("第%v个参数是float64类型，值是%v \n", index, x)
		case string:
			fmt.Printf("第%v个参数是string类型，值是%v \n", index, x)
		case Student27:
			fmt.Printf("第%v个参数是Student27类型，值是%v \n", index, x)
		case *Student27:
			fmt.Printf("第%v个参数是*Student27类型，值是%v \n", index, x)
		case int, int32, int64:
			fmt.Printf("第%v个参数是int系列类型，值是%v \n", index, x)
		default:
			fmt.Printf("第%v个参数是未知类型，值是%v \n", index, x)

		}

	}

}

type Student27 struct {
	Name string
}

func demoAssertdetail02() {

	n1 := "string"
	var n2 int = 100
	n3 := 100.01

	stu := Student27{"Lilei"} //结构体,自定义类型
	stu1 := &Student27{"hanmei"}

	//如何识别 结构体、指针

	TypeJudge(n1, n2, n3, stu, stu1)

}
