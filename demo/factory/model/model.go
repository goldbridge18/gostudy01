package model

type Student struct {
	Id    int
	Name  string
	Score float64
}

//class 首字母小写
type class struct {
	Level int
}

type class01 struct {
	name  string
	level int
}

//class结构体首字母小写因此只能在model包内使用
//如果想让其他的包使用可以通过工厂模式解决

func FactoryClass(n int) *class {
	return &class{
		Level: n,
	}
}

//如果字段名首字母为小写
func Factoryclass01(n int, str string) *class01 {

	return &class01{
		name:  str,
		level: n,
	}
}

func (class01 *class01) GetName() string {

	return class01.name
}

func (class01 *class01) GetLevel() int {
	return class01.level
}
