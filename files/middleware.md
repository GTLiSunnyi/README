# middelware

## 1. 说明

### 1.1 此中间件用于上传存证到链上，或查询链上的存证信息。
- 1. post_invitation 时，会在数据库中存储你的邮箱，返回数据库记录 id。
- 2. register 时，上传上面提到的数据库记录 id 和密码，在数据库中生成一个用户。
- 3. 使用 hmac_sig 工具，使用密码对存证内容、时间戳签名。
- 4. upload 时，上传邮箱、存证内容、时间戳和签名，验证签名正确性，用本地 test.json 文件生成私钥，使用此私钥将数据上传到链上。  

- 不同语言实现的项目不需要额外不同语言的 SDK 去上传存证
- 不需要手动让存证上链

### 1.2 依赖

- postgresql

	```shell
	# macOS，需要安装 homebrew
	brew install postgresql
	# ubuntu 系统可以参考：https://tecadmin.net/install-postgresql-server-on-ubuntu/
	# 查看是否成功
	pg_ctl --version

	# 使用 PostgreSQL 的时候可能会有很多坑，自己多摸索摸索
	# 初始化数据库，只有需要初始化一次
	initdb /usr/local/var/postgres -E utf8

	# 创建超级管理员
	createuser postgres -P -s

	# 启动数据库
	pg_ctl -D /usr/local/var/postgres -l /usr/local/var/postgres/server.log start

	# 创建后面会用到的用户和数据库
	sudo su postgres
	psql
	CREATE ROLE sunny WITH PASSWORD '123456';
	ALTER USER sunny WITH CREATEDB;
	CREATE DATABASE sunny OWNER sunny;
	```

- diesel

	```shell
	cargo install diesel_cli --no-default-features --features postgres
	cd middleware
	echo DATABASE_URL=postgres://username:password@localhost/diesel_demo > .env
	diesel setup
	```

## 2. 

- 
	```shell
	git clone git@gitlab.taidihub.com:flyq/hmac_sig.git
	
	```

## 3. 
![middleware-upload](../images/middleware-upload.png)
