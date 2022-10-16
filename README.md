# BackupAgent 

Gin-MysqlBak 客户端，使用go-micro v2编写，用于完成各种备份任务

版本： v3.0.0

## Get Start

克隆仓库

```shell
git clone https://github.com/noovertime7/gin-mysqlbak-agent.git
```

开始之前，请先将配置文件中的autoInit置为true，以便自动初始化表，但仍需要手动创建数据库

```sql
create datebase `gin-mysqlbak-agent`;
```
### 二进制部署

选择合适环境的安装包，或者自行编译

https://github.com/noovertime7/gin-mysqlbak-agent/releases/tag/v3.0.0


### docker容器部署
开始之前，请先创建config.ini配置文件，并确保数据库已经手动创建

```shell
docker run -itd --name gin-mysql-agent  \
--net=host --restart=always \
-v /root/config.ini:/app/domain/config/config.ini \
-v /root/bakfile:/app/bakfile \
chenteng/gin-mysqlbak-agent:3.0.0
```

### kubernetes部署

开始之前，请先修改agent-conf.yaml配置文件，并确保数据库已经手动创建

PS:如果对端口有个性化需求，请修改agent-deploy.yaml

```shell
## 创建命名空间
kubectl create ns mysqlbak
## 创建配置文件
kubectl apply -f agent-conf.yaml
## 创建deployment & service
kubectl apply -f agent-deploy.yaml
```