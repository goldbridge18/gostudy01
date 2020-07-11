package main

import (
	"fmt"
	"time"
)

func demoChannel04() {

	//1.定义一个管道10 个数据int类型
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	//2.定义一个管道5个数据string类型
	strChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		strChan <- "hello" + fmt.Sprintf("%d", i)
	}

	//如果遍历管道时，没有关闭channel会阻塞导致deadlock

	//解决办法之一：select

	for {
		select {
		case v := <-intChan: //注意：这里，如果intChan一直没有关闭，也不会一直阻塞而deadlock
			fmt.Println("v = ", v)
			time.Sleep(time.Second)
		case v := <-strChan:
			fmt.Println("v = ", v)
			time.Sleep(time.Second)
		default:
			fmt.Println("not fund data ")
			return
		}
	}

}
