package main

import "fmt"

func demoAppend() {

	slice := []int{1, 2, 3, 4}
	slice = append(slice, 6, 7, 8, 9) //追加元素
	fmt.Println("slice is value  ", slice)
	//追加切片，

	slice1 := []int{20, 21, 22, 23}

	slice1 = append(slice1, slice...) //3个点是固定写法。
	fmt.Println(slice1)

	str := "zbac"
	slice3 := str[1:3]
	fmt.Println(slice3)
	arr3 := []byte(str)
	arr3[1] = 'm'
	str = string(arr3)
	fmt.Println(str)

	str1 := "中国共产党千岁！"

	arr4 := []rune(str1)

	for k, v := range arr4 {
		fmt.Println("k ", k, "  v", v)
	}
	arr4[5] = '万'

	str1 = string(arr4)
	fmt.Println(str1)

}
