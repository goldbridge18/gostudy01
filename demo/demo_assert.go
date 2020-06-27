package main

import "fmt"

type Point struct {
	x int
	y int
}

func demoAssert() {
	var a interface{} //定义一个空接口
	var point Point = Point{10, 20}

	a = point
	fmt.Println(a)

	var b Point
	//如果将a赋值给一个Point变量？
	//如果将 b = a 则报错，提示需要断言，如下行
	//b = a  //cannot use a (type interface {}) as type Point in assignment: need type assertion
	b = a.(Point) //类型断言
	fmt.Println(b)

	//类型断言案例
	var x interface{}
	var b2 float32 = 1.1
	x = b2 //x为空接口，可以接收任意类型
	fmt.Println(x)
	//x为接口不知道具体类型，因此需要使用类型断言转化类型
	y := x

	fmt.Printf("y 的类型是%T，值是%v", y, y)

	//类型断言（带检测）
	z, ok := x.(float64)
	if ok { //if语句的其他写法 ：if z, ok := x.(float64) , ok{}
		fmt.Println("convert success")
		fmt.Printf("z的类型是%T， 值是%v \n", z, z)
	} else {
		fmt.Println("convert fail")
	}

	fmt.Println("continue ..")

}
