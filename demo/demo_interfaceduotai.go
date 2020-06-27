package main

import "fmt"

type Usb01 interface {
	start()
	stop()
}

type Phone01 struct {
}

func (phone Phone01) start() {
	fmt.Println("开机。。。。。")
}

func (phone Phone01) stop() {
	fmt.Println("关机。。。。。")
}

type Camera struct {
}

func (Camera Camera) start() {
	fmt.Println("开机。。。。。")
}

func (camera Camera) stop() {
	fmt.Println("关机。。。。。")
}

type Computer01 struct {
}

func (computer Computer01) Working(usb Usb01) { // usb变量根据传入实参确定Phone还是Camera；多态参数
	usb.start()
	usb.stop()
}

func demoInterfaceduotai() {

	//创建结构体变量
	computer := Computer01{} //Computer 指向了接口Usb
	phone := Phone01{}
	camera := Camera{}

	computer.Working(phone)
	computer.Working(camera) //computer既接收了phone也接收了camera，实现了多种形态，即多态

	//定义一个Usb接口数组，可以存放Phone和Camera的结构体
	//实现多态数组
	var usbarr [3]Usb01
	usbarr[0] = Phone{}
	usbarr[1] = Phone01{}
	usbarr[2] = Camera{}
	//因为数组只能存放一种数据类型，如果要实现多种数据类型，则需要interface实现
	fmt.Println(usbarr)
}
