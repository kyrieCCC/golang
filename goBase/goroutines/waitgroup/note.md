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

### 三个方法的对比
1. 在三个方法中，使用Add方法添加计数，使用Done方法减掉一个计数，如果计数不为0，则会阻塞Wait方法执行
2. 一个goroutine调用Add方法来设定等待的goroutine的数量
3. 每个被等待的goroutine在结束时可以调用Done方法
4. 在主goroutine里面可以调用Wait方法阻塞至所有goroutine结束

## 写法盘点
以下为协程进行时阻塞的相关写法
### 1. time
主线程为了等待goroutine都运行完毕，不得不在程序的末尾使用time.Sleep()方法来睡眠一段时间，等待其他线程充分运行。一般来说我们会使用到time.Sleep(time.Second)方法，这个方法表示暂停一秒钟的时间

但是在实际的开发中，一秒的时间往往是不够的，这个时候我们往往就不适用这个方式了

### 2. 通道
首先可以肯定的是通道可以实现我们关于阻塞的需求，但是管道如此使用未免有一些大材小用。通道的设计之初不仅仅只是用作简单的同步处理，当我们假设有一百个甚至是一万个for循环需要运行的时候，make出这么多的通道对内存而言是一个不小的消耗

### 3. waitgroup
同样的为了解决同一个需求但是不想有太多的不足，waitgroup可以更加方便的帮助我们达到这个目的，waitgroup方法有三个核心方法，具体使用见上文

相对于前面两个写法而言，waitgroup更加轻巧方便

## 注意点
1. 计数器不能为负值
2. waitgroup对象不是一个引用类型