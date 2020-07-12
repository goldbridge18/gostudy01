package main

import (
	"fmt"
	"reflect"
)

//基础类型的转换
func reflectTest01(b interface{}) {

	//通过反射获取的传入的变量的type,kind,值
	//1. 先获取到reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType = ", rType)

	//2.获取reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal = ", rVal) //此时的值不是int类型了

	//3.获取 变量对应的Kind
	fmt.Println("kind = ", rVal.Kind())
	fmt.Println("kind = ", rType.Kind())

	var num2 int64 = 10
	//num3 := num2 + rVal  ///直接相加报错: invalid operation: num2 + rVal (mismatched types int and reflect.Value)
	//需要转化类型int才能计算
	num3 := num2 + rVal.Int()
	fmt.Println("num3 = ", num3)

	//4.将rVal转换interface{}
	iv := rVal.Interface()
	//将interface{} 通过断言转化成需要的类型

	num4 := iv.(int)
	fmt.Println("num4 = ", num4)

}

func demoReflect01() {
	var num int = 100
	reflectTest01(num)
}
