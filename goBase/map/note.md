# Map(集合)
map是一种**无序的键值对的集合**

map最重要的一点是通过**key**来快速检索数据，key作为索引，来指向数据的值

map是一种集合，所以我们可以像迭代数组和切片那样迭代它，不过，map是无序的，遍历map时返回的键值对的顺序是不确定的

> 在获取 Map 的值时，如果键不存在，返回该类型的零值，例如 int 类型的零值是 0，string 类型的零值是 ""。

> Map 是引用类型，如果将一个 Map 传递给一个函数或赋值给另一个变量，它们都指向同一个底层数据结构，因此对 Map 的修改会影响到所有引用它的变量。

## 定义map
可以使用**make函数**或者**map关键字**来定义一个map：

```go
//make函数的方式
map_variable := make(map[KeyType]ValueType, initialCapacity)

//map关键字的方式
//中括号内为键类型，右边为值类型
newMap := map[string]string{
    "name": "wlc",
    "age": "18",
    "company": "meituan"
}
```
其中**KeyType是键的类型**，**ValueType是值的类型**，**initialCapacity是可选的参数**，用于指定map的初始容量，map的容量是指map中可以保存的键值对的数量，当map中的键值对的数量达到容量的时候，map会**自动扩容**，如果不指定initialCapacity，Go语言会根据实际情况来选择合适的值

## map的相关操作
### 获取元素
想要获取一个键值对直接通过key即可获取

```go
//通过key获取val
name_val := newMap["name"]

//此时将获取到val值为 wlc
```

### 修改元素
修改一个指定的map元素也可以直接通过key来进行操作，直接赋值更改即可

```go
//通过key获取指定内容进行更改
newMap["age"] = "20"

//此时的"age"属性更改为20
```

### 获取长度
通过len()函数即可

```go
//通过len函数将长度返回给mapLength
mapLength := len(newMap)
```

### 遍历map
遍历一个map可以通过for循环来完成，同时引入range范围

```go
//通过for来进行循环，同时使用range范围来控制循环的次数
for k, val := range newMap {
    fmt.Print(k, val)
}
```

### 删除元素
删除元素的方式与copy方法类似，我们是通过调用一个delete()函数来完成删除功能，其中delete函数需要传入两个参数。第一个参数为需要进行删除操作的map集合；第二个参数为删除的key值，通过key值索引到指定的键值对进行删除

```go
//传入newMap作为操作的原集合，"banana"作为key
delete(newMap, "banana")
```