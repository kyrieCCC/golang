# WaitGroup
## 定义
在Go语言当中，**sync.WaitGroup**结构体对象用于等待一组线程的结束；WaitGroup是Go并发中最常用的工具，我们可以通过WaitGroup来**表达对这一组协程的任务是否完成**，以决定是否继续往下走，或者取任务结果

WaitGroup的结构体
```go
type WaitGroup struct {
    noCopy noCopy
    state1 [3]uint32
}
```
## 主要方法
### 1. Add()
Add方法主要为WaitGroup的等待数+1或者+n

```go
//Add方法
func (*WaitGroup) Add()
```

+ Add()方法内部计数器加上**delta**，delta可以是负数
+ 如果内部计数器变为0，则**Wait()方法会将处于阻塞等待的所有goroutine释放**
+ 如果计数器小于0，则调用panic()函数
+ Add()方法加上正数的调用**应在Wait()方法之前**，否则Wait方法可能只会等待很少的goroutine
+ Add()方法在创建新的goroutine或者其他等待的事件之前被调用

### 2. Done()
Done函数调用的也是Add函数，主要用于-1的操作

```go
//Done方法
func(wg *WaitGroup) Done()
```

+ Done()方法会减少WaitGroup计数器的值，一般在goroutine的最后执行

### 3. Wait()
Wait函数阻塞当前协程，直到等待数目归为0后才继续向下执行

```go
//Wait方法
func (wg * WaitGroup) Wait()
```

+ Wait方法会阻塞，直到WaitGroup的计数器为0