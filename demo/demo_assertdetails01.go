package main

import "fmt"

type  Usb27 interface{
	start()
	stop()
}

type Phone27 struct{
		Name string
}

func (phone27 Phone27) start(){
	fmt.Println("start opening telephone.......")
}
func (phone27 Phone27) call(){
	fmt.Println("calling telephone.......")
}


func (phone27 Phone27) stop(){
	fmt.Println("stop close telephone....")
}

type Camera27 struct{
	Name string
}

func (camera27 Camera27) start(){
	fmt.Println("start opening camera.......")
}
func (camera27 Camera27) stop(){
	fmt.Println("stop close camera.......")
}

type Computer27 struct{

}

func (computer27 Computer27) Wroking(usb27 Usb27) {
	usb27.start()

	//类型断言，实现phone的call()功能
	if phone27, ok := usb27.(Phone27); ok == true {
		phone27.call()
	}

	usb27.stop()
}



func demoAssertdetails01(){

	var   usbArr [3]Usb27
	usbArr[0] = Phone27{"huawei"}
	usbArr[1] = Phone27{"redmi"}
	usbArr[2] = Camera27{"佳能"}

	fmt.Println(usbArr)

	var computer27 Computer27
	for k , v := range usbArr{
		fmt.Println(k,v)
		//如果给phone添加一个call() ,在computer27中利用断言实现
		computer27.Wroking(v)
	}
}
