# account

## 1. 账户类型
- 外部账户：账户余额balance和计数器nonce  
合约账户：balance和nonce之外还有code(代码)、storage(相关状态-存储)，我的理解：合约中的 storage 变量会保存在 MPT 中，像 memory 等数据只会短暂保存在内存中

## 2. nonce
- 每发起一笔交易，账户的 nonce 会增加1，用于防止重放攻击
- nonce 太小，交易会被拒绝
- nonce 太大，交易会 pendding，直到补齐中间缺少的 nonce
