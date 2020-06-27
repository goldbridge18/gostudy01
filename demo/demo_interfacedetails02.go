package main

import "fmt"

//空接口interface{}没有任何方法，所有类型都实现了空接口。即可以把任何一个变量赋值给空接口

type T interface{

}
type Student04 struct{
	Name string
}

func (stu Student04) Prnull(){
	fmt.Println("空接口。。。。。")
}

func demoInterface02(){

	var t T
	fmt.Println(t) //接口是指针为引用类型，空接口返回nil

	var stu Student04
	var t1 T = stu
	fmt.Println(t1)
	//fmt.Println(stu.Prnull())//报错

	var t2  interface{} = stu  //直接定义声明并定义一个空接口
	fmt.Println(t2)

}