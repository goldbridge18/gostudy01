一、协程

1、进程和线程说明
- 1.进程是程序在操作系统中的一次执行过程，是系统进行资源分配和调度的基本单位
- 2.线程是进程的一个执行实例，是程序自行的最小单位，他是比进程更小的能独立运行的基本单位。
- 3.一个进程可以创建和销毁多个线程，同一个进程中的多个线程可以并发执行。
- 4.一个程序至少有一个进程，一个进程至少有一个线程。

2、并发和并行
- 1.多线程程序在单核上运行是并发

    并发特点：
        - 1.多个任务作用在一个cpu
        - 2.从微观的角度看，在一个时间点上，有且只有一个任务在执行。

- 2.多线程程序在多核上运行是并行

    并行特点：
         - 1.多个任务作用在多个cpu
        - 2.从微观的角度看，在一个时间点上，有多个个任务在执行。

![并发和并行示意图](https://github.com/goldbridge18/imagefile/blob/master/goimage/2020-07-05%2016-04-58%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE.png)


3、goroutine

go的主线程:一个go线程上，可以启多个协程。协程是轻量级的线程【编译器做优化】。

go协程的特点：
- 有独立的栈空间
- 共享程序堆空间
- 调度由用户控制
- 协程是轻量级的线程

协程的流程图：
![..](https://github.com/goldbridge18/imagefile/blob/master/goimage/2020-07-05%2018-09-30%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE.png)

总结：
- 1.主线程是一个物理线程，直接作用在cpu上，重量级非常耗费cpu资源
- 2.协程从主线程上开启，是轻量级的线程，是逻辑态，对资源耗费相对小。
- 3.协程是golang的重要特点，可以轻松开启上万个协程。其他语言的并发机制是一般基于线程的，开启过多的线程，资源耗费大。

4、goroutine的调度模型---MPG

- 1.M：操作系统的主线程（物理线程）
- 2.P：协程执行需要的上下文
- 3.G：协程

![1...](https://github.com/goldbridge18/imagefile/blob/master/goimage/2020-07-05%2018-41-02%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE.png)

![2...](https://github.com/goldbridge18/imagefile/blob/master/goimage/2020-07-05%2018-41-44%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE.png)

![3...](https://github.com/goldbridge18/imagefile/blob/master/goimage/2020-07-05%2018-42-30%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE.png)

5、runtime

runtime包提供和go运行时环境的互操作，如控制go程的函数

NumCPU函数：返回本地的逻辑cpu个数

GOMAXPROCS函数:设置可同时执行的最大CPU数。

二、channel管道

- 1.channel本质就是一个数据结构队列
- 2.数据是先进先出【FIFO：first in first out】
- 3.线程安全，多goroutine访问时，不需要加锁，就是说channel本身就是线程安全的
![channel](https://github.com/goldbridge18/imagefile/blob/master/goimage/2020-07-07%2023-26-04%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE.png)
- 4.channel是有类型的，一个string的channel只能存放string类型的数据。

**2.1 channel声明：**

语法：
```
var 变量 chan 数据类型
```
例子：
```
var intChan chan int(intChan用于存放int数据)
var mapChan chan map[int]string(mapChan用于存放map[int]string类型)
var perChan chan *Person// 结构体
```
- 1.channel是引用类型
- 2.channel必须初始化才能写入数据，即make后使用
- 3.管道是有类型的intChan只能写入整数int
```
var intChan chan int
intChan = make(chan int,10)
```

**2.3 channel注意事项：**
- 1.channel中只能存放指定的数据类型
- 2. channel的数据放满后，就不能继续放入数据
- 3. 如果从channel取出数据后，取出的数据的位置被释放 ，可以继续放入数据
- 4.在没有使用协程的情况下，如果channel数据取完了，再取就会报dead lock
```
fatal error: all goroutines are asleep - deadlock!
```

如何实现channel存放多种数据类型

```
package main

import "fmt"

type Cat struct {
	Name string
	age  int
}

func demoChannel01() {

	//定义一个存放任意数据类型的channel，可以存放3个数据

	//声明一个接口类型的channel
	var allChan chan interface{} //空接口类型 的channel，接口可以存放任意数据类型
	allChan = make(chan interface{}, 3)

	//allChan := make(chan interface{}, 3) //声明一个接口类型的channel

	allChan <- 100
	allChan <- "cat007"

	/*
		var cat Cat //初始化Cat结构体
		cat.age = 18
		cat.Name = "008"
	*/
	/*
		var cat Cat = Cat{"flower cat", 10} //初始化Cat结构体
	*/
	cat := Cat{"flower cat", 10}
	allChan <- cat

	//channel是先进先出，如果要取出cat的名字，就需要先把前两个取出来，才能取出Cat结构体的数据
	<-allChan //取出数据，不赋值给任何对象
	<-allChan

	getCat := <-allChan //取出Cat结构体

	fmt.Printf("getCat is %T, getCat is %v\n", getCat, getCat)

	//获取cat的名字，如果使用getCat.Name直接报错
	catName := getCat.(Cat).Name  //使用断言
	fmt.Println("cat's name is ", catName)

}
```
使用断言 ：catName := getCat.(Cat).Name

类型断言，由于接口是一般类型，不知道具体类型，如果要转成具体类型，就需要使用类型断言。


**2.4 channel的关闭**

使用内置函数close可以关闭channel，当channel关闭后，不能继续向channel写入，但是可以读取该channel的数据。

**channel遍历**
利用for-range ，不能使用普通for循环。
- 1. 在遍历时，如果channel没有关闭则会出现deadlock错误
- 2. 在遍历时，如果channel已经关闭则会正常遍历数据，遍历完后，就会退出遍历。