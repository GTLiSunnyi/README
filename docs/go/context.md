# context

## 1. 定义：
```go
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}
```

## 2. 使用
- 定义空的context：
```go
ctx := context.Background()
todoCtx := context.TODO()
```

context 主要用在协程之间传递上下文信息，即使协程中没有用到 context，也建议养成传递 context 的好习惯
- 1. 取消信号
```go
context.WithCancel()
```
- 2. 超时信号
```go
ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second) // 两秒后会自动调用cancel()
```
- 3. 截止时间信号
```go
ctx, cancel := context.WithDeadline(context.Background(), /*unix时间*/) // 和WithTimeout类似
```
- 4. key-value 值
```go
ctx1 := context.WithValue(context.Background(), "key1", "value1")
ctx2 := context.WithValue(ctx1, "key2", "value2")
fmt.Println(ctx2.Value("key1")) // value1
```
