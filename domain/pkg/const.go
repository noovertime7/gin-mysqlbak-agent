package pkg

const (
	HistoryStatusAll     = "all"
	HistoryStatusSuccess = "success"
	HistoryStatusFail    = "fail"
	LargePageSize        = 99999
)

type HostType int64

const (
	MysqlHost   int64 = 1
	ElasticHost int64 = 2
)

// 加解密相关
var b58 = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
