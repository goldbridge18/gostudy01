1、map

map是key-value数据结构，又称为字段或者关联数组。 

基本语法：
var map变量名 map[keytype]valuetype

key可以是很多中类型，比如bool，数字，string，指针，channel，还可以是只包含前面几个类型的接口，结构体，数组，

key type 通常为int、string

valuetype 通常为数字(整数，浮点数)，string，map，struct

注意：slice，map 还有function不可以，因为这几个没法用 == 来判断。


map声明的例子：
    
    var a map[string]string
    var a map[string]int
    var a map[int]string
    var a map[string]map[string]string

注意：声明是不会分配内存的，初始化需要make，分配内存后才能赋值和使用。map的key不能重复，map的值是无序的

make(type, size intertype) 
内置函数make分配并初始化一个类型为slice、map、channel的对象。其第一个实参为类型，而非值。make的返回类型与其参数相同，而非只想他的指针。

slice：size指定了其长度。该切片的容量等于其长度。切片支持第二个整数实参来指定不同的容量，他必须不小于其长度，make([]int, 0 ,10) 分配一个长度为0，容量为10的切片

map: 初始化分配的创建取决于size，但产生的映射为0.size可以省略，这种情况下会分配一个小的起始大小。

channel：通道的缓存根据指定的缓存容量初始化。若size为零或被省略，该信道即为无缓存的。



	//多个map声明使用
```
	studentMap := make(map[string]map[string]string)

	studentMap["NO.001"] = make(map[string]string)
	studentMap["NO.001"]["name"] = "tom"
	studentMap["NO.001"]["age"] = "19"
	studentMap["NO.001"]["sex"] = "m"
```


2、map的增删改查

delete (mapname,mapkey) //删除map元素，如果不存在，不会操作，也不报错

```
//删除所有的key
//1.遍历所有的key 逐个删除
//2.直接make一个新空间

	studentMap = make(map[string]map[string]string)
	fmt.Println(studentMap)
    ```

```
	// 查找
	val, ok := studentMap["NO.002"]

	fmt.Println("查找： ", val, ok) // 如果存在ok 等于true
	if ok {
		fmt.Println("查找的结果：", val)
	} else {
		fmt.Println("not found")
	}

```

3、map遍历

map只能使用个 for-range遍历，因为map的key可能有不为数字或者数字不连续

4、map的切片