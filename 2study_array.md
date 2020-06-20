1、数组
```
格式：var  arr [n]int/string/float32...
                var arr  = [n]int{1,2,3...n}
                 var arr  = [...]int{1,2,3...n} //一定要是3个点 
```

```
下标为0是内存地址是数组的地址，按照类型的大小步长增加
数组的类型可以是值类型也可以是引用类型，但是不可以混用
数组的类型确定之后，数组的值一定要统一；数组长度确定之后，则不可动态变化，否则越界。动态变化需要slice
go的数组属于值类型，在默认情况下值传递，因此会进行值拷贝。数组间不会相互影响。如果想修改数组里的值，可以使用引用传递（指针）。

	//数组定义格式之一
	var arr0 [3]int
	arr0[0] = 1
	arr0[1] = 2
	arr0[2] = 3

	for k, v := range arr0 {
		fmt.Println(k, v)
	}
	//数组定义格式二，
	var arr [3]int = [3]int{1, 2, 3} //还可以这样写 var arr [3]int = [...]int{1,2,3}
	var arr1 [3]int = [...]int{1, 2, 3}
	//遍历数组方法一
	for i := 0; i < len(arr); i++ {
		fmt.Println("key is ", i, "values is", arr[i])
	}
	for i := 0; i < len(arr1); i++ {
		fmt.Println("arr1 array key is ", i, "values is", arr[i])
	}
	//for -range  遍历数组  k,v := range arr ,k和v可以用_表示站位符忽略

	//k和v的作用域只能在for循环中使用
	for k, v := range arr {
		fmt.Println("key is ", k, ", values is ", v)
	}
```
2、切片 slice
```
切片可以简单理解为动态的数组，但是和数组是有区别的。
切片是数组的引用，因此切片是一个引用类型。
切片的长度是可以变化的，是一个动态变化的数组
````
语法格式： 

var 切片名 []类型  

var slice := intarr[startIndex:endIndex]  //切片的取值范围是左闭右开

slice底层的数据结构可以理解为一个结构体

type slice struct{

	ptr *[2]int
	len int
	cap int
}

![slice内存分布图](https://github.com/goldbridge18/imagefile/blob/master/goimage/2020-06-11%2023-27-14%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE.png)

3、切片的使用方式：

第一种：定义一个切片，然后让切片引用一个已创建的数组
```
var intarr [5]int = [...]int{1, 2, 3, 4, 5}

	slice := intarr[1:3]
	for k, v := range slice {
		fmt.Println("slice  key is ", k, "  slice values is ", v)
	}
	fmt.Println("intarr ", intarr)
	fmt.Println("slice meta :", slice)
	fmt.Println("slice length is ", len(slice))
	fmt.Println("slice 容量：", cap(slice)) //容量一般是切片大小的2倍，但是它是可以动态增加的

	fmt.Println("intarr[1] 的地址", &intarr[1])
	fmt.Println("slice[0]", &slice[0])
```

第二种：通过make来创建切片

基本语法：var 切片名 []type = make([]type,len,[cap])
说明：type：数据类型；len：大小；cap：容量，可选。

```
	//make使用
	var slice []float64 = make([]float64, 10, 20)
	slice[6] = 100
	slice[3] = 99
	fmt.Println("slice ", slice)  //返回数组 [0 0 0 99 0 0 100 0 0 0]，这个数组是对外不可见的。
	fmt.Println("cap is size ", cap(slice))


```

![make内存分布](https://github.com/goldbridge18/imagefile/blob/master/goimage/2020-06-11%2023-49-09%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE.png)


a、通过make方式创建分片可以指定切片大小和容量
b、如果没有切片的各个元素赋值，那么就会使用默认值(int返回0，float返回0，string返回""，bool返回false)
c、通过make创建的切片对应的数组是由make底层维护，对外不可见，只能通过slice去访问各个元素。

第三种：定义一个切片，直接指定具体数组，使用原理类似make方式


切片使用注意事项go：
	1.切片初始化时 var slice = arr[startindex:endindex] ,左开右闭
	2.切片初始化时，仍然不能越界。在[0-len(arr)]范围之间，但是可以动态增长。
	3.书写：
				var slice = arr[0:endindex]可以简写 var slice = arr[:endindex]
				var slice = arr[startindex:len(arr)] 可以简写 var slice =arr[startindex:]
				var slice = arr[0:len(arr)] ---->var slice = arr[:]
	4.cap是一个内置函数，用于统计切片的容量，即最大可以放入多少元素。
	5.切片可以再切片
	6.切片定义之后，不能使用，因为本身是一个空，需要让其引用到一个数组或者make一个空间。

	4、slice 的append 函数可以对切片进行动态追加。
	在追加元素为切片时,固定格式为 ：slice1 = append(slice1, slice...) //3个点是固定写法。
```
		slice := []int{1, 2, 3, 4}
	slice = append(slice, 6, 7, 8, 9) //追加元素
	fmt.Println("slice is value  ", slice)
	//追加切片，

	slice1 := []int{20, 21, 22, 23}

	slice1 = append(slice1, slice...) //3个点是固定写法。
	fmt.Println(slice1)
```
append的底层原理：
	- 1.切片append操作的本质就是对数组扩容。
	- 2.go底层创建一个新的数组newArr，按照扩容后的大小进行创建新的数组
	- 3.将slice原来包含的元素拷贝到新的数组newArr
	- 4.slice重新引用到newArr
	- 5.注意newArr是在底层来维护，程序员不可见。

	5、切片的拷贝

	copy(para1,para2) //para1和para2都是切片类型

	6、string和slice

string的底层是byte

sring是不可变的，修改string类型，先将string转为[]byte或[]rune,修改之后转化为string

7、数组排序

排序分类：

	- 内部排序： 指将需要处理的所有数据都加载到内部存储器中进行排序。包括：交换式排序、选择式排序和插入式排序
	交换排序：冒泡排序、快速排序
	- 外部排序：数据量大，无法全部加载到内存中，需要借助外部存储进行排序。包括：合并排序和直接排序

8、递归函数

一个函数在函数体内又调用本身，成为递归调用 


9、二维数组

