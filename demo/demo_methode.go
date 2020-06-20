package main

import "fmt"

//定义一个结构体
type Personff struct {
	Name string
	Age  int
}

//定义一个方法

func (personff Personff) testff() {

	fmt.Println("I'm a fangfa")
}

func (personff Personff) add() {
	res := 0
	for i := 0; i < 1000; i++ {
		res = +i
	}
	fmt.Println(res)
}

//传参
func (personff Personff) add01(n int) {

	res := 0

	for i := 1; i < n; i++ {
		res = +i
	}

	fmt.Println(res)

}
func demoFangfa() {
	//初始化一个方法
	var personff Personff
	personff.testff()
	personff.add()

	n := 100

	personff.add01(n)
}
