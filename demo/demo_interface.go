package main

import "fmt"

type Usb interface {
	start()
	stop()
}

type Phone struct {
}

func (phone Phone) start() {
	fmt.Println("开机。。。。。")
}

func (phone Phone) stop() {
	fmt.Println("关机。。。。。")
}

type Computer struct {
}

func (computer Computer) Working(usb Usb) {
	usb.start()
	usb.stop()
}

func demoInterface() {

	//创建结构体变量
	computer := Computer{} //Computer 指向了接口Usb
	phone := Phone{}

	computer.Working(phone)
}
