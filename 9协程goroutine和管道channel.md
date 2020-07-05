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
