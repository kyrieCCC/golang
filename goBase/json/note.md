# go的序列化操作
## 定义
json是一种数据交换格式，主要是将要传输的数据转换为json特有的格式，其他人拿到这个json数据后，可以通过数据反向转换得到转换前的数据，方便数据传输的一致性

json依赖ket-value键值对来保存数据，而键值对组合中的键名写在前面，并且用双引号 "" 包裹，使用冒号 : 分隔

常见的json格式如下:
```json
{
    "city": "beijing",
    "name": "wlc",
    "age": "18"
}
```

## go语言的json序列化
json序列化是指，将语言中键值对类型的数据转换为json字符串，比如go语言中的结构体、map、数组和切片等等，转换为json格式的过程成为格式化

在go语言中序列化可以通过 **Marshal**方法
```go
func Marshal(v interface{}) ([]byte, error)

//通过文档得知，Marshal函数接收的是一个空接口类型
//意思是我们可以将go的任意类型的数据传入进去
//而返回值是一个byte的切片，以及error
```

接下来我们尝试序列化一个结构体
```go
type myStruct struct {
    Name string
    Age int 
    Sex string
}

func main() {
    subStruct := myStruct {
        Name: "wlc",
        Age: 18,
        Sex: "boy",
    }

    data, err := json.Marshal(subStruct)

    //输出序列化后的结果
    fmt.Println(data)
}
```

### 标签
我们在上面序列化结构体的时候，实际上就是把字段的key作为json的key使用，但在浏览器中我们更希望接收到的是首字母小写的字段，于是我们可以引入标签的方式

标签的使用非常简单，我们只需要在我们定义的结构体后使用``反引号来声明一个别名即可

```go
type Person struct {
    Name string
    Age int `json: "age"`
    Like string
}

//在上方将Age字段声明为age，这样在格式化之后出现的字段名即为小写的age
```