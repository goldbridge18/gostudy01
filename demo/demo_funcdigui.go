package main

import "fmt"

//函数分配在栈
func test(n int){

	if n>2 {
		n --
		test(n)
	}

	fmt.Println("n is ",n)
}

func run(){

	test(4)
}