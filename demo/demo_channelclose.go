package main

import "fmt"

func demoChannelClose() {
	//	声明一个channel
	intChan711 := make(chan int, 3)

	//往intChan711写入数据
	intChan711 <- 100
	intChan711 <- 101
	//关闭inChan,
	close(intChan711)
	//关闭channel之后不能继续写入，
	//intChan711<- 102//报错
	//但是可以读取数据
	fmt.Println(<-intChan711)
	fmt.Println(<-intChan711)

	//遍历channel

	intChan71101 := make(chan int, 100)

	for i := 0; i < 100; i++ {
		intChan71101 <- i
	}
	/*
	   //遍历不能使用普通的for循环
	   //通过len或者cap 不能真正得到channel的大小;
	   for i :=0; i < len(intChan71101); i++{
	   	fmt.Println(<-intChan71101)
	   }
	*/

	//遍历时，没有关闭channel会报deadlock。因此需要使用close关闭管道
	close(intChan71101)
	//遍历 管道
	for v := range intChan71101 {
		fmt.Println(v)
	}
}
