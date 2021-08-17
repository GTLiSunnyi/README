# Quorum

## 1. 说明

- Quorum 是 ConsenSys 和摩根大通合作开发的一种企业以太坊，基于公链以太坊的代码。我们公司投资过 ConsenSys，目前是合作关系，由我们帮助其开拓 Quorum 的中国市场。

- Quorum 相比公链以太坊增强了隐私控制、权限控制、共识机制，以及提高整个链的交易性能，不过 Quorum 的 tps 我还没有找到一个准确的说法。

## 2. 交易过程

![交易流程图](../images/QuorumTransactionProcessing.jpg)

- 这是 2018.1 的资料。

## 3. 逻辑架构

![逻辑架构](../images/逻辑架构.png)

- 在共识方法上, Quorum用QuorumChain取代了以太坊的POW, 最新推出的2.0.0又增加了Raft的共识方法, 最终的目标就是形成一个插件化的共识体系。

- Quorum 包括基于以太坊的节点。Transaction Manager 用于管理私有交易在联盟内部成员之间的同步, Enclave 配合 Transaction Manager 来完成私有交易的密码学相关的处理工作。
