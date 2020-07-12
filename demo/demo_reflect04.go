package main
import(
	"fmt"
)

//定义了一个Monster712结构体
type Monster712 struct {
	Name string
	Age int
	Score float32
	Sex string
}
//声明一个方法,显示s的值
func (s Monster712) Print()  {
	fmt.Println("-------------start-------------")
	fmt.Println(s)
	fmt.Println("---------------end-------------")
}
//声明一个方法 返回两个数的和
func  (s Monster712) GetSum(n1 ,n2 int) int  {
	return n1 + n2
}

//声明一个方法接收四个值
func (s Monster712) Set(name string,age int, score float32, sex string)  {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func demoReflect04()  {
	
	var  a Monster712
	a.Set("xxxxx",18,98.0,"x")
	a.Print()
}