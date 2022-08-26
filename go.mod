module backupAgent

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/aliyun/aliyun-oss-go-sdk v0.0.0-20190307165228-86c17b95fcd5
	github.com/fortytw2/leaktest v1.3.0 // indirect
	github.com/go-errors/errors v1.0.1
	github.com/go-xorm/xorm v0.7.9
	github.com/golang/protobuf v1.5.2
	github.com/gorhill/cronexpr v0.0.0-20180427100037-88b0669f7d75
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1
	github.com/minio/minio-go/v7 v7.0.34
	github.com/noovertime7/mysqlbak v0.0.0-20220612083217-fdb12cd90242
	github.com/olivere/elastic v6.2.37+incompatible
	github.com/opentracing/opentracing-go v1.1.0
	github.com/robfig/cron/v3 v3.0.1
	github.com/sirupsen/logrus v1.9.0
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	google.golang.org/protobuf v1.28.0
	gopkg.in/ini.v1 v1.66.6
	gorm.io/driver/mysql v1.3.5
	gorm.io/gorm v1.23.8
)
