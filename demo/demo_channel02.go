package main

import "fmt"

/*
利用channel;对同一个channel边写边读

*/
//to channel write data
var count int = 5000

func writeData(intChan chan int) {
	for i := 0; i < count; i++ {
		intChan <- i
		fmt.Println("write data is  ", i)
	}
	close(intChan) // close channel
}

//read from channel data

func readData(intChan chan int, exitChan chan bool) {

	for {
		v, ok := <-intChan
		if !ok {
			break
		}

		fmt.Println("read data is ", v)

	}
	//read data alway is not data .task is completed
	exitChan <- true
	close(exitChan)
}

func demoChannel02() {

	// 声明channel
	intChan := make(chan int, count-100)
	exitChan := make(chan bool, 1)

	//write data
	go writeData(intChan)

	//go readData(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

}
