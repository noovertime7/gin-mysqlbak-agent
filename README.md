English | [简体中文](./README_zh_cn.md)
# BackupAgent

Gin-MysqlBak client, written in go-micro v2, for various backup tasks

Version: v3.0.0

## Get Start

Clone the repository

```shell
git clone https://github.com/noovertime7/gin-mysqlbak-agent.git
```

Before you start, set autoInit to true in the configuration file to automatically initialize the tables, but you still need to create the database manually

```sql
create datebase `gin-mysqlbak-agent`;
```
### Binary deployment

Choose the right installation package for your environment, or compile it yourself

https://github.com/noovertime7/gin-mysqlbak-agent/releases/tag/v3.0.0


### docker container deployment
Before you start, create the config.ini configuration file and make sure the database has been created manually

```shell
docker run -itd --name gin-mysql-agent \
--net=host --restart=always \
-v /root/config.ini:/app/domain/config/config.ini \
-v /root/bakfile:/app/bakfile \
chenteng/gin-mysqlbak-agent:3.0.0
```

### kubernetes deployment

Before you start, please modify the agent-conf.yaml configuration file and make sure the database has been created manually

PS: If you have individual requirements for ports, please modify agent-deploy.yaml

```shell
## Create namespace
kubectl create ns mysqlbak
## Create configuration file
kubectl apply -f agent-conf.yaml
## Create deployment & service
kubectl apply -f agent-deploy.yaml
```