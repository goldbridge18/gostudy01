package main

import "fmt"

type Personb struct {
	id int64
	Name string
	age int
	gender string
}

func (personb *Personb) say() string{
	res :=fmt.Sprintf("name is [%v], id is [%v], age is [%v], gender[%v]",personb.Name,personb.id,personb.age,personb.gender)
	return res
}
	func demoObjectsay(){

		var personb Personb
		personb.age = 18
		personb.Name = "li lei"
		personb.id = 1001
		personb.gender ="male"

		fmt.Println(personb.say())
	}