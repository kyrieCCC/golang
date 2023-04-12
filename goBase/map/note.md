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