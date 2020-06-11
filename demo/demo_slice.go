package main

import "fmt"

func demoSlice() {

	//第一种：
	var intarr [5]int = [...]int{1, 2, 3, 4, 5}

	slice := intarr[1:3]
	for k, v := range slice {
		fmt.Println("slice  key is ", k, "  slice values is ", v)
	}
	fmt.Println("intarr ", intarr)
	fmt.Println("slice meta :", slice)
	fmt.Println("slice length is ", len(slice))
	fmt.Println("slice 容量：", cap(slice)) //容量一般是切片大小的2倍，但是它是可以动态增加的

	fmt.Println("intarr[1] 的地址", &intarr[1])
	fmt.Println("slice[0]", &slice[0])

	//第二种：

}
