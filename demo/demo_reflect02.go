package main

import (
	"fmt"
	"reflect"
)

//对结构体的反射
func reflectTest02(b interface{}) {

	//通过反射获取的传入的变量的type,kind,值
	//1. 先获取到reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType = ", rType)

	//2.获取reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal = ", rVal) //此时的值不是int类型了

	fmt.Println("rType  kind = ", rType.Kind())
	fmt.Println("rVal kind = ", rVal.Kind())
	//3.将rVal转换interface{}
	iv := rVal.Interface()
	fmt.Printf("iv = %v, iv = %T  iv = %v \n", iv, iv, iv) //取值报错
	//fmt.Printf("i  iv = %v \n", iv.Name) //取值报错,无法知道类型,需要断言来解决问题

	//将interface{} 通过断言转化成需要的类型
	//使用类型断言检测类型. 可以参照:demo_assertdetails02.go

	stu, ok := iv.(Student712)
	if ok {
		fmt.Println("stu.Name = ", stu.Name)
	}

}

type Student712 struct {
	Name string
	Age  int
}

func demoReflect02() {
	stu := Student712{
		Name: "LiLei",
		Age:  18,
	}

	reflectTest02(stu)
}
