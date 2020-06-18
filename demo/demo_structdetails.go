package main

import "fmt"

//如果结构体的字段类型是：指针、slice、map的零值都是nil，即还没有分配空间，指针需要new来分配空间，
//其他的则使用make分配空间之后才能使用
type Person struct {
	Name  string
	Age   string
	Score [5]float64 //数组
	ptr   *int
	slice []int64
	map1  map[string]string
}

func demoStructDetails() {

	var person Person
	person.Score[1] = 100

	//指针 new来分配空间

	//slice 使用make分配空间
	person.slice = make([]int64, 2)
	person.slice[0] = 555
	person.slice[1] = 666

	//map 使用make分配空间，动态增长大小
	person.map1 = make(map[string]string)
	person.map1["111"] = "222"
	person.map1["333"] = "444"

	fmt.Println(person)
}
