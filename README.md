# BackupAgent Service

Gin-MysqlBak 客户端，使用go-microv2编写，用于完成各种备份任务

Generated with

```
micro new --namespace=go.micro --type=service backupAgent
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.service.backupAgent
- Type: service
- Alias: backupAgent

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./backupAgent-service
```

Build a docker image
```
make docker
```