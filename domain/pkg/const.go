package pkg

const (
	HistoryStatusAll     = "all"
	HistoryStatusSuccess = "success"
	HistoryStatusFail    = "fail"
	LargePageSize        = 99999
)

type HostType int64

const (
	MysqlHost   HostType = 1
	ElasticHost HostType = 2
)