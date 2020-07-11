package main

import "fmt"

func demoChannel() {

	//声明
	var intChan chan int
	//初始化
	intChan = make(chan int, 10) //是个地址，引用类型

	fmt.Println("intChan值：", intChan)
	fmt.Println("intChan地址值：", &intChan)

	//向管道写入数据

	intChan <- 10
	num := 211

	intChan <- num

	//查看管道长度和容量cap,数据不能超过他的容量大小
	fmt.Println("channel len = ", len(intChan), "cap is ", cap(intChan))

	//读取数据

	var num2 int
	num2 = <-intChan
	fmt.Println("num2", num2)

	fmt.Println("channel len = ", len(intChan), "cap is ", cap(intChan)) //读取一个值，长度就会减一。

	var num3 int
	num3 = <-intChan
	fmt.Println("num3", num3)
	//管道长度为2，再继续会报错
	//如果不使用协程的情况下，一个channel没有数据了，继续取，会报错：deadlock
	var num4 int
	num4 = <-intChan
	fmt.Println("num4", num4)

}
