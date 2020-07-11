package main

import (
	"fmt"
	"time"
)

var count711 int = 10

func sayHello() {
	for i := 0; i < count711; i++ {
		time.Sleep(time.Second)
		fmt.Println("ni hao wo guo")

	}
}

func test711() {

	//捕获异常 defer+recover

	defer func() { //匿名函数
		if err := recover(); err != nil {
			fmt.Println("test711 发生error", err)
		}

	}() //调用

	//定义一个map
	var myMap map[int]string
	myMap[0] = "golang" //使用map需要make，因此触发panic 报错
}

func demoChannel05() {

	go sayHello()
	go test711()

	for i := 0; i < 10; i++ {
		fmt.Println("..............", i)
	}
	time.Sleep(time.Second)
}
