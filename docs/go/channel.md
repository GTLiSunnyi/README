# channel

## 无缓冲
- 如果没有协程接收，发送端就会阻塞，deadlock

## 有缓冲
- 只有队列满了，发送端才会阻塞，deadlock
 
- 在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告 deadlock，但是close后继续取不会，会取出该类型的零值  
- close后只能读取，不能写入
- for range channel，会在 channel close 之后停止

## 三种状态
- nil，没有被初始化
- active
- closed
