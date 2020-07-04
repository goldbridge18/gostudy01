package main

import (
	"fmt"
	"os"
	"bufio"
)

func demoFfile02(){

	//创建一个文件test01.txt 
	// 1.打开文件
	filePath := "test01.txt"
	file , err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE,0755)
	if  err != nil {
		fmt.Println("Open file  flase")
		return
	}

	//关闭file句柄
	defer file.Close()

	//写入file的内容
	str := "中国共产党万岁！！！"

	//是bufio的缓存，
	writer := bufio.NewWriter(file)
	for i :=0; i<10; i++{
		writer.WriteString(str)
		writer.WriteString("\n")

	}

	//因为内容写入了bufio的缓存里；需要flush完成落盘操作
	writer.Flush()
	
	fmt.Println("写入完成")

}