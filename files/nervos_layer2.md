# Nervos layer2——Godwoken

## 1. 简介

- 目前 ckb 的 layer2 还只实现了 PoA 的 optimistic rollup，将来会实现 zk rollup 和 PoS 的 optimistic rollup。

	> PoA 实现：<https://github.com/nervosnetwork/clerkb>

## 2. 基本流程

![layer2](../images/nervos_layer2.png)

1. 任何人都可以创建一个 rollup cell 来开启一个 layer2。

	> 此 layer2 有 256bit 的存储空间（TODO）,layer2 智能合约可以自由使用这个存储空间。  
	> rollup cell 中会包含此 layer2 部署的相关信息。  
	> 存储空间由 Sparse Merkle Tree 构成，SMT 的 root hash 也会保存在 rollup cell 中。  
	> rollup cell 中的 typescript 是特定的状态验证脚本，用于验证每个区块 SMT 的 root hash 是否正确。  
	> 开启的 layer2 可以是公开或者私有的，这取决于 rollup cell 的 lockScript 的限制。

2. 聚合者收集 layer2 中的交易，打包成 layer2 区块，再向 layer1 提交区块。

	> 聚合者需要质押 ckb 才能提交区块。

3. 在 layer2 区块最终确认前，挑战者可以通过质押ckb、标记非法交易来发起挑战。

4. 聚合者执行被挑战的交易来证明交易的合法性，聚合者和挑战者的正确方会得到另一个人的抵押。

5. layer2 区块最终确认后就不能发起挑战了。
	
	> 聚合者需要承担 ckb 的流动成本和 layer2 交易在 layer1 中执行的成本。
