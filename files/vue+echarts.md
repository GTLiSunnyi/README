# vue + echarts

## 1. 依赖
- 
	```shell
	npm install vue -g
	npm install vue-cli -g

	# 初始化项目
	vue init webpack "data-demo"
	
	# 运行
	npm run dev
	```

## 2. 在 vue 中引入 echarts
- 分为[全局引用和局部引用](https://www.cnblogs.com/xiongzuyan/p/10417968.html)。在多个组件中使用 echarts 时，使用全局引用。
- ⚠️：使用`import * as echarts from 'echarts'`引用
