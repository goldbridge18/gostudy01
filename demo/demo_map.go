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

	//声明方式二

	ma1 := make(map[int]int, 4)
	ma1[1] = 1
	ma1[2] = 2
	fmt.Println(ma1)

	//声明方式三
	ma2 := map[int]int{
		9: 9,
		8: 8,
		7: 7,
	}
	fmt.Println(ma2)

	//多个map

	studentMap := make(map[string]map[string]string)

	studentMap["NO.001"] = make(map[string]string)
	studentMap["NO.001"]["name"] = "tom"
	studentMap["NO.001"]["age"] = "19"
	studentMap["NO.001"]["sex"] = "m"

	fmt.Println(studentMap)

	//增加
	studentMap["NO.002"] = make(map[string]string)
	studentMap["NO.002"]["name"] = "jack"
	studentMap["NO.002"]["age"] = "18"
	studentMap["NO.002"]["sex"] = "F"

	fmt.Println(studentMap)

	// 查找

	val, ok := studentMap["NO.002"]

	fmt.Println("查找： ", val, ok) // 如果存在ok 等于true
	if ok {
		fmt.Println("查找的结果：", val)
	} else {
		fmt.Println("not found")
	}

	//delete

	delete(studentMap, "NO.001")
	fmt.Println(studentMap)

	//删除所有的key
	//1.遍历所有的key 逐个删除
	//2.直接make一个新空间

	studentMap = make(map[string]map[string]string)
	fmt.Println(studentMap)

}
