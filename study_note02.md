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

![tupian](https://github.com/goldbridge18/imagefile/blob/master/2020-06-10%2000-20-13%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE.png)