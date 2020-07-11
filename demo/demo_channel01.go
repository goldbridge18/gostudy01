package main

import "fmt"

type Cat struct {
	Name string
	age  int
}

func demoChannel01() {

	//定义一个存放任意数据类型的channel，可以存放3个数据

	//声明一个接口类型的channel
	var allChan chan interface{} //空接口类型 的channel，接口可以存放任意数据类型
	allChan = make(chan interface{}, 3)

	//allChan := make(chan interface{}, 3) //声明一个接口类型的channel

	allChan <- 100
	allChan <- "cat007"

	/*
		var cat Cat //初始化Cat结构体
		cat.age = 18
		cat.Name = "008"
	*/
	/*
		var cat Cat = Cat{"flower cat", 10} //初始化Cat结构体
	*/
	cat := Cat{"flower cat", 10}
	allChan <- cat

	//channel是先进先出，如果要取出cat的名字，就需要先把前两个取出来，才能取出Cat结构体的数据
	<-allChan //取出数据，不赋值给任何对象
	<-allChan

	getCat := <-allChan //取出Cat结构体

	fmt.Printf("getCat is %T, getCat is %v\n", getCat, getCat)

	//获取cat的名字，如果使用getCat.Name直接报错
	catName := getCat.(Cat).Name //使用断言
	fmt.Println("cat's name is ", catName)

}
