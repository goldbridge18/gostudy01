package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func test705() {
	for i := 0; i < 10; i++ {
		fmt.Println("test705()  goroutine........" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func test7051() {
	for i := 0; i < 10; i++ {
		fmt.Println("test7051() hello golang......" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

}

func demoGoroutine() {

	fmt.Println(runtime.NumCPU()) //本地有cpu个数
	//runtime.GOMAXPROCS(4)//设置cpu的最大个数，1.8版本之后默认使用多核
	go test705() //开启协程

	test7051()
	/*
		for i := 0; i < 10; i++ {
			fmt.Println("test7051() hello golang......" + strconv.Itoa(i))
			time.Sleep(time.Second)
		}
	*/
}
