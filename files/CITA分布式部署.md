# CITA 辅助文档

## 1. 说明

- 此文档给使用者提供部署 Toolchain 和分布式部署 CITA（v20.2.0） 的步骤、流程。

* [CITA 文档](https://docs.citahub.com/zh-CN/cita/cita-intro)中需要掌握的内容为：
	1. 快速入门
	2. 配置说明（链级配置）
	3. 部署指南（分布式部署）
	4. 命令说明
	5. 节点管理
	6. 经济模型
	7. 账户与权限管理
	8. 隐私保护（了解）
	9. 存证
	10. 进阶使用（solidity 合约开发）
	11. 架构设计（了解）
	12. 工具链（除了 Cyton、IDE、Truffle Box、Web Debugger、Testnet，其他要求会部署）
	13. first-forever-demo

- ⚠️ 文档中 xx.xx.xx.xx 处使用 `hostname -I` 命令获取的第一个 ip（10.10.0.202）替换，而不是使用 47.102.199.70

## 2. 分布式部署（两个节点）

* 两台服务器ip地址分别为  
	node0（ubuntu16.04、superadmin）: 47.102.199.70  
	node1（ubuntu16.04）: 47.102.201.40
  
* 依赖  
	1. wget
		```shell
		apt-get update
		apt-get install wget
		```

	2. 根据[文档](https://docs.citahub.com/zh-CN/cita/getting-started/setup#安装-cita-客户端工具)在 node0、node1 上安装 cita、cita-cli。

	3. docker
		```shell
		# node0
		ssh root@47.102.199.70
		curl -sSL https://get.daocloud.io/docker | sh
		apt install docker-compose

		# node1
		ssh root@47.102.201.40
		curl -sSL https://get.daocloud.io/docker | sh
		```
		- 给 docker 配置镜像：<https://www.runoob.com/docker/docker-mirror-acceleration.html>
	
	4. git
		```shell
		apt install git
		```

	5. rails
		```shell
		# node0
		apt install ruby-railties
		apt install ruby-bundler
		```

	6. npm
		```shell
		# ndoe0
		apt install npm
		npm config set registry https://registry.npm.taobao.org
		npm install n -g
		n 10.13.0
		export NODE_HOME=/usr/local
		export PATH=$NODE_HOME/bin:$PATH
		export NODE_PATH=$NODE_HOME/lib/node_modules:$PATH

		node -v
		```

### 2.1 分别生成 key

- 
	```shell
	cita-cli key create
	```

- node0、node1 地址分别为：0x2d2ed0d5b613ef9078decdc4ba54f99fe9446be4，0x7d5456e143def3724fb9dc97a258e2467df4160b。

### 2.2 管理员使用 ***`真实`*** 的 ip 地址初始化 cita 链
- 
	```shell
	cd /data/cita/cita_secp256k1_sha3
	bin/cita create --super_admin 0x2d2ed0d5b613ef9078decdc4ba54f99fe9446be4  --nodes "47.102.199.70:4000,47.102.201.40:4000" --contract_arguments SysConfig.blockInterval=20000  
	```
- ⚠️：这里设置出块速度为20s，具体其他链的参数配置可参考[文档](https://docs.citahub.com/zh-CN/cita/configuration-guide/chain-config)。

### 2.3 node0 将 test-chain 中的内容复制到 node1

- 
	```shell
	scp -r test-chain 47.102.201.40:/data/cita/cita_secp256k1_sha3
	```
- 这一步如果出错的话，可以参考 <https://www.cnblogs.com/guodavid/p/11004499.html> 的解决方案一，尽量 ***`不要`*** 使用方案二。

### 2.4 分别在服务器上启动节点

- 以下以 node0 做演示
	- 启动
		```shell
		cd /data/cita/cita_secp256k1_sha3
		bin/cita setup test-chain/0
		bin/cita start test-chain/0
		```
	- 查看是否成功启动
		```shell
		bin/cita top test-chain/0
		```

	- 如果显示内容和以下内容相似，则正确启动
		```shell
		root      2542     1  0 15:37 ?        00:00:00 cita-forever  
		root      2559  2542  2 15:37 ?        00:00:00 cita-auth -c auth.toml  
		root      2560  2542  1 15:37 ?        00:00:00 cita-bft -c consensus.toml -p 	privkey  
		root      2563  2542  1 15:37 ?        00:00:00 cita-chain -c chain.toml  
		root      2554  2542  4 15:37 ?        00:00:00 cita-executor -c executor.toml  
		root      2548  2542  0 15:37 ?        00:00:00 cita-jsonrpc -c jsonrpc.toml  
		root      2553  2542  1 15:37 ?        00:00:00 cita-network -c network.toml  
		```
- 在 node1 上重复以上操作
- node1 的 cita-cli 需要使用下面的命令连接到它本地的服务
	```shell
	cita-cli
	switch --url http://127.0.0.1:1338
	```

### 2.5 查看是否正常出块

- 
	```shell
	rpc blockNumber
	```

## 3. 在 node0 上部署 rebirth
- 
	```shell
	git clone https://github.com/citahub/re-birth.git
	```

- [安装 rvm](https://blog.csdn.net/weixin_33970449/article/details/89758010)，将 ruby 切换到 2.5.3 版本。

- 接着执行以下命令启动 rebirth
	```shell
	cd re-birth
	gem install -n /usr/local/bin cocoapods
	apt-get install libpq-dev
	gem install pg
	gem install bundler:1.17.1
	bundle update mimemagic
	make setup
	docker-compose run --rm app bundle exec rails secret
	```

- 最后一步输出结果为
	```shell
	Starting rebirth_db_1 ... 
	Starting rebirth_db_1 ... done
	8ee35bb6809a823ed3ab1779edeb5ac244ffab2bb76017a8dbcd754d44bb8da8a3391cf929dfccd9d94307cfc04632e4b77278f5f543bb15bb74f49526e34284
	```

- 将 .env.local 中的 SECRET_KEY_BASE 字段改为上文中输出的值：8ee35bb6809a823ed3ab1779edeb5ac244ffab2bb76017a8dbcd754d44bb8da8a3391cf929dfccd9d94307cfc04632e4b77278f5f543bb15bb74f49526e34284

- 将 .env 中的 CITA_URL 字段改为 "http://xx.xx.xx.xx:1337/"

- 最后启动 rebirth
	```shell
	make up
	```

- 访问<http://47.102.199.70:8888>，如果出现下面的内容，则为正确启动。  
{"message":"Read more API interface info at https://github.com/cryptape/re-birth"}

## 4. 在 node0 上部署 microscope
- 
	```shell
	git clone https://github.com/citahub/microscope-v2.git
	cd microscope-v2
	docker build -t microscope .
	docker run --name microscope -d -p 80:80 microscope
	```

- 访问<http://47.102.199.70:80>，连接 rebirth 的地址：<http://47.102.199.70:8888>，可以看到出块间隔为 20s，随后区块高度也会刷新，这代表 microscope 成功运行。

## 5. 在 node0 上部署 [first-forever-demo](https://github.com/citahub/first-forever-demo/blob/develop/README-CN.md)

- 
	```shell
	# 在下面的“github.com”后面加上“cnpmjs.org”可以不用翻墙下载，所有的“git clone”都可以泗洪这种方式
	git clone https://github.com.cnpmjs.org/citahub/first-forever-demo.git
	cd first-forever-demo
	npm install
	cp .env.example .env
	cp ./src/config.js.example ./src/config.js
	npm run deploy

	# copy contractAddress to src/config.js
	# configure the other two variables
	# Just fill in a legal private key
	# chain: 'http://47.102.199.70:1337',
  	# privateKey: '0x934c13d78dbc092da41c10efb26092cc6d5062f92668823e024a583ce7b7ccf9',

	npm start
	```

- 打开网址<http://47.102.199.70:3000>，就可以了。

## 6. 在 node0 上部署 monitor

- 
	```shell
	git clone https://github.com.cnpmjs.org/citahub/cita-monitor.git
	cd cita-monitor
	git submodule update --init
	```

### 6.1 部署 agent

#### 6.1.1 单节点部署
- 
	```shell
	cd agent/cita_exporter
	docker build -t citamon/agent-cita-exporter .
	cd ..

	docker run -d --name="citamon_agent_cita_exporter_1337" \
	--pid="host" \
	-p 1920:1923 \
	-v /etc/localtime:/etc/localtime \
	-v "/data/cita/cita_secp256k1_sha3/":"/data/cita/cita_secp256k1_sha3/" \
	-v "/data/cita/cita_secp256k1_sha3/test-chain/0":"/data/cita/cita_secp256k1_sha3/test-chain/0" \
	-e NODE_IP_PORT="xx.xx.xx.xx:1337" \
	-e NODE_DIR="/data/cita/cita_secp256k1_sha3/test-chain/0" \
	citamon/agent-cita-exporter

	# 检查是否正常运行
	curl http://localhost:1920/metrics/cita
	```

#### 6.1.2 编排部署

- 
	```shell
	cd agent
	cp .env.example .env
	vi .env
	```

- 参考配置
	```shell
	# citamon_agent_cita_exporter HostName
	HOSTNAME=iZuf6098m9sg0sswpg1mcyZ

	# CITA node ip and port
	NODE_IP=xx.xx.xx.xx
	NODE_PORT=1337

	# CITA node directory analysis
	NODE_DIR=/data/cita/cita_secp256k1_sha3/test-chain/0
	SOFT_PATH=/data/cita/cita_secp256k1_sha3/

	# CITA NODE INFO
	CITA_NODENAME=node_0
	CITA_CHAIN_ID=test-chain/0
	CITA_NETWORKPORT=4000
	```

- 启动
	```shell
	docker-compose up -d

	# 查看数据采集信息
	#citamon_agent_host_exporter
	curl http://localhost:1920/metrics/host

	#citamon_agent_process_exporter
	curl http://localhost:1920/metrics/process

	#citamon_agent_rabbitmq_exporter
	curl http://localhost:1920/metrics/rabbitmq

	#citamon_agent_cita_exporter
	curl http://localhost:1920/metrics/cita
	```

### 6.2 部署 server

- 
	```shell
	cd ../server
	cp .env.example .env
	# 可以在 .env 中将语言设置为 cn，也可以不配置 .env
	```

- 修改 prometheus 配置
	```shell
	cd config
	vi prometheus.yml
	```

- 将其中的 citamon_server_alertmanager 和 citamon_agent_host_ip 换成 xx.xx.xx.xx

- 运行
	```shell
	cd ..
	export VERSION=`cat ../VERSION`
	docker-compose build
	docker-compose up -d

	# 查看是否正常运行
	docker ps -a
	```

- 另外还可以配置alert、grafana 等，详见<https://github.com/citahub/cita-monitor/blob/master/server/README.md#步骤>中的第二步。

### 6.4 访问 monitor
- 进入[页面](http://47.102.199.70:1919)后，点击左上角的 `Home`，随后可以选择查看的内容。

## 7. CITA 服务重启
* 重启节点
* 重启每个节点的 agent 服务
* 存储不够可以删除 re-birth/logs 下的文件
* 使用 `docker-compose logs -f citamon_server_alertmanager` 命令查看 alertmanager 是否报错
