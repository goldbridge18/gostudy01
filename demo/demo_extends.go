package main

import "fmt"

type Goods struct {
	Name string
	Price float64
}

type Brand struct {

	Name string
	Address string
}

type TV struct{
	Goods
	Brand
}

type PTV struct {
	*Goods
	*Brand
}

func demoExtends(){

	var   tv TV
	tv.Goods.Name = "小米"
	tv.Goods.Price = 99.0
	tv.Brand.Address = "bj"
	tv.Brand.Name = "0001"
	fmt.Println(tv)

	var tv1 = TV{Goods{"hw",100},Brand{"000","003"},}
	fmt.Println(tv1)

	var tv2 = TV{
		Goods{
			Name:"hw",
			Price:100,
			},
			Brand{
				Name:"000",
				Address: "003",
				},}
	fmt.Println(tv2)
 //指针

ptv := PTV{&Goods{"cw001",98.0},&Brand{"cw","004"},}
fmt.Println(*ptv.Goods)
fmt.Println(*ptv.Brand)


//匿名字段的使用

type NM struct{
	Goods
	int   //匿名字段  int类型的匿名字段不能重复出现，只能出现一次int类型的匿名子弹
}

var nm NM
nm.int = 110  //匿名字段，可以直接使用使用字段类型  ； 

fmt.Println(nm.int)

}