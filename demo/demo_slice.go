package main

import "fmt"

func demoSlice() {
	var intarr [5]int = [...]int{1, 2, 3, 4, 5}

	slice := intarr[1:3]
	for k, v := range slice {
		fmt.Println("slice  key is ", k, "  slice values is ", v)
	}
	fmt.Println("intarr ", intarr)
	fmt.Println("slice meta :", slice)
	fmt.Println("slice length is ", len(slice))
	fmt.Println("slice 容量：", cap(slice)) //容量一般是切片大小的2倍，但是它是可以动态增加的

}
