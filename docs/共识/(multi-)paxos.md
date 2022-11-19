# (multi-)paxos

## 1. paxos
- paxos 可以同时存在几个 Proposer
![paxos 流程图](../../images/paxos.png)

- 1. Proposer生成一个Proposal ID n，为了保证Proposal ID唯一，可采用时间戳+Server ID生成；
- 2. Prepare: Proposer向所有Acceptors广播Prepare(n)请求；
- 3. Promise: Acceptor比较n和minN，如果n>minN，将 minN 和 acceptedValue 返回；
- 4. Proposer接收到过半数回复后，如果发现有acceptedValue返回，将所有回复中minN最大的acceptedValue作为本次提案的value，否则可以任意决定本次提案的value；
- 5. Propose: 广播Accept (n,value) 到所有节点；
- 6. Accept: Acceptor比较n和minN，如果n>=minN，则minN=n，acceptedValue=value，本地持久化后，返回；否则，返回minN。
- 7. Learn: Proposer接收到过半数请求后，如果发现有返回值result >n，表示有更新的提议，跳转到1；否则value达成一致。

## 2. multi-paxos
Paxos只能对一个值形成决议，决议的形成至少需要两次网络来回，在高并发情况下可能需要更多的网络来回，极端情况下甚至可能形成活锁。如果想连续确定多个值，Paxos搞不定。因此Paxos几乎只是用来做理论研究，并不直接应用在实际工程中。
