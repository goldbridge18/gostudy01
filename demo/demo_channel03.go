package main

import "fmt"

func demoChannel03() {

	//channel的只读、只写定义方式

	//1. 默认情况，管道是双向的可读可写
	var chan1 chan int
	fmt.Println(chan1)

	//2.声明只写的channel
	var chan2 chan<- int
	chan2 = make(chan int, 2)
	chan2 <- 10

	//num := <-chan2  //读取chan2的值，直接报错
	fmt.Println(chan2)

	//3.声明只读
	var chan3 <-chan int
	num1 := <-chan3
	//chan3 <- 11//x写入报错
	fmt.Println(num1)

}
