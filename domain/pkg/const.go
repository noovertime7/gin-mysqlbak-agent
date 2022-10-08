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
