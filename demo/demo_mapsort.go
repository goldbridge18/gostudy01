package main

import (
	"fmt"
	"sort"
)


func demoMapsort(){

	m := make(map[int]int,10)
	m[11] = 9
	m[10] = 7
	m[23] = 11
	m[5] = 16

//如果按照map的key的顺序进行排序
//1. 先将map的key放入一个切片中
//2 对切片进行排序
// 3 遍历切片，按照key的顺序输出值

var keys []int

for  k,_ := range m {

//	fmt.Println(k)
	keys = append(keys,k)
}

//keys值排序
sort.Ints(keys)
fmt.Println(keys)

for _, k := range keys{

	fmt.Printf("m[%v] = %v \n",k,m[k])
}

}
