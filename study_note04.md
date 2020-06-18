面向对象编程

golang支持面型对象编程(oop)，并不是纯粹的面向对象语言。golang支持面向对象编程特性。

golang没有类，struct来实现oop特性。

golang面向对象编程非常简洁，去掉了oop语言的继承、方法重载、构造函数、析构函数、隐藏的this指针

golang有面向对象编程的继承、封装、和多态特性。golang的继承没有extends关键字。继承是通过匿名字段来实现。

oop本身就是语言类型系统的一部分，通过接口关联，耦合性低。也就是说golang中面向接口编程是非常重要的特性。

1、结构体 struct

语法：

type 结构体名称 struct{
    field1 type
    field2 type
}

field的声明同变量，示例： 字段名 类型
字段类型可以为基本类型、数组、引用类型、指针、slice、map

如果结构体的字段类型是：指针、slice、map的零值都为nil，即还没有分配空间，需要先make才能使用。

2、结构体声明方式

方式一：
 var person Person
 方式二：{}
 var person Person = Person{

p := Person{}
 
 方式三：&

 var person *Person = new(Person)  //结构体指针

方法四： {}

var person *Person = &Person{} //结构体指针

结构体指针的字段访问方式，标准方式：(*ptr).field = _简化方式：ptr.field= _

