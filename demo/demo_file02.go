package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount int
	NumCount int 
	SpaceCount int
	OtherCount int 
}

func  demoFile02() {

	//打开文件
	fileName := "test.txt"
	file, err := os.Open(fileName)
	//判断文件是否错误
	if  err != nil {
		fmt.Println("open file has ",err)
		return
	}
//关闭文件
	defer file.Close()

	//初始化 结构体
	var count CharCount
//创建一个reader，把文件内容读到buffer里
reader := bufio.NewReader(file)

//开始循环读取fileName的内容
for{
	str, err := reader.ReadString('\n')
	if err == io.EOF {
		break
	}
	//遍历str
	for _, v := range str {
		//fmt.Printf("k = %v; v = %v \n",k,v)
		switch  {
		case v >= 'a' && v <= 'z'  :
			fallthrough //穿透，下一行
		case v >= 'A' && v <= 'Z' :
			count.ChCount++
		case v == ' ' || v == '\t' :
			count.SpaceCount++
		case v >='0' && v <='9' :
			count.NumCount++
		default :
			count.OtherCount++
	
		}
		
	}
}
fmt.Printf("string = %v; number = %v;space = %v,other = %v\n",count.ChCount,count.SpaceCount,count.NumCount,count.OtherCount)
}
