package main

import "fmt"

func demoMap() {
	//声明map
	var ma map[string]string
	//使用map前，需要使用make为map分配数据空间 ，否则报错
	ma = make(map[string]string, 10) // make(type,size intertype)

	ma["num1"] = "张三"
	ma["num3"] = "王五"
	ma["num2"] = "李四"
	ma["num3"] = "李四001" //key重复，则被后面的重复的值覆盖
	// map的值是无序的
	fmt.Println(ma)

}
