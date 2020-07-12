package main

import (
	"fmt"
	"reflect"
)

//通过反射,修改 值

func reflect03(b interface{}) {

	//获取reflect.Vlaue
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal= ", rVal)

	//使用setxxxx 修改 ,因为传入的值时指针不能直接使用setxxx,需要Elem
	//rVal.SetInt(20) //报错
	rVal.Elem().SetInt(20)

	/*
		//如何理解rVal.Elem()

		num := 10
		ptr *int = &num
		num1 := *ptr  //类似于 rVal.Elem()
	*/
}

func demoReflect03() {
	var num int = 10
	reflect03(&num)
	fmt.Println(num)
}
