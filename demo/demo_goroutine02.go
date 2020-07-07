package main

import (
	"fmt"
	"sync"
	"time"
)

//解决：

//此时可能无法运行出结果，因为主线程结束，协程还没有来的及执行，就结束了
//如何解决这种问题有两种办法：1.加全局锁
var (
	myMap02 = make(map[int]int, 10)

	//声明一个全局互斥锁 lock
	//需要引入sync包:synchornized 同步
	//Mutex: 是互斥
	lock sync.Mutex
)

func test70602(n int) {
	res := 1
	for i := 1; i < n; i++ {
		res *= i
	}

	//把res放入myMap
	//枷锁
	lock.Lock()
	myMap02[n] = res
	//解锁
	lock.Unlock()
}
func demoGoroutine02() {

	//	lock.Lock()
	for i := 1; i < 100; i++ { //如果10的值比较大 好多v的值可能为0说明myMap越界了
		test70602(i)
	}
	//lock.Unlock()
	//主线程休眠10秒，为让协程完成任务
	time.Sleep(time.Second * 10)
	lock.Lock()
	for k, v := range myMap02 {
		fmt.Println("k = ", k, "; v = ", v)
	}
	lock.Unlock()
	//go build -race xxxx.go 运行时，增加参数-race参数，确实会发现现有资源竞争问题。
}
