# (multi-)paxos

## 1. paxos
- paxos 可以同时存在几个 Proposer
![paxos 流程图](../../images/paxos.png)

- 1. Proposer 生成一个 Proposal ID n，为了保证 Proposal ID 唯一，可采用时间戳+ Server ID 生成；
- 2. Prepare: Proposer 向所有 Acceptors 广播 Prepare(n) 请求；
- 3. Promise: Acceptor 比较 n 和 minN，如果 n>minN，将 minN 和 acceptedValue 返回；
- 4. Proposer 接收到过半数回复后，将所有回复中 minN 最大的 acceptedValue 作为本次提案的 value，否则可以任意决定本次提案的 value；
- 5. Propose: 广播 Accept (n,value) 到所有节点；
- 6. Accept: Acceptor 比较 n 和 minN，如果 n>=minN，则 minN=n，acceptedValue=value，本地持久化后，返回；否则，返回 minN。
- 7. Learn: Proposer 接收到过半数请求后，如果发现有返回值 result >n，表示有更新的提议，跳转到1；否则 value 达成一致。

## 2. multi-paxos
Paxos 只能对一个值形成决议，决议的形成至少需要两次网络来回，在高并发情况下可能需要更多的网络来回，极端情况下甚至可能形成活锁。如果想连续确定多个值，Paxos 搞不定。因此 Paxos 几乎只是用来做理论研究，并不直接应用在实际工程中。
