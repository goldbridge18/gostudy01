package main

import "fmt"

func demoMake() {

	//make使用
	var slice []float64 = make([]float64, 10, 20)
	slice[6] = 100
	slice[3] = 99
	fmt.Println("slice ", slice) //返回数组 [0 0 0 99 0 0 100 0 0 0]，这个数组是对外不可见的。
	fmt.Println("cap is size ", cap(slice))

	var strslice []string = []string{"jack", "tom", "jon"}
	fmt.Println("strslice = ", strslice)
	fmt.Println("strslice size ,", len(strslice))
	fmt.Println("strslice cap is ", cap(strslice)) //此时的容量和实际大小一样都为3

}
