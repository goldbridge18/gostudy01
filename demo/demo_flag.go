package main

import (
	"fmt"
	"flag"
)
func demoFlag()  {
	//定义变量，用于接收命令的参数值
	var user string
	var pwd string
	var host string
	var port int

	//&user 就是接收命令行的输入的 -u 的参数
	// "u" 就是 -u指定参数
	//"" 默认值
	//"用户名默认为空"
	flag.StringVar(&user,"u","","用户名默认为空")
	flag.StringVar(&pwd,"pwd","","密码默认为空")
	flag.StringVar(&host,"h","localhost","")
	flag.IntVar(&port,"port",3306,"")


	//调用Parse方法来转换参数值
	flag.Parse()


	fmt.Println(user,pwd,host,port)


}