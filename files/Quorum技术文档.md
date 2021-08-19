# Quorum 技术文档

## 1. Qubernetes

### 1.1 说明

- Qubernetes 支持 Raft 和 IBFT 共识算法，可以兼容多个版本、任意数量的节点。

### 1.2 依赖
- 需要给 docker 6.25G 的内存，<https://docs.docker.com/desktop/mac/>
- 安装 qctl
	```shell
	# qctl，和 Qubernetes 交互的命令行工具
	GO111MODULE=on go get github.com/ConsenSys/qubernetes/qctl
	qctl --help
	```

### 1.3 启动一个网络
- 方式一
	```shell
	git clone https://github.com/ConsenSys/qubernetes.git
	cd qubernetes
	./quickest-start.sh
	
	./connect.sh node1 quorum

	# 关闭
	./quickest-stop.sh
	```

- `推荐`方式二：
	```shell
	minikube start --memory 6144

	qctl init
	# 按照上一步执行结果 export
	export QUBE_CONFIG=/Users/sunny/test/qubernetes.generate.yaml
	qctl generate network --create
	# 按照上一步执行结果 export
	export QUBE_K8S_DIR=/Users/sunny/test/out
	qctl deploy network --wait
	qctl geth exec quorum-node1 'eth.blockNumber'

	# 关闭
	qctl delete network
	docker stop minikube
	```

## 2. 节点

### 2.1 依赖
- 
	```shell
	# latest 版本拉不下来，使用21.7
	docker pull quorumengineering/quorum:21.7
	docker pull quorumengineering/tessera:21.7
	```

## 3. 工具

### 3.1 cakeshop

- 使用 quorum-wizard 命令，选择 `Simple Network` 或者 `Custom Network`，后面最好选择 `docker-compose` 的部署方式，因为这种方式可以使用所有的工具，其他选项默认就行。

- Cakeshop 包含区块链浏览器，可以查看到交易、合约和链的配置。它自带了一个 sandbox，和 Remix-IDE 类似。还有权限管理、钱包、节点管理的功能。

- Reporting














