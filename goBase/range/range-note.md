# range(范围)
## 定义
在go语言当中，range关键字用于for循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回 **元素的索引和索引对应的值**，在集合中返回键值对

例如我们读取一个map：
```go
newMap := map[string]string {
    "name": "wlc",
    "age": "18",
    "home": "hainan"
}

//读取这个map
for key, value :=  range newMap {
    fmt.Println(key, value)
}
```

当然，如果我们只需要key或者是value的时候，我们可以这样做：
```go
//只需要key的时候，直接省略value即可
for key := range newMap {
    fmt.Println(key)
}

//当我们只需要value的时候，用下划线_代替key表示省略
for _, value := range newMap {
    fmt.Println(value)
}
```

让我们来看看range在其他数据结构的使用：
```go
//数组和切片当中，其实与map类似，我们可以遍历索引与值
//例如我们可以尝试遍历一个slice
slice := []int {1,2,3,4,5}

for i, val := range slice {
    fmt.Println(i, val)
}
//此时输出每一个切片对应的索引与值
```