# Go语言的类型转换
类型转换用于将一种数据类型的变量转换为另外一种类型的变量，其基本格式如下：

> type_name(expression)

type_name 为类型，expression 为表达式

## 数值类型转换
将整型转换为浮点型
```go
//定义一个整型变量
var a int = 10

//将上面的整型变量转换为浮点型
var b float64 = float64(a)
```

## 字符串类型转换
### 字符串转整型
将一个字符串转换为另一个类型，可以使用以下语法：
```go
var str string = "10"
var num int 

num, _ = strconv.Atoi(str)
```

以上代码将字符串变量 str 转换为整型变量 num。

注意，strconv.Atoi 函数返回两个值，第一个是转换后的整型值，第二个是可能发生的错误，我们可以使用空白标识符 _ 来忽略这个错误

### 整数转字符串
与上面的例子类似，将整数转为字符串的方法可以使用：

> strconv.Itoa()

示例如下：
```go
num := 123
str := strconv.Itoa(num)

fmt.Print(str)
//输出"123"
```

### 字符串转浮点数
将字符串转换为浮点数的方法可以使用：

> strconv.ParseFloat(str, 64)

示例如下：
```go
str := "3.14"
//类型转换
num, err := strconv.ParseFloat(str, 64)
if err != nil {
        fmt.Println("转换错误:", err)
    } else {
        fmt.Printf("字符串 '%s' 转为浮点型为：%f\n", str, num)
    }
```

### 浮点型转换为字符串
将字符串转换为浮点数的方法可以使用：

> strconv.FormatFloat(num, 'f', 2, 64)

实例如下：
```go
num := 3.14
str := strconv.FormatFloat(num, 'f', 2, 64)

fmt.Printf("浮点数 %f 转为字符串为：'%s'\n", num, str)
//输出3.140000 "3.14"
```