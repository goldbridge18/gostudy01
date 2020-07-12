1、浮点类型 float32 float64 浮点型声明默认为float64位 浮点型= 符号位+指数位+尾数位

科学计数法：

5.1234E2 表示浮点数，E表示科学计算法 5.1234*10^2
十进制计数法：

1.234
.1234 等价 0.1234
2、字符类型

golang没有专门的字符类型。如果保存字符(字母)，用byte保存。 字符串是一串固定长度的字符链接起来的字符序列。(字符串是有字节组成的)

如果var c byte = '字' 会overflow溢出；使用c int='字'

如果保存的字符在ASCII表上比如A-Z，a-z,0-9 直接用byte 如果保存的字符对应的码值大于255，则使用int 如果按照字符的方式输出，使用fmt.Printf()

3、string 字符串

字符串时不可变的

字符串的两种形式：1、双引号会识别转义字符 2、反引号以原生的字符串输出，包括换行和特殊字符

4、数据类型转换

var a int32 = 100； var a1 int64 = int64(a)//把a的值转换成int64

被转换的是变量存储的数据(即值)，变量本身的数据类型并没有改变

如果int64转换成int8，编译不会报错，只是转换结果按照溢出处理

6、基本类型转string类型

方式1：
使用fmt.Sprintf("%参数",表达式)
Sprintf根据format参数生成格式化的字符串并返回该字符串。

方式2：strconv()包函数
7、指针 使用&获得一个变量的地址

var i int =10 var ptr *int = &i

ptr 是指针变量 ptr的类型是 *int ptr本身的值是&i

*ptr 是指针的值 &ptr 是指针的地址

&i 是i的地址 10 是地址里的值，也就是i的值

值可以改变，地址不可以改变

var i int = 11 var ptr *int32 = &i //错误，必须时同类型的

8、值类型和引用类型

值类型：基本数据类型 int系列，float系列，bool，string，数组和结构体

引用类型：指针、slice切片、map、管道chan、interface等都是引用类型。

使用特点：

值类型：变量直接存储值，内存通常(受逃逸分析的影响可能不在栈中)在栈中分配

引用类型：变量存储的是一个地址，这个地址对应的空间才真正存储数据，内存通常在堆上分配。当没有任何变量引用这个地址时，改地址对应的数据空间成为一个垃圾，由GC来回收。

9、三大流程控制

    顺序控制    
    分支控制
    循环控制
10、 字符串遍历 传统方法：
```
var str string = "h1h1h1h1h1h1h1"

for i := 0; i < len(str);i++{ fmt.Printf("%c \n",str[i]) }
````
方法二：利用for-range str := "abc=abc" for index,val := range str{ fmt.Printf("index=%d,val=%c \n",index,val) }

11、递归调用

一个函数在函数体内有调用了本身，称为递归调用

12、函数

为完成某一个功能的程序指令的集合，称为函数

13、defer

1.当go执行到一个defer时，不会立即执行defer后的语句，而是将defer后的语句压入到一个栈中，然后继续执行下一行函数。
2. 当函数执行完毕后，在从defer栈中，一次从栈顶取出语句执行（栈的原则：先进后出）
3. 在defer将语句放入到栈时，也会将相关的值拷贝同时入栈。
14、函数参数的值传递方式

值类型：基本数据类型int系列，float系列，bool、string、数组和结构体struct 引用类型：指针、slice切片、map、管道chan、interface等

15、字符串处理函数

len()

[]rune() // 字符串遍历，把字符串转换正切片。解决处理中文的问题，因为一个中文占用3个字节
```
strconv.Atoi() //字符串转整数

num ,err = strconv.Atoi("123") //err可以用来判断错误

str = strconv.Itoa(12345) //整数转字符串

//字符串转[]byte var bytes = []byte("hello world") fmt.Printf("byte=%v\n",bytes)

//[]byte 转字符串
str1 = string([]byte{97, 98, 99})
fmt.Printf("str = %v \n", str1)

//10进制转2\8\16
var str2 = strconv.FormatInt(12, 2)
fmt.Println("二进制：", str2)
//查找字符串是否在指定的字符串中 str3 := strings.Contains("hello word", "wo") fmt.Println("查找wo的是否存在", str3)

//统计指定字符串的个数
str4 := strings.Count("123123456123", "123")
fmt.Println("total :", str4)

//不区分大小写的字符串比较（== 时区分字母大小写的）
str5 := strings.EqualFold("abc", "ABC")
fmt.Println(str5)

fmt.Println("比较 ：", "abc" == "ABC") //区分大小写，返回false

//返回指定字符在字符串第一次出现的index值，如果没有返回-1：
index := strings.Index("00123,456", "123")
fmt.Println("index :", index) //位置从0开始

// strings.LastIndex 返回指定字符在字符串的最后一个index

//替换字符串
str6 := strings.Replace("abc,abc,dbs", "abc", "go", 1) //最后的数字为-1表示全部替换，否则表示希望替换几个
fmt.Println("replace after new_string: ", str6)

//分割字符串,拆分成数组

str7 := strings.Split("abc,abc", ",")
fmt.Println("split :", str7)

//将字符串的字母进行大小写转换

str8 := strings.ToLower("HELLO worlD")
fmt.Println(str8)
//ToUpper 转大写
```
16、日期和时间函数

引入包 time
```
now := time.Now().Year() fmt.Println("time :",now)

	//"2006/01/02 15:04:05" 格式化日期的固定写法。
now1 := time.Now().Format("2006/01/02 15:04:05")
fmt.Println("time :",now1)

/*时间常量
Nanosecond Durarion=1//纳秒
Microsecond =1000*Microsecond //毫秒
Millisecond =1000*Millsecond //秒
Scond =60*Second //分钟
Minute = 60*Seond //分钟
Hour =60*Minute//小时
*/

starttime := time.Now()
fmt.Println(starttime)
time.Sleep(time.Second*60)
endtime := time.Now()
fmt.Println(endtime)
```
17、内置函数 buildin function
```
len:用来求长度
new：用来分配内存，主要用来分配值类型，比如int、float32、struct等返回的时指针

num := new(int)

fmt.Printf("num的类型：%T，num的值是：%v，num的指针地址：%v，num指向的值%v", num, num, &num, *num)

//num的类型：*int，num的值是：0xc000018128，num的指针地址：0xc00000e030，num指向的值0

// 0xc000018128是num指针的地址；0xc00000e030是指针指向值的地址 

//0是指针指向的值。如果指针没有值则为默认0


make：用来分配内存，主要用来分配引用类型，如channel、map、slice
```
18、错误处理机制 go没有try...catch...finally go中引入的处理方式为：defer、panic、recover

func test(){ //使用defer+recover来捕捉和处理异常 defer func (){ err := recover() //recover()内置函数，可以捕捉异常 if err !=nil{ //说明捕捉到的错误 fmt.Println("error : ",err) } }() num1 := 10 num2 := 0 res := num1/num2 fmt.Println("res :",res)

}

func run(){

test()//报错不阻碍下一面的代码执行
for i := 0; i <10; i++ {
	fmt.Println("next ....")
}
}

自定义处理错误 使用errors.New和panic 内置函数。

error.New("错误说明")，返回一个error类型的值，表示一个错误
panic内置函数，接收一个interface{}类型的值（也就是任何值）作为参数，可以接收error类型的变量，**输出错误信息，并退出程序**


19\ const修饰常量

- 1.常量使用const修饰
- 2.常量在定义的时候,必须初始化赋值
- 3.常量不能修改
- 4.常量只能修饰bool\int\float\string

语法:
```
const identifier [type] = value
```

常量写法:

方式一

```
const (
    a = 1
    b = 2
)
```

方式二:

```
const(
    a = iota //表示给a赋值为0
    b   //b的写法表示在a的基础上+1 ,b的值为1
    c  //c的写法表示在b的基础上+1 ,c的值为2

)
```

常量的首字母大小写,决定public还是private

注意:

```

const(

a = iota //a = 0
b = iota //b = 1
c , d  = iota,iota // c = 2 ;d = 2

)
```
不换行不加1