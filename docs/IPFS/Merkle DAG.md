# Merkle DAG(Merkle directed acyclic graph, 默克尔有向无环图)

[参考文档 2018-1-22](https://cloud.tencent.com/developer/news/59427)

## 1. 特点
- 内容寻址：使用多重哈希来唯一识别一个数据块的内容 ?
- 防篡改：可以方便的检查哈希值来确认数据是否被篡改
- 去重：由于内容相同的数据块哈希是相同的，可以很容去掉重复的数据，节省存储空间 ?

⚠️注意：IPFS 的 Merkle DAG 中不存储数据

## 2. 将一个文件构建为 Merkle DAG
- 文件大小：3646K
- 文件会被分成15个 block，每个 block 256k（除了最后一个）

![DAG](../../images/DAG1.png)

## 3. 将一个文件夹构建为 Merkle DAG
![DAG](../../images/DAG2.png)
![DAG](../../images/DAG3.png)
