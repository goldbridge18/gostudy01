package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func demoFile() {

	//file的叫法：file对象、file指针、file文件句柄
	file, err := os.Open("test.txt")

	//defer 当程序退出时执行
	defer file.Close()

	if err != nil {

		fmt.Printf("file error  %v \n", err)

	}
	fmt.Printf("file = %v \n", file)

	reader := bufio.NewReader((file)) //file是一个指针

	for {
		str, err := reader.ReadString('\n') //读到一个换行符结束
		if err == io.EOF {
			break
		}
		fmt.Println(str)
		//fmt.Println(err)
	}

	fmt.Println("over................")

}
