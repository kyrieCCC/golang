# Go错误处理
go语言通过内置的错误接口提供了非常简单的错误处理机制

error类型是一个接口类型，定义如下：
```go
type error interface {
    Error() string
}
```

我们可以在编码中通过实现error接口类型来生成错误信息

函数通常在最后的返回值中返回错误信息。使用error.New可以返回一个错误信息：
```go
func Sqrt(f float64)(float64, error) {
    if f < 0 {
        return 0, error.New("negative number!")
    }
    else {
        return f**2
    }
}
```