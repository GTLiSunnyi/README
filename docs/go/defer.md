# defer

- 先进后出  
如果先将前面申请的资源释放掉了，对于后面的资源可能会造成影响。
- go 中 return 是 1.先把返回值赋给一个变量，再 2.return。defer是在1和2之间执行。
- defer + recover 可以捕获代码中的 panic  
```go
defer func() {
	switch p := recover(); p {
	case EightyPanic{}:
		fmt.Println("EightyPanic")
	default:
		panic(p)
	}
}()
```

- 使用场景
> 延迟 close 和 recover panic
