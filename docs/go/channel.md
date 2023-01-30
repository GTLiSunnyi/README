# channel

## 无缓冲
```go
ch := make(chan int)
```
写入阻塞条件：同一时间没有另外一个线程对该chan进行读操作
取出阻塞条件：同一时间没有另外一个线程对该chan进行取操作

## 有缓冲
```go
ch := make(chan int, 1)
```
写入阻塞条件：缓冲区满
取出阻塞条件：缓冲区没有数据
 
- 在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告 deadlock，但是 close后继续取不会，会取出该类型的零值  
- close 后只能读取，不能写入
- for range channel，会在 channel close 之后停止

## 三种状态
- nil，没有被初始化
- active
- closed
