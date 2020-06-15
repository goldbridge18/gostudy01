package main

import (
	"fmt"
)

func demoMapSlice() {

	var monsters []map[string]string        //声明一个map切片
	monsters = make([]map[string]string, 2) //为切片创建空间

	if monsters[0] == nil {
		monsters[0] = make(map[string]string) //为map创建一个空间
		monsters[0]["name"] = "蜘蛛精"
		monsters[0]["address"] = "盘丝洞"
		monsters[0]["num"] = "six"

	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string) //为map创建一个空间
		monsters[1]["name"] = "白骨精"
		monsters[1]["address"] = "白骨洞"
		monsters[1]["num"] = "thrid"

	}

	//append 动态增加 元素
	newmonsters := map[string]string{

		"name":    "玉兔精",
		"address": "月宫",
		"num":     "one",
	}
	fmt.Println(newmonsters)
	monsters = append(monsters, newmonsters)
	fmt.Println(monsters)
}
