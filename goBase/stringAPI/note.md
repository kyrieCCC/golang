# Go语言中字符串的方法
## API
### 1. strings.Contains
该方法用于判断字符串中是否包含指定的字符和字符串，返回布尔值的结果

使用方法：
```go
strings.Contains(str, targetStr)
//返回boolean值
```

参数：
1. str，需要进行判断的源字符串
2. targetStr，需要判断是否存在的字符或者字符串

### 2. strings.Count
该方法用于判断字符串中的某个字符或者字符串出现的次数，返回int类型

使用方法：
```go
strings.Count(str, targetStr)
//返回一个数值表示出现的次数
```

参数：
1. str，需要进行判断的源字符串
2. targetStr，需要判断次数的字符或者字符串

### 3. strings.HasPrefix
该方法用于检测字符串是否以指定的前缀开头，返回布尔值的结果

使用方法
```go
strings.HasPrefix(str, preStr)
//返回布尔值
```

参数：
1. str，需要进行判断的源字符串
2. preStr，指定的前缀（一般是一个字符串）

### 4. strings.HasSuffix
该方法用于检测字符串是否以指定的内容结尾，返回布尔值的结果

使用方法
```go
strings.HasSuffix(str, endStr)
//返回布尔值
```

参数：
1. str，需要进行判断的源字符串
2. endStr，指定的结尾（一般是一个字符串）

### 5. strings.Index
该方法用于检测指定的字符串第一次出现的位置，返回index下标

使用方法：
```go
strings.Index(str, targetStr)
//返回第一次出现的下标
```

参数：
1. str，需要进行判断的源字符串
2. targetStr，需要判断第一次出现位置的字符或者字符串

### 6. strings.Join
该方法可以将一个字符串切片中的元素连接成为单个字符串，其中可以设置连接符，返回一个string字符串

这里的join方法其实与JS中的join方法类似，但是有一个不同点，go语言是一个强类型的语言，当JS在执行join方法的时候会存在一个隐式类型转换的过程，也就是说我们可以将一个纯数字的数组转换为一个字符串，但是go语言中是不允许这样做的，因此我们传入的切片或者数组必须是string类型

使用方法：
```go 
strings.Join(string_slice, connector)
//返回字符串切片通过连接符拼接后的结果字符串
```

参数：
1. string_slice，需要进行连接的字符串切片
2. connector，连接符

### 7. strings.Repeat
该方法可以将一段内容重复叠加，类似于str+str的操作，返回一个新的字符串

使用方法：
```go
strings.Repeat(str, num)
//返回str经过num叠加后的结果字符串
```

参数：
1. str，需要进行重复的源字符串
2. num，重复的次数

### 8. strings.Replace
该方法会使用一个新的字符来替换源字符串中的字符，我们可以指定替换的某几个元素或者是全部，返回一个新的字符串

使用方法：
```go
strings.Replace(str, oldStr, newStr, n)
//返回一个替换后的字符串
```

参数：
1. str，需要进行替换操作的源字符串
2. oldStr，需要被替换的字符
3. newStr，需要新替换进来的字符
4. n，表示操作的次数，如果n < 0，则全部匹配

### 9. strings.Split
该方法会将一个字符串通过某个连接符转换为一个切片的类型，返回一个切片或者数组

该方法与JS中的Split方法类似

使用方法：
```go
strings.Split(str, connector)
//返回一个切片
```

参数：
1. str，需要进行替换操作的源字符串
2. connector，用于分割的连接符

### 10.strings.ToLower
该方法会将字符串全部转换为小写的形式，返回一个新的字符串

使用方法：
```go
strings.ToLower(str)
//返回全为小写的新的字符串 
```
参数：
1. str，需要进行小写转换操作的源字符串

### 11.strings.ToUpper
该方法会将字符串全部转换为大写的形式，返回一个新的字符串

使用方法：
```go
strings.ToUpper(str)
//返回全为大写的新的字符串 
```
参数：
1. str，需要进行大写转换操作的源字符串
