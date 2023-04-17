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

## json反序列化
与序列化相反，反序列化就是将我们的序列化转换为go使用的数据类型

在go语言中反序列化可以通过 **Unmarshal**方法
```go
func Unmarshal(data []byte, v interface{}) error
 
//传入的值是一个切片类型
//后面的v interface是指定我们用于接收json数据的结构体
//是否接收成功返回一个error的值
```

使用示例：
```go

func unmarshalStruct(){
	//我们将这个json写入到变量，不过因为他本身有双引号，这里通过反斜杠将双斜杠转义
	//在实际开发中，这个str的字符串是通过前端POST上传或者读取文件得到的
	str := "{\"Name\":\"牛魔王\",\"Age\":500,\"Birthady\":\"2011-11-11\",\"Sal\":8000,\"Skill\":\"牛魔拳\"}"
 
 
	var monster Monster       //定义一个结构体，去接收反序列化json的结果
 
	err := json.Unmarshal([]byte(str),&monster)  //反序列化，通过[]byte(str)类型断言将str转换为切片
	if err != nil{
		fmt.Printf("unmarshal err=%v\n",err)
	}
 
	fmt.Printf("反序列化后 monster=%v",monster)
}
```
具体操作见相应的json.go文件