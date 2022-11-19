# Raft

## 1. 与 PBFT 的区别
- Raft 的容错是 2f + 1
- Raft 的 follower 是绝对相信 leader 节点的，而 PBFT 不是
