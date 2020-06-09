package main

import "fmt"

func demoSlice() {
	//数组定义格式之一
	var arr0 [3]int
	arr0[0] = 1
	arr0[1] = 2
	arr0[2] = 3

	for k, v := range arr0 {
		fmt.Println(k, v)
	}
	//数组定义格式二，
	var arr [3]int = [3]int{1, 2, 3} //还可以这样写 var arr [3]int = [...]int{1,2,3}
	var arr1 [3]int = [...]int{1, 2, 3}
	//遍历数组方法一
	for i := 0; i < len(arr); i++ {
		fmt.Println("key is ", i, "values is", arr[i])
	}
	for i := 0; i < len(arr1); i++ {
		fmt.Println("arr1 array key is ", i, "values is", arr[i])
	}
	//for -range  遍历数组  k,v := range arr ,k和v可以用_表示忽略

	//k和v的作用域只能在for循环中使用
	for k, v := range arr {
		fmt.Println("key is ", k, ", values is ", v)
	}
}
