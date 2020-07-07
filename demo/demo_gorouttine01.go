package main

import "fmt"

var (
	myMap = make(map[int]int, 10)
)

func test706(n int) {
	res := 1
	for i := 1; i < n; i++ {
		res *= i
	}

	//把res放入myMap
	myMap[n] = res
	//fmt.Println(myMap)
}
func demoGoroutine01() {

	for i := 0; i < 100; i++ {
		go test706(i) //此时可能无法运行出结果，因为主线程结束，协程还没有来的及执行，就结束了
		//如何解决这种问题有两种办法：1.加全局锁 2.channel。例子请看demo_goroutine02.go demo_goroutine03.go
		//test706(i)
	}

	for k, v := range myMap {
		fmt.Println("k = ", k, " ; v = ", v)
	}

}
