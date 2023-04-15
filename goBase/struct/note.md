# go的结构体
## 定义
go语言中数组可以存储同一类型的数据，但在**结构体中我们可以为不同项定义不同的数据类型**

结构体是由**一系列具有相同类型或不同类型的数据构成的集合**

结构体表示一项记录，比如保存图书馆的书籍记录等等

## 定义结构体
结构体定义需要使用type和struct语句

struct语句定义一个新的数据类型，结构体中有一个或多个成员

type语句设定了结构体的名称

格式如下：
```go
//定义一个新的结构体
type struct_val_type struct {
    //前面是成员名称 后面为类型
    member_name1 definition
    member_name2 definition
    member_name3 definition
    member_name4 definition
}

//当我们定义好一个结构体类型后，我们就可以用于变量的声明
val1 := struct_val_type {member_name1}
```

## 访问结构体成员
访问结构体的成员，其实跟js编程语言中的对象差不多，可以直接通过 . 一个属性来获取这个成员

> struct_name.option

go语言如果要访问结构体成员，需要使用 . 操作符格式为
> 结构体.成员名

## 结构体作为函数参数
我们在go语言当中也可以像其他数据类型一样将结构体类型作为参数传递给函数，并以以上示例的方式来进行访问

```go
type books struct {
    name string
    size string
    id int
}

//将结构体作为参数来传入函数当中 
func printBookName(book books) {
    //获取结构体当中的name属性
    fmt.Println(book.name)
}
```

## 结构体指针
我们可以定义指向结构体的指针类似于其他指针变量

格式为：
> var struct_hand *struct_name

当我们想查看这个结构体的存储地址时：
> struct_hand = &struct_name

为什么要在结构体引入指针这个概念呢？其实在指针的文档中我们提到，一般go语言的函数进行修改操作的时候，会把参数拷贝一层，这样我们无法直接操作源数据，引入指针的原因正因为如此

引入指针的概念可以更方便于我们针对结构体当中的结构体成员进行修改

具体操作可以见对应的操作文件struct.go