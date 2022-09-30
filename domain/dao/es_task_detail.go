package dao

type EsTaskDetail struct {
	HostInfo   *HostDatabase `json:"host_info"`
	ESTaskInfo *EsTaskDB     `json:"es_task_info"`
}
