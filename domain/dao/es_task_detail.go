package dao

type EsTaskDetail struct {
	HostInfo   *HostDatabase `json:"host_info"`
	ESTaskInfo *TaskInfo     `json:"es_task_info"`
}
