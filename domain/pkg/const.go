package pkg

const (
	HistoryStatusAll     = "all"
	HistoryStatusSuccess = "success"
	HistoryStatusFail    = "fail"
)

type HostType int64

const (
	MysqlHost   HostType = 1
	ElasticHost HostType = 2
)
