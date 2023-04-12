# go语言的数组与切片
## 数组(Array)
go语言的数组与其他语言的数组类似

go语言的数组定义方式如下：
```go
//指定长度
var a = [5]int{1,1,1,1,1}
a := [3]float32{1.0,1.0,1.0}

//非指定长度(使用...来表示)
a := [...]int{1,2,3,4}
```

当我们需要访问其中一个数值 ，直接访问其下标即可

但是使用数组的最大一个缺点就是**需要指定长度**，于是我们会经常使用到另一个类型——**切片**

## 切片(slice)
Go 语言切片是**对数组的抽象**

**Go 数组的长度不可改变**，在特定场景中这样的集合就不太适用，Go 中提供了一种灵活，功能强悍的内置类型切片("动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

定义切片的方式：
```go
//可以直接声明一个未确定大小的数组作为切片
var slice []int

//或使用make()函数来创建切片（常用方式）
var slice1 []type = make([]type, len())

//也可以简写为
slice2 := make([]type, len())

//定义一个新切片
slice_use := []int {1,2,3,4,5}
```

### 切片的截取
go的切片提供了截取的功能，类似于js中的slice，可以将我们的切片截取为任意一段子切片

这里 切片 的截取很类似于python中的语法，**使用:（冒号）进行分割**，左右两侧写入切取的左端和右端

示例：
```go
slice_use := []int {1,2,3,4,5}

//新建一个新的切片，截取上面下标1到4的元素
slice_new := slice_use[1:4]
```

### len()与cap()函数
切片是可以进行索引的，并且可以由**len()**方法**获取长度**

切片提供了计算容量的方法**cap()**可以**测量切片最长可以达到多少**

示例：
```go
//新建一个数组且长度为3
slice := make([]int, 3)

fmt.Println(len(slice), cap(slice))
//输出3 3
//前面的3表示这个切片的长度为3，后面的表示可以最长达到3

```

### 空切片
一个切片在未初始化之前默认为nil，长度为0

示例:
```go
//新建一个未初始化的切片 
slice := []int

fmt.Println(len(slice), cap(slice))
//输出0 0,这是应为切片此时只是初始化了，是空的
```

### append() 和 copy() 函数
如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容全都拷贝过来

append()函数的第一个参数传入需要添加元素的切片，后续填入参数即可

> 需要注意的是，append函数并不与其他编程语言的push方法类似，这是一个赋值操作，我们需要将append函数赋值给原切片才能完成添加的操作

copy()函数需要传入两个参数，第一个参数是粘贴的对象，第二个参数是被拷贝的切片

示例：
```go
//新建有值的切片
slice := []int {1,2,3}

//新建一个空切片
slice_empty := make([]int)

//向第一个切片添加元素
slice = append(slice, 1)

//拷贝slice的内容到slice_empty
copy(slice_empty, slice)
```


